package outpacket

import (
	"goms/maple"
	"goms/mongodb/model/character"
	"goms/mongodb/model/social"
	"goms/network/server/channel"
	"goms/nx"
	"goms/opcode"
	"goms/util"
)

// CWvsContext::SetAccountInfo
func NewSetAccountInfo() []byte {
	p := newOutPacket(opcode.CWvsContext_OnSetAccountInfo)
	p.EncodeUint64(0) // m_llNexonOID
	p.EncodeStr("")   // m_sNexonClubID
	return p.buf
}

// CWvsContext::OnEventNameTag
func NewEventNameTag() []byte {
	p := newOutPacket(opcode.CWvsContext_OnEventNameTag)
	for i := 0; i < 5; i++ {
		p.EncodeStr("")
		p.EncodeByte(0xFF)
	}
	return p.buf
}

// CWvsContext::OnRequestEventList
func NewRequestEventList() []byte {
	p := newOutPacket(opcode.CWvsContext_OnRequestEventList)
	p.EncodeInt32(0) // EventInfo::ms_nDefaultLevelLimit
	condition := false
	p.EncodeBool(condition)
	if !condition {
		CWvsContextOnLiveEvent(&p)
		return p.buf
	}
	if CWvsContextMakeEventListImg(&p) {
		CWvsContextOnLiveEvent(&p)
	}
	return p.buf
}

// CWvsContext::OnHourChanged
func NewHourChanged(day int, hour int) []byte {
	p := newOutPacket(opcode.CWvsContext_OnHourChanged)
	p.EncodeUint16(uint16(day))  // wDayOfWeek
	p.EncodeUint16(uint16(hour)) // m_bPassiveMode
	return p.buf
}

// CWvsContext::OnSetTamingMobInfo
func NewSetTamingMobInfo(char *character.Character, isLevelUp bool) []byte {
	p := newOutPacket(opcode.CWvsContext_OnSetTamingMobInfo)
	p.EncodeUint32(char.ID)                // dwCharacterID
	p.EncodeUint32(char.TamingMob.Level)   // m_nTamingMobLevel
	p.EncodeUint32(char.TamingMob.Exp)     // m_nTamingMobExp
	p.EncodeUint32(char.TamingMob.Fatigue) // m_nTamingMobFatigue
	p.EncodeBool(isLevelUp)
	return p.buf
}

// CWvsContext::UNK441
func NewUNK441() []byte {
	p := newOutPacket(opcode.CWvsContext_UNK441)
	p.EncodeByte(0) // range is 0-5
	return p.buf
}

// CWvsContext::UNK399
func NewUNK399(characterID uint32) []byte {
	p := newOutPacket(opcode.CWvsContext_UNK399)
	p.EncodeUint32(characterID)
	if characterID > 0 {
		p.EncodeUint32(0)
	}
	return p.buf
}

// CWvsContext::OnSetClaimSvrAvailableTime
func NewSetClaimSvrAvailableTime() []byte {
	p := newOutPacket(opcode.CWvsContext_OnSetClaimSvrAvailableTime)
	p.EncodeByte(0) // m_nClaimSvrOpenTime
	p.EncodeByte(0) // m_nClaimSvrCloseTime
	return p.buf
}

// CWvsContext::OnClaimSvrStatusChanged
func NewClaimSvrStatusChanged() []byte {
	p := newOutPacket(opcode.CWvsContext_OnClaimSvrStatusChanged)
	p.EncodeBool(true) // m_bClaimSvrConnected
	return p.buf
}

// CWvsContext::OnStarPlanetUserCount
func NewStarPlanetUserCount() []byte {
	p := newOutPacket(opcode.CWvsContext_OnStarPlanetUserCount)
	count := 4
	p.EncodeUint32(uint32(count))
	for i := 0; i < count; i++ {
		// p_m_mUserCountByServerIdx
		p.EncodeUint32(uint32(i + 3)) // key 3,4,5,6?
		p.EncodeUint32(0)             // value
	}
	return p.buf
}

