package outpacket

import (
	"goms/maple"
	"goms/mongodb/model/character"
	"goms/mongodb/model/inventory"
	"goms/mongodb/model/social"
	"goms/opcode"
)

// CUserPool::OnPacket list
// 1.CUserPool::OnUserEnterField
// 2.CUserPool::OnUserLeaveField
// 3.CUserPool::OnUserCommonPacket
// 3.1.CUser::OnPetPacket
// 3.2.CUser::OnDragonPacket
// 3.3.CUser::OnAndroidPacket
// 3.4.CUser::OnFoxManPacket
// 3.5.CUser::OnSkillPetPacket
// 4.CUserPool::OnUserRemotePacket
// 5.CUserLocal::OnPacket

// CUserPool::OnUserEnterField
func NewUserEnterField(char *character.Character, inv *inventory.Inventory, soc *social.Social) []byte {
	p := newOutPacket(opcode.CUserPool_OnUserEnterField)
	p.EncodeUint32(char.ID)
	CUserRemoteInit(&p, char, inv, soc)
	return p.buf
}

// Call by CUserPool::OnUserEnterField
// CUserRemote::Init
func CUserRemoteInit(p *outPacket, char *character.Character, inv *inventory.Inventory, social *social.Social) {
	p.EncodeByte(char.Stat.Level)          // m_nLevel
	p.EncodeLocalStr(char.Name)            // sCharacterName
	p.EncodeStr("")                        // m_sParentName, deprecated
	p.EncodeLocalStr(social.Guild.Name)    // m_sGuildName
	p.EncodeUint16(social.Guild.MarkBg)    // m_nGuildMarkBg
	p.EncodeByte(social.Guild.MarkBgColor) // m_nGuildMarkBgColor
	p.EncodeUint16(social.Guild.Mark)      // m_nGuildMark
	p.EncodeByte(social.Guild.MarkColor)   // m_nGuildMarkColor
	p.EncodeBool(char.Gender)
	p.EncodeBool(false) // m_nAccountGender
	// Reputation
	p.EncodeUint32(char.Stat.Pop) // m_nPopularity
	p.EncodeUint32(0)             // m_nFarmLevel
	p.EncodeUint32(0)             // m_nNameTagMark
	SecondaryStatEncodeForRemote(p)
	p.EncodeUint16(char.Job)       // m_nJobCode
	p.EncodeUint16(char.SubJob)    // m_nSubJobCode
	p.EncodeUint32(char.TotalCHUC) // m_nTotalCHUC
	AvatarLookEncode(p, char, inv) //  AvatarLook::AvatarLook
	if maple.IsZeroJob(char.Job) {
		AvatarLookEncode(p, nil, nil) // Additional appearance
	}
	sub_751FC0(p)     // UNK
	p.EncodeUint32(0) // m_dwDriverID
	p.EncodeUint32(0) // m_dwPassenserID
	sub_190CD20(p)
	// pAvatarHairEquip for CAvatar::ForcingAppearance?
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeStr("")
	p.EncodeStr("")
	p.EncodeUint16(0)
	p.EncodeUint16(0)
	p.EncodeByte(0)
	p.EncodeUint16(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	// TextChairInfo start
	portableChairID := 0
	p.EncodeUint32(uint32(portableChairID)) // m_nPortableChairID
	if portableChairID > 0 {
		p.EncodeStr("") // m_sPortableChairMsg
	}
	// TextChairInfo end
	// TowerChairInfo start
	count := 0
	p.EncodeUint32(uint32(count)) // m_lTowerChairIDList
	for i := 0; i < count; i++ {
		p.EncodeUint32(0) // m_lTowerChairID
	}
	// TowerChairInfo end
	condition := false
	p.EncodeBool(condition)
	p.EncodeUint16(0)   // posX
	p.EncodeUint16(0)   // posY
	p.EncodeByte(0)     // m_nMoveAction
	p.EncodeUint16(0)   // dwSN use for CWvsPhysicalSpace2D::GetFoothold
	p.EncodeByte(0)     // related to Kaiser skill 60000219
	p.EncodeBool(false) // isCustomChair
	// TamingMob
	p.EncodeUint32(char.TamingMob.Level)   // m_nTamingMobLevel
	p.EncodeUint32(char.TamingMob.Exp)     // m_nTamingMobExp
	p.EncodeUint32(char.TamingMob.Fatigue) // m_nTamingMobFatigue
	// MiniRoom
	var miniRoomType byte = 0
	p.EncodeByte(miniRoomType) // m_nMiniRoomType
	if miniRoomType > 0 {
		p.EncodeUint32(0)   // m_dwMiniRoomSN
		p.EncodeStr("")     // m_sMiniRoomTitle
		p.EncodeBool(false) // m_bPrivate
		p.EncodeByte(0)     // m_nGameKind
		p.EncodeByte(0)     // m_nCurUsers
		p.EncodeByte(0)     // m_nMaxUsers
		p.EncodeBool(false) // m_bGameOn
	}
	// ADBoard
	adBoardRemote := false // m_bADBoardRemote
	p.EncodeBool(adBoardRemote)
	if adBoardRemote {
		p.EncodeStr("")
	}
	// Couple Record
	isCouple := false
	p.EncodeBool(isCouple)
	if isCouple {
		// CUserPool::OnCoupleRecordAdd
		p.EncodeUint64(0) // m_liCoupleItemSN
		p.EncodeUint64(0) // m_liPairItemSN
		p.EncodeUint32(0) // nItemID
	}
	// Friend Record
	isFriend := false
	p.EncodeBool(isFriend)
	if isFriend {
		// CUserPool::OnFriendRecordAdd
		p.EncodeUint64(0) // m_liFriendshipItemSN
		p.EncodeUint64(0) // m_liFriendshipPairItemSN
		p.EncodeUint32(0) // nItemID
	}
	// Marriage Record
	isMarried := false
	p.EncodeBool(isMarried)
	if isMarried {
		// CUserPool::OnMarriageRecordAdd
		p.EncodeUint32(0) // m_dwMarriageCharacterID
		p.EncodeUint32(0) // m_dwMarriagePairCharacterID
		p.EncodeUint32(0) // nRingID
	}
	// UNK Record
	condition = false
	p.EncodeBool(condition)
	if condition {
		count = 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			p.EncodeUint32(0)
		}
	}
	var flag byte = 0
	p.EncodeByte(flag)
	// flag&1 !=0 use for CUser::LoadReincarnationEffect
	// flag&2 !=0 use for CDragon::CreateEffect
	if flag&0x8 != 0 {
		p.EncodeUint32(0) // m_tDelayedPvPEffectTime
		// call CUser::LoadPvPRageEffect
	} else if flag&0x10 != 0 {
		p.EncodeUint32(0) // m_tDelayedPvPEffectTime
		// call CUser::LoadPvPChampionEffect
	}
	if flag&0x20 != 0 {
		p.EncodeUint32(0) // m_tHitPeriodRemain_Revive
	}
	p.EncodeUint32(0) // m_nEvanDragonGlide_Riding
	if maple.IsKaiserJob(char.Job) {
		p.EncodeUint32(0)   // m_nKaiserMorphRotateHueExtern
		p.EncodeUint32(0)   // m_nKaiserMorphRotateHueInnner
		p.EncodeBool(false) // m_bKaiserMorphPrimiumBlack
	}
	p.EncodeUint32(0) // nSkillID use for CUser::SetMakingMeisterSkillEff
	for i := 0; i < 5; i++ {
		p.EncodeByte(0) // m_aActiveEventNameTag
	}
	itemID := 0
	p.EncodeUint32(uint32(itemID)) //  nItemID use for CUser::SetCustomizeEffect
	if itemID > 0 {
		p.EncodeStr("")
	}
	isSoulEffect := false
	p.EncodeBool(isSoulEffect)
	condition = false
	p.EncodeBool(condition)
	if condition {
		condition = false
		if condition {
			// CUser::SetFlareBlink
			p.EncodeUint32(0) // nSLV use for
			p.EncodeUint32(0)
			p.EncodeUint16(0) // ptPos.x
			p.EncodeUint16(0) // ptPos.y
		}
	}
	CUserStarPlanetRankEncode(p)
	CUserEncodeStarPlanetTrendShopLook(p)
	CUserEncodeTextEquipInfo(p)
	CUserEncodeFreezeHotEventInfo(p, 0, 0)
	CUserEncodeEventBestFriendInfo(p)
	p.EncodeBool(false) // bOnOff use for CUserRemote::OnKinesisPsychicEnergyShieldEffect
	p.EncodeBool(false) // m_bBeastFormWingOnOff
	p.EncodeUint32(0)   // nMeso use for CUser::SetMesoChairCount
	p.EncodeUint32(0)
	sub_1851B20(p)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	p.EncodeUint32(0)
	CUserEncodeSundayMaple(p)
	sub_1A92CB0(p)
}

