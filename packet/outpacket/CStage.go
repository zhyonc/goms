package outpacket

import (
	"goms/maple"
	"goms/mongodb/model/character"
	"goms/mongodb/model/inventory"
	"goms/network/server/channel"
	"goms/nx"
	"goms/opcode"
	"time"
)

// CStage::OnSetField
func NewSetField(channelIndex uint8, bCharacterData bool, posMap *nx.MapNX, notifiter *channel.Notifiter, damageSeed [3]uint32,
	flag maple.CharFlag, char *character.Character, inv *inventory.Inventory, maplePoint, cahsPoint uint32,
	usingBuffProtector bool, whiteFadeInOut bool, mobStatAdjustRate uint32,
	canNotifyAnnouncedQuest bool, stackEventGauge uint32) []byte {
	p := newOutPacket(opcode.CStage_OnSetField)
	count := 1
	p.EncodeUint16(uint16(count))
	for i := 0; i < count; i++ {
		p.EncodeUint32(1)
		p.EncodeUint32(0)
	}
	p.EncodeUint32(uint32(channelIndex)) // m_nChannelID start at 0
	p.EncodeBool(false)                  // m_bDev
	p.EncodeUint32(0)                    // m_dwOldDriverID
	// Logging Into Handler (1) Changing Map (2) bPopupDlg
	if bCharacterData {
		p.EncodeByte(1)
	} else {
		p.EncodeByte(2)
	}
	p.EncodeByte(0) // unk
	if !bCharacterData {
		p.EncodeInt32(posMap.Info.FieldType)
	} else {
		p.EncodeUint32(0)
	}
	p.EncodeInt32(posMap.GetWidth())  // nFieldWidth
	p.EncodeInt32(posMap.GetHeight()) // nFieldHeight
	p.EncodeBool(bCharacterData)      // bCharacterData
	notifiterCheck := len(notifiter.Msgs)
	p.EncodeUint16(uint16(notifiterCheck)) // nNotifierCheck (used for block reason)
	if notifiterCheck > 0 {
		p.EncodeStr(notifiter.BlockReason)
		for _, msg := range notifiter.Msgs {
			p.EncodeStr(msg)
		}
	}
	if bCharacterData {
		// CalcDamage::SetSeed
		p.EncodeUint32(damageSeed[0])
		p.EncodeUint32(damageSeed[1])
		p.EncodeUint32(damageSeed[2])
		CharacterDataEncode(&p, flag, char, inv, maplePoint, cahsPoint)
		CWvsContextLogoutEvent(&p)
	} else {
		p.EncodeBool(usingBuffProtector) // m_bUsingBuffProtector
		p.EncodeUint32(char.PosMap)      // p_dwPosMap
		p.EncodeByte(char.Portal)        // nPortal
		p.EncodeUint32(char.Stat.HP)     // characterStat.nHP
	}
	p.EncodeBool(whiteFadeInOut) // m_bWhiteFadeInOut
	p.EncodeBool(false)          // set true overlapping screen animation
	p.EncodeFT(time.Now())       // ftServer
	p.EncodeUint32(100)          // nMobStatAdjustRate
	isCustomField := false
	p.EncodeBool(isCustomField)
	if isCustomField {
		CFieldCustomEncode(&p) // CFieldCustom::Decode
	}
	p.EncodeBool(false)                   // CWvsContext::OnInitPvPStat
	p.EncodeBool(canNotifyAnnouncedQuest) // bCanNotifyAnnouncedQuest
	p.EncodeBool(true)                    // bCanUseFamiliarSystem
	buf := FamiliarSystemEncode(inv, bCharacterData)
	p.EncodeUint32(uint32(len(buf)))
	p.EncodeBuffer(buf)
	p.EncodeBool(stackEventGauge > 0)
	if stackEventGauge > 0 {
		p.EncodeUint32(stackEventGauge) //  CField::DrawStackEventGauge
	}
	if bCharacterData && maple.IsBanBanBaseField(char.PosMap) {
		notifierMessage := 0 // sNotifierMessage
		p.EncodeByte(byte(notifierMessage))
		for i := 0; i < notifierMessage; i++ {
			p.EncodeStr("") // sMsg2
		}
	}
	// else if maple.IsBanBanBaseField(field.ID) {
	// 	notifierMessage := 0 // sNotifierMessage
	// 	p.EncodeByte(byte(notifierMessage))
	// 	for i := 0; i < notifierMessage; i++ {
	// 		p.EncodeStr("") // sMsg2
	// 	}
	// }
	CUserStarPlanetRankEncode(&p)                        // TODO CUser::StarPlanetRank::Decode
	CWvsContextEncodeStarPlanetRoundInfo(&p)             // TODO CWvsContext::DecodeStarPlanetRoundInfo
	CUserEncodeTextEquipInfo(&p)                         // TODO CUser::DecodeTextEquipInfo
	CUserEncodeFreezeHotEventInfo(&p, 0, char.AccountID) // TODO CUser::DecodeFreezeHotEventInfo
	CUserEncodeEventBestFriendInfo(&p)                   // TODO CUser::DecodeEventBestFriendInfo
	CUserEncodeSundayMaple(&p)                           // TODO CUser::DecodeSundayMaple
	return p.buf
}