// CWvsContext::OnResultInstanceTable
func NewResultInstanceTable(table maple.InstanceTable) []byte {
	p := newOutPacket(opcode.CWvsContext_OnResultInstanceTable)
	p.EncodeStr(table.Name)         // sTableName
	p.EncodeInt32(table.Col)        // type?
	p.EncodeInt32(table.Row)        // subType?
	p.EncodeBool(true)              // bRightResult
	p.EncodeInt32(table.GetValue()) // nValue
	return p.buf
}

// CWvsContext::OnBroadcastMsg
func NewBroadcastMsg(broadcast channel.Broadcast, channelIndex uint8) []byte {
	p := newOutPacket(opcode.CWvsContext_OnBroadcastMsg)
	p.EncodeByte(byte(broadcast.Type))
	if broadcast.Type == maple.SlideNotice {
		p.EncodeBool(true) // bSlideNoticeExist
	}
	p.EncodeLocalStr(broadcast.Msg)
	switch broadcast.Type {
	case maple.Megaphone,
		maple.MegaphoneNoMessage:
		p.EncodeByte(channelIndex)          // m_pStr
		p.EncodeBool(broadcast.WhisperIcon) // bWhisperIcon
	case maple.ItemMegaphone:
		p.EncodeByte(channelIndex)          // m_pStr
		p.EncodeBool(broadcast.WhisperIcon) // bWhisperIcon
		if broadcast.Item.ItemID > 0 {
			p.EncodeBool(true)
			GWItemSlotBaseEncode(&p, maple.ItemTypeBundle, broadcast.Item)
		} else {
			p.EncodeBool(false)
		}
	case maple.ItemMegaphoneNoItem:
		p.EncodeByte(channelIndex) // m_pStr

	case maple.TripleMegaphone:
		p.EncodeByte(byte(len(broadcast.NextMsg)))
		for _, msg := range broadcast.NextMsg {
			p.EncodeLocalStr(msg)
		}
		p.EncodeByte(channelIndex)
		p.EncodeBool(broadcast.WhisperIcon)
	case maple.BlowWeather:
		p.EncodeUint32(broadcast.Item.ItemID)
	case maple.BalloonMessage:
		p.EncodeUint32(broadcast.Item.ItemID)
		p.EncodeUint32(broadcast.Timeout)
		condition := false
		p.EncodeBool(condition) // hasPackedChararacterLook?
		if condition {
			p.Fill(120) //  CInPacket::DecodeBuffer(0x78u)
		}
	case maple.WhiteYellow_ItemInfo:
		if broadcast.Item.ItemID == 0 {
			p.EncodeBool(false)
		} else {
			p.EncodeBool(true)
			GWItemSlotBaseEncode(&p, maple.ItemTypeBundle, broadcast.Item)
		}
	case maple.Yellow,
		maple.Yellow2:
		GWItemSlotBaseEncode(&p, maple.ItemTypeBundle, broadcast.Item)
	case maple.BlueChatItemInfo,
		maple.BlueChatItemInfo2:
		p.EncodeUint32(broadcast.Item.ItemID)
		GWItemSlotBaseEncode(&p, maple.ItemTypeBundle, broadcast.Item)
	case maple.GMErrorMessage:
		p.EncodeUint32(broadcast.NPCID)
	case maple.YellowChatFiledItemInfo:
		p.EncodeUint32(broadcast.Item.ItemID)
		p.EncodeBool(broadcast.Item.ItemID > 0)
		if broadcast.Item.ItemID > 0 {
			GWItemSlotBaseEncode(&p, maple.ItemTypeBundle, broadcast.Item)
		}
	case maple.TryRegisterAutoStartQuest:
		p.EncodeUint32(broadcast.QuestID)
		p.EncodeUint32(broadcast.Timeout)
	case maple.TryRegisterAutoStartQuest_NoAnnouncement:
		p.EncodeUint32(broadcast.QuestID)
	case maple.RedWithChannelInfo:
		p.EncodeUint32(broadcast.CharID)
	case maple.PopUpNotice:
		p.EncodeUint32(broadcast.Width)
		p.EncodeUint32(broadcast.Height)
	}
	return p.buf
}

