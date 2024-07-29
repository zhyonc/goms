package server

import (
	"context"
	"crypto/rand"
	"fmt"
	"goms/maple"
	"goms/mongodb"
	"goms/mongodb/model/character"
	"goms/mongodb/model/inventory"
	"goms/mongodb/model/social"
	"goms/network"
	"goms/network/client"
	"goms/network/server/channel"
	"goms/network/server/listener"
	"goms/nx"
	"goms/nxfile"
	"goms/packet/inpacket"
	"goms/packet/object"
	"goms/packet/outpacket"
	"goms/util"
	"log/slog"
	"net"
	"path"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type channelServer struct {
	worldServer network.WorldServer
	worldID     maple.WorldID
	index       uint8
	addr        string
	dbClient    *mongodb.DBClient
	tcpListener listener.Listener
	clientMap   sync.Map
	fieldMap    sync.Map
	cancel      context.CancelFunc
}

func NewChannelServer(worldServer network.WorldServer, worldID maple.WorldID, index uint8, addr string, dbClient *mongodb.DBClient) network.ChannelServer {
	s := &channelServer{
		worldServer: worldServer,
		worldID:     worldID,
		index:       index,
		addr:        addr,
		dbClient:    dbClient,
	}
	s.tcpListener = listener.NewTCPListener(addr, s.HandleTCPConn)
	return s
}

// HandleTCPConn implements network.ChannelServer.
func (s *channelServer) HandleTCPConn(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	slog.Info("New client connected", "addr", addr)
	ip, _, _ := net.SplitHostPort(addr)
	ok := s.worldServer.CheckAuthorizedIP(ip)
	if !ok {
		slog.Warn("Unauthorized IP", "addr", addr)
		conn.Close()
		return
	}
	var recvIV [4]byte
	var sendIV [4]byte
	rand.Read(recvIV[:])
	rand.Read(sendIV[:])
	c := client.NewChannelClient(conn, recvIV, sendIV, s)
	go c.RecvPacket()
	_, err := conn.Write(outpacket.NewConnect(recvIV[:], sendIV[:]))
	if err != nil {
		slog.Error("Failed to send connect packet", "addr", conn.RemoteAddr())
		c.Disconnect()
		return
	}
	s.KickClient(ip)
	s.clientMap.Store(ip, c)
}

// Run implements network.ChannelServer.
func (s *channelServer) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				s.clientMap.Range(func(key, value any) bool {
					client, ok := value.(network.ChannelClient)
					if !ok {
						return true
					}
					// Save online client game data
					s.SaveGameData(client)
					return true
				})
				s.fieldMap.Range(func(key, value any) bool {
					field, ok := value.(*channel.Field)
					// Clear inactive field
					if ok && field.GetIPsLength() == 0 {
						s.fieldMap.Delete(key)
					}
					return true
				})
			case <-ctx.Done():
				return
			}
		}
	}()
	// Start tcp listener
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.tcpListener.Start()
	}()
	wg.Wait()
}

// Stop implements network.ChannelServer.
func (s *channelServer) Stop() {
	s.tcpListener.Stop()
}

// KickClient implements network.ChannelServer.
func (s *channelServer) KickClient(ip string) {
	temp, ok := s.clientMap.Load(ip)
	if !ok {
		return
	}
	c, ok := temp.(network.ChannelClient)
	if !ok {
		return
	}
	s.clientMap.Delete(ip)
	s.SaveGameData(c)
}