// Call by CUserRemote::Init
// SecondaryStat::DecodeForRemote
func SecondaryStatEncodeForRemote(p *outPacket) {
	p.Fill(68) // CInPacket::DecodeBuffer(iPacketa, &uFlagTemp, 0x44u)
	// TODO
	p.EncodeByte(0) // nDefenseAtt
	p.EncodeByte(0) // nDefenseState
	p.EncodeByte(0) // nPVPDamage
	SecondaryStatStopForceAtomEncode(p)
	p.EncodeUint32(0)
}

// Call by SecondaryStat::DecodeForRemote
// SecondaryStat::StopForceAtom::Decode
func SecondaryStatStopForceAtomEncode(p *outPacket) {
	p.EncodeUint32(0) // nIdx
	p.EncodeUint32(0) // nCount
	p.EncodeUint32(0) // nWeaponID
	count := 0
	p.EncodeUint32(uint32(count))
	for i := 0; i < count; i++ {
		p.EncodeUint32(0) // p_aAngleInfo
	}
}

func sub_751FC0(p *outPacket) {
	// 00 00 00 00 FF 00 00 00 00 FF
	for i := 0; i < 2; i++ {
		p.EncodeUint32(0)
		count := 0
		for i := 0; i < count; i++ {
			p.EncodeByte(0)
			p.EncodeUint32(0)
		}
		p.EncodeByte(0xFF)
	}
}