// CWvsContext::OnPartyResult
func NewPartyResult(partyType maple.PartyType, soc *social.Social) []byte {
	p := newOutPacket(opcode.CWvsContext_OnPartyResult)
	p.EncodeByte(byte(partyType))
	switch partyType {
	case maple.PartyReq_InviteParty:
		p.EncodeUint32(0) // m_nPartyID
		p.EncodeStr("")   // sApplierName
		p.EncodeInt32(0)  // dwCID
		p.EncodeInt32(0)  // nGrade
		p.EncodeInt32(0)  // nCharacterJobCode.m_bLoopback
		p.EncodeByte(0)   // nSkillID
		condition := false
		p.EncodeBool(condition)
		if condition {
			p.EncodeByte(0) // nResult
		}
	case maple.PartyReq_InviteIntrusion:
		p.EncodeUint32(0) // m_nPartyID
		p.EncodeStr("")   // sApplierName
		p.EncodeInt32(0)  // nGrade
		p.EncodeInt32(0)  // nCharacterJobCode.m_bLoopbac
		p.EncodeInt32(0)  // nSkillID
	case maple.PartyReq_ApplyParty:
		p.EncodeUint32(0) // m_nPartyID
		p.EncodeStr("")   // sApplierName
		p.EncodeInt32(0)  // nLevel
		p.EncodeInt32(0)  // nCharacterJobCode.m_bLoopbac
		p.EncodeInt32(0)  // nSkillID
	case maple.PartyRes_LoadParty_Done:
		p.EncodeUint32(0)        // m_nPartyID set 0 means no team is created
		PartyDataEncode(&p, soc) // PARTYDATA::Decode
	case maple.PartyRes_CreateNewParty_Done:
		p.EncodeUint32(0) // m_nPartyID
		p.EncodeInt32(0)  // dwTownID
		p.EncodeInt32(0)  // dwFieldID
		p.EncodeInt32(0)  // nGrade
		p.EncodeInt16(0)  // pt.x
		p.EncodeInt16(0)  // pt.y
		p.EncodeByte(0)   // dwCID
		p.EncodeByte(0)   // bAppliable
		p.EncodeStr("")   // sPartyName
	case maple.PartyRes_WithdrawParty_Done:
		p.EncodeUint32(0)      // m_nPartyID
		p.EncodeInt32(0)       // member char id
		isPartyExists := false // bPartyExists
		p.EncodeBool(isPartyExists)
		if isPartyExists {
			p.EncodeByte(0) // bExpelled?
			p.EncodeStr("") // member sCharacterName
			PartyDataEncode(&p, soc)
		}
	case maple.PartyRes_JoinParty_Done:
		p.EncodeUint32(0) // m_nPartyID
		p.EncodeStr("")   // sJoinerName
		p.EncodeByte(0)   // v138 new
		p.EncodeInt32(0)  // v138 new
		PartyDataEncode(&p, soc)
	case maple.PartyRes_JoinParty_Done2:
		p.EncodeByte(0)  // v138 new
		p.EncodeInt32(0) // v138 new
	case maple.PartyRes_InviteParty_Sent,
		maple.PartyRes_InviteIntrusion_Sent:
		p.EncodeStr("") // CUtilDlg::Notice sMsg
	case maple.PartyRes_ChangePartyBoss_Done:
		p.EncodeUint32(0) // CWvsContext::FindUser dwMemberID
		p.EncodeBool(false)
	case maple.PartyRes_UserMigration:
		p.EncodeUint32(0) // m_nPartyID
		PartyDataEncode(&p, soc)
	case maple.PartyRes_ChangeLevelOrJob:
		p.EncodeInt32(0) // CWvsContext::FindUser dwMemberID
		p.EncodeInt32(0) // m_party.anLevel
		p.EncodeInt32(0) // m_party.anJob
	case maple.PartyRes_UpdateShutdownStatus:
		p.EncodeUint32(0) // CWvsContext::FindUser dwMemberID
		p.EncodeByte(0)   //  m_party.abAccountShutdown
	case maple.PartyRes_UNK62:
		p.EncodeUint32(0) // v138 new
		p.EncodeByte(0)   // v138 new
	case maple.PartyRes_SetAppliable:
		p.EncodeBool(false) // bAppliable
	case maple.PartyRes_SuccessToSelectPQReward:
		p.EncodeUint32(0) // m_dwCharacterID
		p.EncodeStr("")   // nGrade
		p.EncodeByte(0)   // nSelectedIdx
	case maple.PartyRes_FailToSelectPQReward:
		p.EncodeByte(0) // nRetCode
	case maple.PartyRes_ApplyParty_Sent:
		p.EncodeStr("") // CUtilDlg::Notice sMsg
	case maple.PartyRes_FoundPossibleMember,
		maple.PartyRes_FoundPossibleParty:
		p.EncodeStr("")   // sApplierName
		fadeWndCount := 1 // CWvsContext::GetFadeWndCount
		if fadeWndCount < 3 {
			p.EncodeInt32(0) // nGrade
			p.EncodeInt32(0) // dwCID
			p.EncodeInt32(0) // nResult
			condition := 78
			if condition != 78 {
				p.EncodeInt32(0) // nPartyID
			}
		}
	case maple.PartyRes_PartySettingDone:
		p.EncodeBool(false) // bAppliable
		p.EncodeStr("")     // nGrade
	case maple.PartyRes_Load_StarGrade_Result:
		count := 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			p.EncodeInt32(0) // dwCID
			p.EncodeInt32(0) // nGrade
		}
	case maple.PartyRes_Load_StarGrade_Result2:
		p.EncodeUint32(0) // m_nPartyID
		count := 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			p.EncodeInt32(0) // dwCID
			p.EncodeInt32(0) // nGrade
		}
	case maple.PartyRes_Member_Rename:
		p.EncodeUint32(0) // m_nPartyID
	case maple.PartyRes_UNK84:
		p.EncodeByte(0) // v138 new
	case maple.PartyRes_UNK85:
		p.EncodeByte(0)  // aTownPortal
		p.EncodeInt32(0) // dwTownID
		p.EncodeInt32(0) // dwFieldID
		p.EncodeInt32(0) // nSkillId
		p.EncodeInt16(0) // pt.x
		p.EncodeInt16(0) // pt.y
	case maple.AdverNoti_GetAll:
		p.EncodeUint32(0) // dwValue
		p.EncodeStr("")   // sApplierName
		p.EncodeInt32(0)  // nLevel
		p.EncodeInt32(0)  // dwCID
		p.EncodeInt32(0)  // nGrade
	}
	return p.buf
}