// GetClientCount implements network.ChannelServer.
func (s *channelServer) GetClientCount() uint32 {
	var count uint32 = 0
	s.clientMap.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// SaveGameData implements network.ChannelServer.
func (s *channelServer) SaveGameData(c network.ChannelClient) {
	char := c.GetChar()
	inv := c.GetInv()
	soc := c.GetSoc()
	if char == nil || inv == nil || soc == nil {
		return
	}
	s.dbClient.WithTransaction(func(ctx mongo.SessionContext) (any, error) {
		ok := s.dbClient.CharacterUsecase.UpdateModelByID(ctx, char.ID, char)
		if !ok {
			return nil, fmt.Errorf("failed to update character by %d", char.ID)
		}
		ok = s.dbClient.InventoryUsecase.UpdateModelByID(ctx, char.ID, inv)
		if !ok {
			return nil, fmt.Errorf("failed to update inventory by %d", char.ID)
		}
		ok = s.dbClient.SocialUsecase.UpdateModelByID(ctx, char.ID, soc)
		if !ok {
			return nil, fmt.Errorf("failed to update social by %d", char.ID)
		}
		return nil, nil
	})
}

func (s *channelServer) getField(posMap uint32) *channel.Field {
	fieldTemp, ok := s.fieldMap.Load(posMap)
	if !ok {
		return nil
	}
	field, ok := fieldTemp.(*channel.Field)
	if !ok {
		slog.Error("Failed to assert *channel.Field")
		return nil
	}
	return field
}

// RecvMigrateIn implements network.ChannelServer.
func (s *channelServer) RecvMigrateIn(c network.ChannelClient, data []byte) {
	in := inpacket.NewMigrateIn(data)
	if in.WorldID != uint32(s.worldID) {
		c.Disconnect()
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	charTemp := s.dbClient.CharacterUsecase.FindModelByID(ctx, in.CharacterID)
	invTemp := s.dbClient.InventoryUsecase.FindModelByID(ctx, in.CharacterID)
	socTemp := s.dbClient.SocialUsecase.FindModelByID(ctx, in.CharacterID)
	char, ok := charTemp.(*character.Character)
	if !ok {
		c.Disconnect()
		return
	}
	inv, ok := invTemp.(*inventory.Inventory)
	if !ok {
		c.Disconnect()
		return
	}
	soc, ok := socTemp.(*social.Social)
	if !ok {
		c.Disconnect()
		return
	}
	// Bind game data to client
	c.BindGameData(char, inv, soc)
	field := s.getField(char.PosMap)
	if field == nil {
		mapNX := nxfile.GetMap(char.PosMap)
		if mapNX == nil {
			c.Disconnect()
			return
		}
		field = channel.NewField(c.GetClientIP(), mapNX)
	}
	notifiter := &channel.Notifiter{
		BlockReason: "",
	}

	// @Crowns: Hide 5 crowns on character head firstly
	// CWvsContext::OnEventNameTag opcode=286
	c.SendPacket(outpacket.NewEventNameTag())

	// @EventList
	// CWvsContext::OnRequestEventList opcode=282 TODO
	c.SendPacket(outpacket.NewRequestEventList())

	// @Field: Warp the character to a given field
	// Otherwise certain packets specific to the map will not work if it is being sent late
	// CStage::OnSetField opcode=442
	c.SendPacket(outpacket.NewSetField(s.index, true, field.MapNX, notifiter, c.GetDamageSeed(),
		maple.FlagAll, char, inv, 0, 0,
		false, false, 10, false, 0))

	// Others. The orders of how these are placed should be based on its priority.
	// Such as clones, pets which may not really affect the player if should an exception occur should be placed last.
	// Or stuff that may not possibility cause exception could be placed on top with higher priority too.

	// @Ping
	// CClientSocket::OnPingCheckResult opcode=20
	c.SendPacket(outpacket.NewPingCheckResult())

	// @Time: Prepare server time for client events?
	// CWvsContext::OnHourChanged opcode=161
	now := time.Now()
	c.SendPacket(outpacket.NewHourChanged(now.Day(), now.Hour()))

	// @QuestClear: If not send this packet will -10000 HP when collision with mob?
	// CFiled::SetQuestClear opcode=463
	c.SendPacket(outpacket.NewSetQuestClear())

	// @TamingMob
	// CWvsContext::OnSetTamingMobInfo opcode=105
	c.SendPacket(outpacket.NewSetTamingMobInfo(char, false))

	// @Pet
	// CFuncKeyMappedMan opcode=1480-1484
	// Use hp/mp/potion/skill once
	c.SendPacket(outpacket.NewCFuncKeyMappedMan())
	c.SendPacket(outpacket.NewPetConsumeItemID())
	c.SendPacket(outpacket.NewPetConsumeMPItemID())
	c.SendPacket(outpacket.NewPetConsumeSkillID())
	c.SendPacket(outpacket.NewPetConsumeUNKID())
	// CPet::OnActivated opcode=609 TODO
	// Show pet action
	c.SendPacket(outpacket.CPetOnActivated(true))

	// @Input: Keyboard and Mouse setting
	// CFiled::OnMouseMove opcode=529
	c.SendPacket(outpacket.NewMouseMove())
	// CQuickslotKeyMappedMan::OnInit opcode=471
	c.SendPacket(outpacket.NewCQuickslotKeyMappedMan(char.Keymap.QuickSlotKeys))

	// @UNK
	// CwvsContext opcode=441&399
	c.SendPacket(outpacket.NewUNK441())
	c.SendPacket(outpacket.NewUNK399(char.ID))

	// @ClaimSvr
	// CWvsContext::OnSetClaimSvrAvailableTime opcode=102
	c.SendPacket(outpacket.NewSetClaimSvrAvailableTime())
	// CWvsContext::OnClaimSvrStatusChanged opcode=103
	c.SendPacket(outpacket.NewClaimSvrStatusChanged())

	// @StarPlanet
	// CWvsContext::OnStarPlanetUserCount opcode=104
	c.SendPacket(outpacket.NewStarPlanetUserCount())
	// CField::OnStarPlanetBurningTimeInfo opcode=511
	c.SendPacket(outpacket.NewStarPlanetBurningTimeInfo())

	// @GrowthHelper
	// CUIContext::OnGrowthHelper opcode=1434
	// CGrowthHelperMan::SetCheckCount length=46
	c.SendPacket(outpacket.NewGrowthHelper(char.PosMap, 1))
	// CGrowthHelperMan::OnRecommendItemList length=174
	c.SendPacket(outpacket.NewGrowthHelper(char.PosMap, 3))

	// CUIContext::OnContentsMap opcode=1435
	// CContentsMapMan::OnReceiveFieldContentRewardData length=1534
	c.SendPacket(outpacket.NewContentsMap(char.PosMap))

	// @Admin: TODO GM effect?
	// CField::OnAdminResult opcode=457
	c.SendPacket(outpacket.NewAdminResult())

	// @Mob
	// CMobPool::OnMobCrcKeyChanged opcode=951 unused
	// c.SendPacket(outpacket.NewMobCrcKeyChanged())

	// @QuestRecords
	// CWvsContext_OnMessage opcode=95 TODO
	quest := &nx.QuestNX{
		ID: 1000,
	}
	c.SendPacket(outpacket.NewQuestRecordExMessage(quest))

	// @OtherPlayer: Broadcast new player enter field
	// CUserPool::OnUserEnterField opcode=531
	ips := field.GetIPs()
	for ip := range ips {
		clientTemp, ok := s.clientMap.Load(ip)
		if !ok {
			field.RemoveIP(ip)
			continue
		}
		client, ok := clientTemp.(network.ChannelClient)
		if !ok {
			slog.Error("Failed to assert network.ChannelClient")
			field.RemoveIP(ip)
			continue
		}
		client.SendPacket(outpacket.NewUserEnterField(client.GetChar(), client.GetInv(), client.GetSoc()))
	}

	// @NPC: Create npc entity
	for _, life := range field.MapNX.Lifes {
		// CNpcPool::NewNpcEnterField opcode=1012
		if life.IsNPC() {
			c.SendPacket(outpacket.NewNpcEnterField(&life))
			// CNpcPool::NpcChangeController opcode=1015
			c.SendPacket(outpacket.NewNpcChangeController(true, &life))
		}
	}

	// @EmotionItem
	// CUser::OnSetActiveEmoticonItem opcode=584
	c.SendPacket(outpacket.NewSetActiveEmoticonItem())

	// @Broadcast: Public screen messages
	// CWvsContext::OnBroadcastMsg opcode=135
	broadcast := channel.NewBroadcast("Welcome to goms")
	c.SendPacket(outpacket.NewBroadcastMsg(broadcast, s.index))

	// CField::OnSetQuickMoveInfo opcode=486 TODO
	c.SendPacket(outpacket.NewSetQuickMoveInfo())

	// CWvsContext::OnSessionValue opcode=166
	c.SendPacket(outpacket.NewSessionValue())

	// CWvsContext::OnInventoryOperation opcode=77 TODO
	c.SendPacket(outpacket.NewInventoryOperation(false, false, nil, maple.InvTypeCash, maple.ItemTypeBundle, nil))

	// CWvsContext::OnMonsterBattleSystemResult opcode=301 TODO
	c.SendPacket(outpacket.NewMonsterBattleSystemResult())

	// CField::OnMomentAreaOnOffAll opcode=503 TODO
	c.SendPacket(outpacket.NewMomentAreaOnOffAll())

	// CUserLocal::OnEnterFieldPsychicInfo opcode=852 TODO
	c.SendPacket(outpacket.NewEnterFieldPsychicInfo())

	// CWvsContext::OnTemporaryStatSet opcode=80 TODO
	c.SendPacket(outpacket.NewTemporaryStatSet())

	// CWvsContext::OnChangeSkillRecordResult opcode=84 TODO
	c.SendPacket(outpacket.NewChangeSkillRecordResult(false, false, false, char, 2))

	// CWvsContext::OnInventoryOperation opcode=77 TODO
	c.SendPacket(outpacket.NewInventoryOperation(false, false, nil, maple.InvTypeConsume, maple.ItemTypeBundle, nil))

	// CWvsContext::OnChangeSkillRecordResult opcode=84 TODO
	c.SendPacket(outpacket.NewChangeSkillRecordResult(false, true, false, char, 3))

	// CPet::OnActivated opcode=609 TODO
	c.SendPacket(outpacket.CPetOnActivated(true))

	// CWvsContext::OnTownPortal opcode=134 TODO
	c.SendPacket(outpacket.NewTownPortal(maple.EmptyPortalID, maple.EmptyPortalID))

	// CWvsContext::OnSetPotionDiscountRate opcode=153
	c.SendPacket(outpacket.NewSetPotionDiscountRate())

	// CWvsContext::OnSetBuyEquipExt opcode=185
	c.SendPacket(outpacket.NewSetBuyEquipExt())

	// CWvsContext::UNK411 opcode=411
	c.SendPacket(outpacket.NewUNK411())

	// CUserLocal::UNK883 opcode=883
	c.SendPacket(outpacket.NewUNK883())

	// CUserLocal::OnIsUniverse opcode=808
	c.SendPacket(outpacket.NewIsUniverse())

	// CUserLocal::OnClientResolution opcode=837
	c.SendPacket(outpacket.NewClientResolution())

	// CUserLocal::OnMonsterBattleCapture opcode=807
	c.SendPacket(outpacket.NewMonsterBattleCapture(2))

	// CUser::UNK608 opcode=608 TODO
	c.SendPacket(outpacket.NewUNK608()) // maybe effect switch?

	// @Stat
	// CWvsContext::OnForcedStatReset opcode=83 TODO
	c.SendPacket(outpacket.NewForcedStatReset())
	// CWvsContext::OnTemporaryStatReset opcode=81 TODO
	c.SendPacket(outpacket.NewTemporaryStatReset())
	// CWvsContext::OnTemporaryStatSet opcode=80
	c.SendPacket(outpacket.NewTemporaryStatSet())
	// CWvsContext::OnStatChanged opcode=79
	c.SendPacket(outpacket.NewStatChanged(false, maple.CS_None, char, 0, 1))
	c.SendPacket(outpacket.NewStatChanged(false, maple.CS_None, char, 0, 4))
	c.SendPacket(outpacket.NewStatChanged(false, maple.CS_None, char, 0, 5))

	// @Coupon: Types of hairstyles available in the current town
	// CWvsContext::OnHairStyleCoupon opcode=388
	c.SendPacket(outpacket.NewHairStyleCoupon())

	// @Expedtion
	// CWvsContext_OnExpedtionResult opcode=128
	c.SendPacket(outpacket.NewExpedtionResult(maple.ExpNoti_UNK))
}

// RecvClientResolution implements network.ChannelServer.
func (s *channelServer) RecvClientResolution(c network.ChannelClient, data []byte) {
	in := inpacket.NewClientResolution(data)
	screenWidth := 800
	screenHeight := 600
	switch in.ResolutionType {
	case 1:
		screenWidth = 800
		screenHeight = 600
	case 2:
		screenWidth = 1024
		screenHeight = 768
	case 3:
		screenWidth = 1366
		screenHeight = 768
	}
	c.SetResolution(screenWidth, screenHeight)
}

// RecvRequestInstanceTable implements network.ChannelServer.
func (s *channelServer) RecvRequestInstanceTable(c network.ChannelClient, data []byte) {
	in := inpacket.NewRequestInstanceTable(data)
	table := maple.NewInstanceTable(in.TableName, in.Col, in.Row)
	c.SendPacket(outpacket.NewResultInstanceTable(table))
}

// RecvSaveQuickslotKeyMap implements network.ChannelServer.
func (s *channelServer) RecvSaveQuickslotKeyMap(c network.ChannelClient, data []byte) {
	char := c.GetChar()
	if char == nil {
		return
	}
	in := inpacket.NewSaveQuickslotKeyMap(data)
	char.Keymap.QuickSlotKeys = in.Keycode
}

// RecvTalkToNpc implements network.ChannelServer.
func (s *channelServer) RecvTalkToNpc(c network.ChannelClient, data []byte) {
	in := inpacket.NewTalkToNpc(data)
	c.SendPacket(outpacket.NewChatMsg(maple.CT_Normal, fmt.Sprintf("Talk to NPC with objectID=%d&charPosX=%d&charPosY=%d", in.ObjectID, in.CharPosX, in.CharPosY)))
	scriptName := strconv.Itoa(int(in.ObjectID)) + ".js"
	scriptPath := path.Join(maple.NPCScriptDir, scriptName)
	_, program, err := util.ComplieProgramFromFile(scriptPath)
	if err != nil {
		slog.Error("Failed to load npc script", "err", err, "scriptName", scriptName)
		return
	}
	sm := &object.ScriptMessage{
		SpeakerTemplateID: in.ObjectID,
		MsgType:           maple.Say,
	}
	c.SetScriptMessage(sm)
	c.StartNPCScript(program)
}

// RecvTalkToNpcStep implements network.ChannelServer.
func (s *channelServer) RecvTalkToNpcStep(c network.ChannelClient, data []byte) {
	in := inpacket.NewTalkToNpcAction(data)
	c.ResumeNPCScript(in.Type, in.Mode, in.Selection)
}