func sub_190CD20(p *outPacket) {
	p.EncodeUint32(0) // pfh?
	p.EncodeUint32(0) // nMonkeyEffectItemID?
	count := 0
	p.EncodeUint32(uint32(count)) // nKaiserTailID?
	for i := 0; i < count; i++ {
		p.EncodeUint32(0)
		p.EncodeUint32(0)
	}
}

func sub_1851B20(p *outPacket) {
	count := 0
	p.EncodeUint32(uint32(count))
	for i := 0; i < count; i++ {
		p.EncodeUint32(0)
		p.EncodeUint32(0)
	}
}

func sub_1A92CB0(p *outPacket) {
	count := 0
	p.EncodeUint32(uint32(count))
	for i := 0; i < count; i++ {
		p.EncodeUint32(0)
		p.EncodeUint32(0)
		p.EncodeUint32(0)
	}
}

// CUser::OnSetActiveEmoticonItem
func NewSetActiveEmoticonItem() []byte {
	p := newOutPacket(opcode.CUser_OnSetActiveEmoticonItem)
	p.EncodeUint32(0) // m_nActiveEmoticonItemID
	p.EncodeUint32(0) // m_nActiveEmoticonItemPos
	p.EncodeInt32(0)  // unk
	return p.buf
}

// CPet::OnActivated v138 new
func CPetOnActivated(isActive bool) []byte {
	p := newOutPacket(opcode.CPet_OnActivated)
	p.EncodeUint32(0) // ownerID
	p.EncodeInt32(0)  // idx?
	p.EncodeBool(isActive)
	if isActive {
		p.EncodeBool(true)   // bInit
		p.EncodeUint32(0)    // templateID
		p.EncodeLocalStr("") // petName
		p.EncodeUint64(0)    // itemID?
		p.EncodeInt16(0)     // posX
		p.EncodeInt16(0)     // posY
		p.EncodeByte(0)      // moveAction
		p.EncodeInt16(0)     // Fh
		p.EncodeInt32(0)     // Hue
		p.EncodeInt16(0)     // giant rate
		p.EncodeBool(false)  // bTransformed
		p.EncodeBool(false)  // bReinforced
	} else {
		p.EncodeByte(0) // removed reason
	}
	return p.buf
}

// CUserLocal::OnEnterFieldPsychicInfo
func NewEnterFieldPsychicInfo() []byte {
	p := newOutPacket(opcode.CUserLocal_OnEnterFieldPsychicInfo)
	p.EncodeBool(false)
	return p.buf
}

// CUserLocal::UNK883
func NewUNK883() []byte {
	p := newOutPacket(opcode.CUserLocal_UNK883)
	p.EncodeByte(2)
	return p.buf
}

// CUserLocal::OnIsUniverse
func NewIsUniverse() []byte {
	p := newOutPacket(opcode.CUserLocal_OnIsUniverse)
	p.EncodeBool(false) // m_bUniverse
	return p.buf
}

// CUserLocal::OnClientResolution
func NewClientResolution() []byte {
	p := newOutPacket(opcode.CUserLocal_OnClientResolution)
	return p.buf
}

// CUserLocal::OnMonsterBattleCapture
func NewMonsterBattleCapture(captureType int16) []byte {
	p := newOutPacket(opcode.CUserLocal_OnMonsterBattleCapture)
	p.EncodeInt16(captureType)
	if captureType > 0 {
		condition := captureType - 1
		if condition == 1 {
			p.EncodeInt16(0)    // m_nCaptureGauge
			p.EncodeInt32(0)    // m_tCaptureDelay 52 B3 0A 00？
			p.EncodeBool(false) // m_bDecCaptureGauge
		} else {
			p.EncodeInt16(0) // m_nCaptureGauge
			p.EncodeBool(true)
		}
	} else {
		p.EncodeUint32(0)   // dwMobID
		p.EncodeByte(0)     // m_nCaptureLevel
		p.EncodeBool(false) // bSuccess
	}
	return p.buf
}

// CUser::UNK608
func NewUNK608() []byte {
	p := newOutPacket(opcode.CUser_UNK608)
	count := 0
	p.EncodeUint32(uint32(count))
	for i := 0; i < count; i++ {
		p.EncodeInt32(0)
	}
	p.EncodeBool(false)
	return p.buf
}

// CUserLocal::OnChatMsg
func NewChatMsg(chatType maple.ChatType, msg string) []byte {
	p := newOutPacket(opcode.CUserLocal_OnChatMsg)
	p.EncodeInt16(int16(chatType)) // nType
	p.EncodeLocalStr(msg)          // sChat
	return p.buf
}