// CWvsContext::OnSessionValue
func NewSessionValue() []byte {
	p := newOutPacket(opcode.CWvsContext_OnSessionValue)
	p.EncodeStr("kill_count") // sKey
	p.EncodeStr("0")          // sVal
	return p.buf
}

// CWvsContext::OnInventoryOperation
func NewInventoryOperation(isSetExclRequestSent, isNotRemoveAddInfo bool, invOps []maple.InvOps, invType maple.InvType, itemType maple.ItemType, item any) []byte {
	p := newOutPacket(opcode.CWvsContext_OnInventoryOperation)
	p.EncodeBool(isSetExclRequestSent) // bSetExclRequestSent
	p.EncodeByte(byte(len(invOps)))    // nCount
	p.EncodeBool(isNotRemoveAddInfo)   // bNotRemoveAddInfo
	for _, ops := range invOps {
		p.EncodeByte(byte(ops))
		p.EncodeByte(byte(invType)) // inv type
		p.EncodeInt16(0)            // old nPos
		switch ops {
		case maple.InvOpsAdd:
			GWItemSlotBaseEncode(&p, itemType, item)
		case maple.InvOpsUpdateNumber:
			p.EncodeInt16(0) // Number
		case maple.InvOpsMove:
			p.EncodeInt16(0) // new BagIndex
		case maple.InvOpsRemove:
			// TODO
		case maple.InvOpsItemExp:
			p.EncodeInt64(0) // exp
		case maple.InvOpsUpdateBagIndex:
			p.EncodeInt32(0) // BagIndex
		case maple.InvOpsUpdateBagNumber:
			p.EncodeInt16(0) // new BagIndex
		case maple.InvOpsBagToBag:
			p.EncodeInt16(0) // Number
		case maple.InvOpsBagNewItem:
			GWItemSlotBaseEncode(&p, itemType, item)
		}
	}
	return p.buf
}

// CWvsContext::OnMonsterBattleSystemResult
func NewMonsterBattleSystemResult() []byte {
	p := newOutPacket(opcode.CWvsContext_OnMonsterBattleSystemResult)
	condition := 0x32
	p.EncodeByte(byte(condition))
	switch condition {
	case 0x32:
		// CMonsterBattleSystem::SetBlockByAdminPacket
		p.EncodeBool(true) // m_bBlockByAdmin
	}
	return p.buf
}

// CWvsContext::OnForcedStatReset
func NewForcedStatReset() []byte {
	p := newOutPacket(opcode.CWvsContext_OnForcedStatReset)
	// ForcedStat::Clear
	return p.buf
}

// CWvsContext::OnForcedStatSet
func NewForcedStatSet(mask maple.ForcedStat) []byte {
	p := newOutPacket(opcode.CWvsContext_OnForcedStatSet)
	ForcedStatEncode(&p, mask)
	return p.buf
}

// Call by CWvsContext::OnForcedStatSet
// ForcedStat::Decode
func ForcedStatEncode(p *outPacket, mask maple.ForcedStat) {
	p.EncodeUint32(uint32(mask))
	if mask&maple.FS_STR != 0 {
		p.EncodeUint16(0)
	}
	if mask&maple.FS_DEX != 0 {
		p.EncodeUint16(0)
	}
	if mask&maple.FS_INT != 0 {
		p.EncodeUint16(0)
	}
	if mask&maple.FS_LUK != 0 {
		p.EncodeUint16(0)
	}
	if mask&maple.FS_PAD != 0 {
		p.EncodeUint16(0)
	}
	if mask&maple.FS_MAD != 0 {
		p.EncodeUint16(0)
	}
	if mask&maple.FS_PDD != 0 {
		p.EncodeUint16(0)
	}
	if mask&maple.FS_MDD != 0 {
		p.EncodeUint16(0)
	}
	if mask&maple.FS_ACC != 0 {
		p.EncodeUint16(0)
	}
	if mask&maple.FS_EVA != 0 {
		p.EncodeUint16(0)
	}
	if mask&maple.FS_Speed != 0 {
		p.EncodeByte(0)
	}
	if mask&maple.FS_Jump != 0 {
		p.EncodeByte(0)
	}
	if mask&maple.FS_SpeedMax != 0 {
		p.EncodeByte(0)
	}
	if mask&maple.FS_OptOff != 0 {
		p.EncodeByte(0)
	}
	if mask&maple.FS_AddMHP != 0 {
		p.EncodeUint32(0)
	}
}

// CWvsContext::OnTemporaryStatReset
func NewTemporaryStatReset() []byte {
	p := newOutPacket(opcode.CWvsContext_OnTemporaryStatReset)
	p.Fill(68)
	return p.buf
}

// CWvsContext::OnTemporaryStatSet
func NewTemporaryStatSet() []byte {
	p := newOutPacket(opcode.CWvsContext_OnTemporaryStatSet)
	p.Fill(86)
	return p.buf
}

// CWvsContext::OnStatChanged
func NewStatChanged(isSetExclRequestSent bool, mask maple.CharStat, char *character.Character, money uint64, sn byte) []byte {
	p := newOutPacket(opcode.CWvsContext_OnStatChanged)
	p.EncodeBool(isSetExclRequestSent)
	GWCharacterStatEncodeChangeStat(&p, mask, char, money)
	GWCharacterStatEncodeChangeMixHairStat(&p, char)
	p.EncodeBool(sn > 0)
	if sn > 0 {
		// CUserLocal::SetSecondaryStatChangedPoint
		p.EncodeByte(sn) // bSN?
	}
	isSetBattleRecoveryInfo := false
	p.EncodeBool(isSetBattleRecoveryInfo)
	if isSetBattleRecoveryInfo {
		// CBattleRecordMan::SetBattleRecoveryInfo
		p.EncodeInt32(0) // nHPRecovery
		p.EncodeInt32(0) // nMPRecovery
	}
	return p.buf
}

// CWvsContext::OnChangeSkillRecordResult
func NewChangeSkillRecordResult(isSetExclRequestSent, isShowResult, isRemoveLinkSkill bool, char *character.Character, sn byte) []byte {
	p := newOutPacket(opcode.CWvsContext_OnChangeSkillRecordResult)
	p.EncodeBool(isSetExclRequestSent) // bSetExclRequestSent
	p.EncodeBool(isShowResult)         // bShowResult
	p.EncodeBool(isRemoveLinkSkill)    // bRemoveLinkSkill
	p.EncodeUint16(uint16(len(char.Skill.Lists)))
	for _, skill := range char.Skill.Lists {
		p.EncodeUint32(skill.ID)           // nSkillID
		p.EncodeUint32(skill.CurrentLevel) // nInfo
		p.EncodeUint32(skill.MasterLevel)  // nInfo
		p.EncodeFT(util.ZeroTime)          // dateExpire
	}
	p.EncodeByte(0) // bSN?
	return p.buf
}

// CWvsContext::OnTownPortal
func NewTownPortal(townID, fieldID int32) []byte {
	p := newOutPacket(opcode.CWvsContext_OnTownPortal)
	p.EncodeInt32(townID)  // m_dwTownID
	p.EncodeInt32(fieldID) // m_dwFieldID
	if townID != maple.EmptyPortalID && fieldID != maple.EmptyPortalID {
		p.EncodeInt32(0) // m_nSKillID
		p.EncodeInt16(0) // m_ptFieldPortal.x
		p.EncodeInt16(0) // m_ptFieldPortal.y
	}
	return p.buf
}

// CWvsContext::OnSetPotionDiscountRate
func NewSetPotionDiscountRate() []byte {
	p := newOutPacket(opcode.CWvsContext_OnSetPotionDiscountRate)
	p.EncodeByte(0) // m_nPotionDiscountRate
	return p.buf
}

// CWvsContext::OnSetBuyEquipExt
func NewSetBuyEquipExt() []byte {
	p := newOutPacket(opcode.CWvsContext_OnSetBuyEquipExt)
	p.EncodeByte(1) // m_bBuyEquipExt
	return p.buf
}

// CWvsContext::UNK411
func NewUNK411() []byte {
	p := newOutPacket(opcode.CWvsContext_UNK411)
	p.EncodeInt32(0)
	condition1 := false
	p.EncodeBool(condition1) // if true call CWvsContext::UI_Open(this, 562, -1, 0, 0);
	condition2 := false
	p.EncodeBool(condition2) // if true call CWvsContext::UI_Close(this, 562);
	condition3 := false
	p.EncodeBool(condition3)
	if !condition2 && condition3 {
		p.EncodeInt32(0)
	}
	return p.buf
}

// CWvsContext::OnHairStyleCoupon?
func NewHairStyleCoupon() []byte {
	p := newOutPacket(opcode.CWvsContext_OnHairStyleCoupon)
	v9 := 2
	for i := 0; i < v9; i++ {
		v8 := 0
		p.EncodeUint32(uint32(v8))
		for i := 0; i < v8; i++ {
			p.EncodeInt32(0) // Hair Style Coupon ID and CBD Hair Style Coupon (VIP)?
			v5 := 0
			p.EncodeUint32(uint32(v5)) // Hair count?
			for i := 0; i < v5; i++ {
				p.EncodeInt32(0) // Hair ID?
				p.EncodeByte(0xFF)
			}
		}
	}
	return p.buf
}

// CWvsContext_OnExpedtionResult
func NewExpedtionResult(retCode maple.ExpedtionRetCode) []byte {
	p := newOutPacket(opcode.CWvsContext_OnExpedtionResult)
	p.EncodeByte(byte(retCode)) // nRetCode
	return p.buf
}

// CWvsContext::OnMessage case 1
// CWvsContext::QuestRecordMessage
func NewQuestRecordMessage(quest *nx.QuestNX) []byte {
	p := newOutPacket(opcode.CWvsContext_OnMessage)
	p.EncodeByte(byte(maple.QuestRecordMessage))
	CWvsContextQuestRecordMessage(&p, quest)
	return p.buf
}

// CWvsContext::OnMessage case 13
// CWvsContext::QuestRecordMessage
func NewQuestRecordExMessage(quest *nx.QuestNX) []byte {
	p := newOutPacket(opcode.CWvsContext_OnMessage)
	p.EncodeByte(byte(maple.QuestRecordExMessage))
	CWvsContextQuestRecordExMessage(&p, quest)
	return p.buf
}
