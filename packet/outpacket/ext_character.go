package outpacket

import (
	"goms/maple"
	"goms/mongodb/model/character"
	"goms/mongodb/model/inventory"
	"goms/util"
	"time"
)

// Call by CLogin::OnCreateNewCharacterResult
// Call by CLogin::OnSelectWorldResult
// AvatarData::Decode
func AvatarDataEncode(p *outPacket, char *character.Character, inv *inventory.Inventory) {
	GWCharacterStatEncode(p, char)
	AvatarLookEncode(p, char, inv) // sub_7870F0
	if maple.IsZeroJob(char.Job) {
		AvatarLookEncode(p, nil, nil) // Additional appearance
	}
}

// Call by AvatarData::Decode
// GW_CharacterStat::Decode
func GWCharacterStatEncode(p *outPacket, char *character.Character) {
	p.EncodeUint32(char.ID)                                 // dwCharacterID
	p.EncodeUint32(char.ID)                                 // dwCharacterIDForLog
	p.EncodeUint32(uint32(char.WorldID))                    // dwWorldIDForLog
	p.EncodeLocalName(char.Name, maple.CharacterNameLength) // sCharacterName
	p.EncodeBool(char.Gender)                               // nGender
	p.EncodeByte(char.Look.SkinColor)                       // nSkin
	p.EncodeUint32(char.Look.Face)                          // nFace
	p.EncodeUint32(char.Look.Hair)                          // nHair
	p.EncodeByte(char.Look.MixBaseHairColor)                // nMixBaseHairColor
	p.EncodeByte(char.Look.MixAddHairColor)                 // nMixAddHairColor
	p.EncodeByte(char.Look.MixHairBaseProb)                 // nMixHairBaseProb
	p.EncodeByte(char.Stat.Level)                           // Limit at 250
	p.EncodeUint16(char.Job)
	p.EncodeUint16(char.Stat.Str)
	p.EncodeUint16(char.Stat.Dex)
	p.EncodeUint16(char.Stat.Int)
	p.EncodeUint16(char.Stat.Luk)
	p.EncodeUint32(char.Stat.HP)
	p.EncodeUint32(char.Stat.MaxHP)
	p.EncodeUint32(char.Stat.MP)
	p.EncodeUint32(char.Stat.MaxMP)
	p.EncodeUint16(char.Stat.AP)
	if maple.IsNotExtendSPJob(char.Job) {
		p.EncodeUint16(char.Skill.SP)
	} else {
		ExtendSPEncode(p, char)
	}
	p.EncodeUint64(char.Stat.Exp)       // nEXP64
	p.EncodeUint32(char.Stat.Pop)       // nPOP
	p.EncodeUint32(char.WeaponPoint)    // nWP Zero job use WeaponPoint
	p.EncodeUint32(char.GachaponExp)    // Gachapon Exp
	p.EncodeInt32(0)                    // UNK
	p.EncodeInt64(util.ZeroTime.Unix()) // playtime in seconds?
	p.EncodeUint32(char.PosMap)         // dwPosMap
	p.EncodeByte(char.Portal)           // nPortal
	p.EncodeUint16(char.SubJob)
	if maple.IsDamonJob(char.Job) || maple.IsXenonJob(char.Job) || maple.IsZeroJob(char.Job) {
		p.EncodeUint32(char.Look.DefFaceAcc) // nDefFaceAcc
	}
	// CharacterTrait
	p.EncodeByte(char.Trait.Fatigue)
	p.EncodeUint32(util.Time2YMDH(char.Trait.LastFatigueUpdateTime)) // YYYYMMDDHH
	p.EncodeUint32(maple.TraitTotalExp[char.Trait.CharismaLevel])    // nCharismaEXP[0]
	p.EncodeUint32(maple.TraitTotalExp[char.Trait.InsightLevel])     // nInsightEXP[0]
	p.EncodeUint32(maple.TraitTotalExp[char.Trait.WillLevel])        // nWillEXP[0]
	p.EncodeUint32(maple.TraitTotalExp[char.Trait.CraftLevel])       // nCraftEXP[0]
	p.EncodeUint32(maple.TraitTotalExp[char.Trait.SenseLevel])       // nSenseEXP[0]
	p.EncodeUint32(maple.TraitTotalExp[char.Trait.CharmLevel])       // nCharmEXP[0]
	// NonCombatStatDayLimit CInPacket::DecodeBuffer(21) Start
	p.EncodeUint16(char.Trait.CharismaExp)
	p.EncodeUint16(char.Trait.InsightExp)
	p.EncodeUint16(char.Trait.WillExp)
	p.EncodeUint16(char.Trait.CraftExp)
	p.EncodeUint16(char.Trait.SenseExp)
	p.EncodeUint16(char.Trait.CharmExp)
	p.EncodeByte(char.Trait.CharmByCashPR)
	p.EncodeFT(char.Trait.LastUpdateCharmByCashPR)
	// NonCombatStatDayLimit CInPacket::DecodeBuffer(21) End
	// CharacterPVP
	p.EncodeUint32(char.PVP.Exp)
	p.EncodeByte(char.PVP.Grade)
	p.EncodeUint32(char.PVP.Point)
	p.EncodeByte(char.PVP.ModeLevel)
	p.EncodeByte(char.PVP.ModeType)
	p.EncodeUint32(char.PVP.EventPoint)
	// addPartTimeJob start
	p.EncodeByte(0)           // 0<nAlbaActivityID<6
	p.EncodeFT(util.ZeroTime) // AlbaStartTime
	p.EncodeInt32(0)          // nAlbaDuratio
	p.EncodeBool(false)       // bAlbaSpecialReward
	// addPartTimeJob end
	CharacterCardEncode(p, char)
	p.EncodeFT(util.ZeroTime)       // stAccount_LastLogout
	CharacterBurningEncode(p, char) // bBurning->CharacterBurning
	p.EncodeInt32(0)                // UNK
}

// Call by GW_CharacterStat::Decode
// ExtendSP::Decode
func ExtendSPEncode(p *outPacket, char *character.Character) {
	p.EncodeByte(byte(len(char.Skill.ExtendSPs)))
	for _, extendSP := range char.Skill.ExtendSPs {
		p.EncodeByte(extendSP.JobLevel)
		p.EncodeUint32(extendSP.JobSP)
	}
}

// Call by GW_CharacterStat::Decode
// CHARACTERCARD::Decode
func CharacterCardEncode(p *outPacket, char *character.Character) {
	//  p_nJob = &this->card[0].nJob
	for i := 0; i < 9; i++ {
		card := char.Cards[i]
		p.EncodeUint32(card.CharacterID)
		p.EncodeByte(card.CharacterLevel)
		p.EncodeUint32(card.CharacterClass)
	}
}

// Call by GW_CharacterStat::Decode
// CharaterBurning::Decode
func CharacterBurningEncode(p *outPacket, char *character.Character) {
	p.EncodeFT(char.BurningStartDate)
	p.EncodeFT(char.BurningEndDate)
	p.EncodeUint32(maple.BurningMinLevel)
	p.EncodeUint32(maple.BurningMaxLevel[char.BurningType])
	p.EncodeInt32(0)
	p.EncodeByte(char.BurningType)
}

// Call by AvatarData::Decode
// Call by CUserRemote::Init
// AvatarLook::Decode
func AvatarLookEncode(p *outPacket, char *character.Character, inv *inventory.Inventory) {
	p.EncodeBool(char.Gender)         // nGender
	p.EncodeByte(char.Look.SkinColor) // nSkin
	p.EncodeUint32(char.Look.Face)    // nFace
	p.EncodeUint32(char.JobClass)     // nJob
	p.EncodeByte(0)                   // mega? 0:1
	p.EncodeUint32(char.Look.Hair)
	// 1.anHairEquip
	equipLooks := inv.EquipInv.GetEquipLook()
	for _, equipLook := range equipLooks {
		p.EncodeByte(byte(equipLook.BagIndex))
		p.EncodeUint32(equipLook.ItemID)
	}
	p.EncodeByte(0xFF)
	// 2.anUnseenEquip
	p.EncodeByte(0xFF)
	// 3.totem
	totems := inv.EquipInv.Totems
	for i := 0; i < 3; i++ {
		if totems[i].ItemID == 0 {
			continue
		}
		p.EncodeByte(byte(i))
		p.EncodeUint32(totems[i].ItemID)
	}
	p.EncodeByte(0xFF)
	p.EncodeUint32(inv.EquipInv.CashEquip.Weapon.ItemID) // nWeaponStickerID
	p.EncodeUint32(inv.EquipInv.Equip.Weapon.ItemID)     // nWeaponID
	p.EncodeUint32(inv.EquipInv.Equip.SubWeapon.ItemID)  // nSubWeaponID
	p.EncodeBool(false)                                  // bDrawElfEar
	for i := 0; i < 3; i++ {
		p.EncodeUint32(inv.EquipInv.Pets[i].ItemID) // pet ID
	}
	if maple.IsDamonJob(char.Job) || maple.IsXenonJob(char.Job) {
		//nDemonSlayerDefFaceAcc/nXenonDefFaceAcc
		p.EncodeUint32(char.Look.DefFaceAcc)
	} else if maple.IsZeroJob(char.Job) {
		// bIsZeroBetaLook
		p.EncodeBool(false)
	} else if maple.IsBeastTamer(char.Job) {
		p.EncodeUint32(char.Look.DefFaceAcc)
		p.EncodeBool(char.Look.Ear > 0)
		p.EncodeUint32(char.Look.Ear)
		p.EncodeBool(char.Look.Tail > 0)
		p.EncodeUint32(char.Look.Tail)
	}
	p.EncodeByte(char.Look.MixBaseHairColor) // nMixedHairColor
	p.EncodeByte(char.Look.MixHairPercent)   // nMixHairPercent
	p.Fill(5)                                // UNK
}

// Call by CStage::OnSetField
// CharacterData::Decode
func CharacterDataEncode(p *outPacket, dbcharFlag maple.CharFlag, char *character.Character, inv *inventory.Inventory, maplePoint, cashPoint uint32) {
	p.EncodeUint64(uint64(dbcharFlag)) // dbcharFlag
	p.EncodeByte(0)                    // nCombatOrders guessing combat orders level
	for i := 0; i < 3; i++ {
		p.EncodeInt32(0) // aPetActiveSkillCoolTime
	}
	p.EncodeByte(0)     // nPvPExp size
	p.EncodeUint32(0)   // nWillEXP size
	p.EncodeBool(false) // UNK
	// 0x1
	if dbcharFlag&maple.FlagCharacter != 0 {
		GWCharacterStatEncode(p, char)
		p.EncodeByte(50) // Buddylist.Capacity?
		hasFairyBlessing := false
		p.EncodeBool(hasFairyBlessing)
		if hasFairyBlessing {
			p.EncodeStr("")
		}
		hasEmpressBlessing := false
		p.EncodeBool(hasFairyBlessing)
		if hasEmpressBlessing {
			p.EncodeStr("")
		}
		hasUltimateExplorer := false // ultimate explorer, deprecated
		p.EncodeBool(hasUltimateExplorer)
		if hasUltimateExplorer {
			p.EncodeStr("")
		}
		p.EncodeInt32(0)
		p.EncodeByte(255)
		p.EncodeInt32(0)
		p.EncodeByte(255)
	}
	// 0x2
	if dbcharFlag&maple.FlagMoney != 0 {
		p.EncodeUint64(inv.ItemInv.Money) // GW_CharacterStat::DecodeMoney
		p.EncodeUint32(char.ID)
		p.EncodeUint32(maplePoint)
		p.EncodeUint32(cashPoint)
	}
	// 0x8 || 0x2000000 GW_ExpConsumeItem
	if dbcharFlag&maple.FlagItemSlotConsume != 0 || dbcharFlag&maple.FlagExpConsumeItem != 0 {
		expConsumeItems := inv.ItemInv.Consume.ExpConsumeItems
		p.EncodeUint32(uint32(len(expConsumeItems)))
		for _, item := range expConsumeItems {
			p.EncodeUint32(item.ItemID)
			p.EncodeUint32(item.MinLev)
			p.EncodeUint32(item.MaxLev)
			p.EncodeUint64(item.RemainingExp)
		}
	}
	// 0x8000
	if dbcharFlag&maple.FlagMonsterBattleInfo != 0 {
		count := 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			MonsterBattleMobInfoEncode(p)
		}
		p.EncodeUint32(char.ID)
		for i := 0; i < 3; i++ {
			p.EncodeInt32(0)
			p.EncodeInt32(0)
		}
		count = 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			p.EncodeUint32(0)
		}
		condition := false
		p.EncodeBool(condition)
		if condition {
			GWMonsterBattleLadderUserInfoEncode(p)
		}
		count = 0
		p.EncodeByte(byte(count))
		for i := 0; i < count; i++ {
			GWMonsterBattleRankInfoEncode(p)
		}
		count = 0
		p.EncodeByte(byte(count))
		for i := 0; i < count; i++ {
			GWMonsterBattleRankInfoEncode(p)
		}
	}
	// 0x4000000 CMS Feature
	if dbcharFlag&maple.FlagPotionPot != 0 {
		p.EncodeUint32(uint32(len(inv.PotionPotInv.Pots)))
		for _, pot := range inv.PotionPotInv.Pots {
			p.EncodeUint32(pot.ItemID)
			p.EncodeUint32(char.ID)
			p.EncodeUint32(pot.MaxCapcity)
			p.EncodeUint32(pot.HP)
			p.EncodeUint32(pot.MP)
			p.EncodeFT(pot.StartTime)
			p.EncodeFT(pot.EndTime)
		}
	}
	// 0x80
	if dbcharFlag&maple.FlagInventorySize != 0 {
		p.EncodeByte(inv.ItemInv.Equip.SlotSize)
		p.EncodeByte(inv.ItemInv.Consume.SlotSize)
		p.EncodeByte(inv.ItemInv.SetUp.SlotSize)
		p.EncodeByte(inv.ItemInv.Etc.SlotSize)
		p.EncodeByte(inv.ItemInv.Cash.SlotSize)
	}
	// 0x100000
	if dbcharFlag&maple.FlagPendantExt != 0 {
		p.EncodeFT(inv.EquipInv.PendantExtExpireDate)
	}
	// 0x4
	if dbcharFlag&maple.FlagItemSlotEquip != 0 {
		onlyEquipped := false
		p.EncodeBool(onlyEquipped)
		// EquipInventory.Equip
		for _, equip := range inv.EquipInv.GetEquips() {
			p.EncodeUint16(equip.BagIndex) // maple.EquipPart
			GWItemSlotBaseEncode(p, maple.ItemTypeEquip, equip)
		}
		p.EncodeUint16(0)
		// EquipInventory.CashEquip
		for _, cashEquip := range inv.EquipInv.GetCashEquips() {
			p.EncodeUint16(cashEquip.BagIndex) // maple.EquipPart
			GWItemSlotBaseEncode(p, maple.ItemTypeEquip, cashEquip)
		}
		p.EncodeUint16(0)
		if !onlyEquipped {
			p.EncodeUint16(0) // ItemInventory.Equip
		}
		NonBPEquipEncode(p, inv)       // NonBPEquip::Decode
		VirtualEquipInventoryEncode(p) // VirtualEquipInventory::Decode
	}
	// 0x10
	if dbcharFlag&maple.FlagItemSlotSetUp != 0 {
		// Maybe skill skin
		p.EncodeUint16(0) // 20001 ~ 20048
		p.EncodeUint16(0) // 20049~20051
	}
	// 0x8
	if dbcharFlag&maple.FlagItemSlotConsume != 0 {
		p.EncodeByte(0)
	}
	// 0x10
	if dbcharFlag&maple.FlagItemSlotSetUp != 0 {
		p.EncodeByte(0)
	}
	// 0x20
	if dbcharFlag&maple.FlagItemSlotEtc != 0 {
		p.EncodeByte(0)
	}
	// 0x40
	if dbcharFlag&maple.FlagItemSlotCash != 0 {
		p.EncodeByte(0)
	}
	// ItemInventory end
	// BagData start
	// 0x8
	if dbcharFlag&maple.FlagItemSlotConsume != 0 {
		p.EncodeInt32(0)
	}
	// 0x10
	if dbcharFlag&maple.FlagItemSlotSetUp != 0 {
		p.EncodeInt32(0)
	}
	// 0x20
	if dbcharFlag&maple.FlagItemSlotEtc != 0 {
		p.EncodeInt32(0)
	}
	// BagData end
	// 0x1000000
	if dbcharFlag&maple.FlagCoreInfo != 0 {
		p.EncodeInt32(0) // nSenseEXP
	}
	// 0x40000000
	if dbcharFlag&maple.FlagDayLimit != 0 {
		p.EncodeInt32(0) // DayLimit.nWill
	}
	// 0x800000 Gemstone Imp
	if dbcharFlag&maple.FlagItemPot != 0 {
		for i := 0; i < 3; i++ {
			itemPot := inv.ItemPotInv.Pots[i]
			if itemPot.LifeID == 0 {
				continue
			}
			p.EncodeByte(byte(i))
			GWItemPotSlotEncode(p, &itemPot)
		}
		p.EncodeByte(0)
	}
	// 0x100
	if dbcharFlag&maple.FlagSkillRecord != 0 {
		condition := true
		p.EncodeBool(condition)
		if condition {
			p.EncodeUint16(uint16(len(char.Skill.Lists))) // skills size
			for _, s := range char.Skill.Lists {
				p.EncodeUint32(s.ID)
				p.EncodeUint32(s.CurrentLevel)
				p.EncodeFT(s.ExpirationTime)
				if maple.IsSkillNeedMasterLevel(s.ID) {
					p.EncodeUint32(s.MasterLevel)
				}
			}
			p.EncodeUint16(uint16(len(char.Skill.Links)))
			for _, s := range char.Skill.Links {
				p.EncodeUint32(s.ID)
				p.EncodeUint16(s.Level)
			}
		} else {
			count := 0
			p.EncodeUint16(uint16(count))
			for i := 0; i < count; i++ {
				p.EncodeInt32(0) // nTI
				p.EncodeInt32(0) // sValue
				// SetAt nEventPoint
			}
			count = 0
			p.EncodeUint16(0)
			for i := 0; i < count; i++ {
				p.EncodeInt32(0) // nTI
				// RemoveKey nEventPoint
			}
			count = 0
			p.EncodeUint16(0)
			for i := 0; i < count; i++ {
				p.EncodeInt32(0)          // nTI
				p.EncodeFT(util.ZeroTime) // pInfo
				// SetAt nLevel_CS
			}
			count = 0
			p.EncodeUint16(0)
			for i := 0; i < count; i++ {
				p.EncodeInt32(0) // nTI
				// RemoveKey nLevel_CS
			}
			count = 0
			p.EncodeUint16(0)
			for i := 0; i < count; i++ {
				p.EncodeInt32(0) // nTI
				p.EncodeInt32(0) // sValue
				// SetAt p_nJob
			}
			count = 0
			p.EncodeUint16(0)
			for i := 0; i < count; i++ {
				p.EncodeInt32(0) // nTI
				// RemoveKey p_nJob
			}
		}
	}
	// 0x8000
	if dbcharFlag&maple.FlagSkillCooltime != 0 {
		p.EncodeUint16(uint16(len(char.Skill.CoolTimes)))
		for skillID, coolTime := range char.Skill.CoolTimes {
			p.EncodeUint32(skillID)
			remainingTime := time.Until(coolTime)
			if remainingTime > 0 {
				p.EncodeUint32(uint32(remainingTime.Seconds()))
			} else {
				p.EncodeUint32(0)
			}
		}
	}
	// 0x200
	if dbcharFlag&maple.FlagQuestRecord != 0 {
		condition := true
		p.EncodeBool(condition) // if ture RemoveAll->(mQuestRecord/lQuestRecordModified/lQuestRecordDeleted)
		count := 0
		p.EncodeUint16(0) // QuestsInProgress size
		for i := 0; i < count; i++ {
			// Use for CharacterData::SetQuest
			p.EncodeUint32(0) // nQRKey
			p.EncodeStr("")   // sQRValue
		}
		if !condition {
			count = 0
			p.EncodeUint16(uint16(count))
			for i := 0; i < count; i++ {
				// Use for for CharacterData::RemoveQuest
				p.EncodeUint32(0) // nQRKey
			}
		}
	}
	// 0x4000
	if dbcharFlag&maple.FlagQuestComplete != 0 {
		condition := true
		p.EncodeBool(condition) // if ture RemoveAll->mQuestComplete
		count := 0
		p.EncodeUint16(0) // mQuestComplete size
		for i := 0; i < count; i++ {
			p.EncodeUint32(0)                             // QRKey
			p.EncodeUint32(util.Time2YMDH(util.ZeroTime)) // Timestamp of completion
		}
		if !condition {
			count = 0
			p.EncodeUint16(uint16(count))
			for i := 0; i < count; i++ {
				p.EncodeUint32(0) // sValue
			}
		}
	}
	// 0x400
	if dbcharFlag&maple.FlagMiniGameRecord != 0 {
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			// MONSTERLIFE_INVITEINFO
			p.EncodeUint32(0) // FT.HighDateTime
			p.EncodeUint32(0) // ptr+4  = sOwnerName?
			p.EncodeUint32(0) // ptr+8  = lRewardGradeQ?
			p.EncodeUint32(0) // ptr+12 = lRewardGradeQ->_Myhead?
			p.EncodeUint32(0) // ptr+16 = lRewardGradeQ->_Mysize?
		}
	}
	// 0x800
	if dbcharFlag&maple.FlagRingRecord != 0 {
		// Couple Ring
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.Fill(33)
		}
		// Friend Ring
		count = 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.Fill(37)
		}
		// Marry Ring
		count = 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.Fill(48)
		}
	}
	// 0x1000 RocksInfo
	if dbcharFlag&maple.FlagMapTransfer != 0 {
		for i := 0; i < 5; i++ {
			p.EncodeUint32(0) // RegRocks
		}
		for i := 0; i < 10; i++ {
			p.EncodeUint32(0) // Rocks
		}
		for i := 0; i < 13; i++ {
			p.EncodeUint32(0) // HyperRocks
		}
	}
	// 0x40000
	if dbcharFlag&maple.FlagQuestRecordEx != 0 {
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.EncodeInt32(0) // key
			p.EncodeStr("")  // CSimpleStrMap::InitFromRawString
		}
	}
	// 0x2000
	if dbcharFlag&maple.FlagAvatar != 0 {
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.EncodeUint32(0) // sValue
			AvatarLookEncode(p, char, inv)
		}
	}
	// 0x1000
	if dbcharFlag&0x1000 != 0 { // unk
		count := 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			// nWillEXP_CS
			p.EncodeUint32(0) // nCount
			p.EncodeUint32(0) // sValue
		}
	}
	// 0x200000
	if dbcharFlag&maple.FlagWildHunterInfo != 0 && maple.IsWildHunter(char.Job) {
		GWWildHunterInfoEncode(p)
	}
	// 0x800
	if dbcharFlag&maple.FlagZeroInfo != 0 {
		if maple.IsZeroJob(char.Job) {
			ZeroInfoEncode(p, maple.ZeroFlagAll)
		}
	}
	// 0x4000000
	if dbcharFlag&maple.FlagShopBuyLimit != 0 {
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			count = 0
			p.EncodeUint16(uint16(count)) // nCount
			p.EncodeUint32(0)             // use npc id as shop id?
			for i := 0; i < count; i++ {
				BuyLimitDataEncode(p)
			}
		}
	}
	// 0x20000000
	if dbcharFlag&maple.FlagStolenSkills != 0 {
		for i := 0; i < 15; i++ {
			p.EncodeUint32(0)
		}
	}
	// 0x10000000
	if dbcharFlag&maple.FlagChosenSkills != 0 {
		for i := 0; i < 5; i++ {
			p.EncodeUint32(0) // EquippedStealSkill
		}
	}
	// WTF!dbcharFlag<0 in CMS138
	// Replace with 0x80000000
	if dbcharFlag&maple.FlagCharacterPotential != 0 {
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.EncodeByte(0)  // Key
			p.EncodeInt32(0) // SkillID
			p.EncodeByte(0)  // SkillLevel
			p.EncodeByte(0)  // Grade is C,B,A,S
		}
	}
	// 0x10000
	if dbcharFlag&maple.FlagSoulCollection != 0 {
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.EncodeInt32(0) // Key
			p.EncodeInt32(0) // SkillID
		}
	}
	// 0x1
	if dbcharFlag&maple.FlagCharacter != 0 {
		p.EncodeInt32(0)
		p.EncodeInt32(0)
	}
	// 0x2
	if dbcharFlag&maple.FlagItemSlot != 0 {
		condition := false
		p.EncodeBool(condition)
		if condition {
			count := 0
			p.EncodeUint16(uint16(count))
			for i := 0; i < count; i++ {
				p.EncodeUint16(0) // nCategory
				count = 0
				p.EncodeUint16(uint16(count))
				for i := 0; i < count; i++ {
					p.EncodeInt32(0) // nItemId
					p.EncodeInt32(0) // nCount
				}
			}
		} else {
			count := 0
			p.EncodeUint16(uint16(count))
			for i := 0; i < count; i++ {
				p.EncodeUint16(0) // nCategory
				p.EncodeInt32(0)  // nItemId
				p.EncodeInt32(0)  // nCount
			}
		}
	}
	// 0x4
	if dbcharFlag&maple.FlagReturnEffectInfo != 0 {
		ReturnEffectInfoEncode(p, inventory.ItemSlotBundle{})
	}
	// 0x8
	if dbcharFlag&maple.FlagDressUpInfo != 0 {
		GWDressUpInfoEncode(p, char)
	}
	// 0x200000
	if dbcharFlag&maple.FlagFarmPotential != 0 {
		FarmPotentialEncode(p) // not sure
	}
	// 0x10
	if dbcharFlag&maple.FlagEvolutionInfo != 0 {
		// GW_Core
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.EncodeUint16(0) // nPos
			GWCoreEncode(p)
		}
		count = 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.EncodeUint16(0) // nPos
			GWCoreEncode(p)
		}
	}
	// 0x80
	if dbcharFlag&maple.FloagMemorialCubeInfo != 0 {
		MemorialCubeInfoEncode(p)
	}
	// 0x400
	if dbcharFlag&maple.FlagLikePoint != 0 {
		GWLikePointEncode(p)
	}
	// 0x20000
	if dbcharFlag&maple.FlagRunnerGameRecord != 0 {
		RunnerGameRecordEncode(p, char.ID)
	}
	// 0x80000
	if dbcharFlag&maple.FlagFamiliar != 0 {
		count := 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			p.EncodeUint32(0)
			p.EncodeByte(0)
			p.EncodeByte(0)
			p.EncodeByte(0)
		}
		p.EncodeUint32(0)
		p.EncodeUint64(0)
	}
	count := 0
	p.EncodeUint16(uint16(count)) // OX_System?
	for i := 0; i < count; i++ {
		p.EncodeUint32(0)
		p.EncodeStr("")
	}
	// 0x40000
	if dbcharFlag&maple.FlagMonsterCollection != 0 {
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.EncodeUint32(0)
			p.EncodeStr("")
		}
	}
	p.EncodeBool(false) // m_bFarmOnline
	CharacterDataEncodeTextEquipInfo(p)
	// 0x100000
	if dbcharFlag&0x100000 != 0 {
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			p.EncodeUint32(0)
			p.EncodeUint32(0)
		}
	}
	// 0x8000000
	if dbcharFlag&0x8000000 != 0 {
		condition := false
		p.EncodeBool(condition)
		if condition {
			// sub_751E40 start
			p.EncodeByte(0)
			p.EncodeUint32(0)
			p.EncodeUint32(0)
			p.EncodeUint32(0)
			p.EncodeFT(time.Now())
			// sub_751E40 end
		}
		count := 0
		p.EncodeUint16(uint16(count)) // QuestInfo?
		for i := 0; i < count; i++ {
			// sub_751E10 start
			p.EncodeByte(0)
			p.EncodeUint32(0)
			p.EncodeUint32(0)
			// sub_751E10 end
		}
		count = 0
		p.EncodeUint16(uint16(count)) // Commerci trade voyage Coins?
		for i := 0; i < count; i++ {
			// sub_751E90 start
			p.EncodeUint32(0)
			p.EncodeUint32(0)
			p.EncodeUint64(0)
			// sub_751E90 end
		}
	}
	// 0x10000000
	if dbcharFlag&0x10000000 != 0 {
		p.EncodeByte(0)
	}
	// 0x20000000
	if dbcharFlag&0x20000000 != 0 {
		// sub_770F40 start
		count := 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			p.EncodeUint16(0)
			p.EncodeUint16(0)
		}
		count = 0
		p.EncodeUint32(uint32(count))
		for i := 0; i < count; i++ {
			p.EncodeUint16(0)
			p.EncodeUint32(0)
		}
		// sub_770F40 end
	}
	// 0x400000
	if dbcharFlag&maple.FlagCoreAura != 0 {
		CoreAuraEncode(p)
		p.EncodeByte(1) // UNK
	}
	// 0x80000
	if dbcharFlag&0x80000 != 0 { // unsure the flag num
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			// sub_7687D0 start
			p.EncodeUint32(0)
			p.EncodeUint32(0)
			p.EncodeStr("")
			p.EncodeByte(0)
			p.EncodeUint64(0)
			p.EncodeUint32(0)
			p.EncodeStr("")
			p.EncodeByte(0)
			p.EncodeByte(0)
			p.EncodeUint64(0)
			p.EncodeStr("")
			// sub_7687D0 end
		}
	}
	// 0x100000
	if dbcharFlag&0x100000 != 0 { // unsure the flag num
		count := 0
		p.EncodeUint16(uint16(count))
		for i := 0; i < count; i++ {
			// sub_752120 start
			// CInPacket::DecodeBuffer(0x14)
			p.Fill(20)
			// sub_752120 end
		}
	}
	// 0x200000
	if dbcharFlag&maple.FlagRedLeafInfo != 0 {
		ReadLeafInfoEncode(p, char.AccountID, char.ID)
	}
	if dbcharFlag&0x40000000 != 0 {
		p.Fill(5) // unk
	}
}

// Call by CharacterData::Decode
// NonBPEquip::Decode 10 equip type
// p.EncodeUint16(0) as the end flag
func NonBPEquipEncode(p *outPacket, inv *inventory.Inventory) {
	equipInv := inv.EquipInv
	// 1.Evan dragon equip
	p.EncodeUint16(0)
	// 2.Resistance mechanic equip
	p.EncodeUint16(0)
	// 3.Android CashEquip
	p.EncodeUint16(0)
	// 4.Angelic Buster equip
	p.EncodeUint16(0)
	// 5.Bits Inventory
	p.EncodeUint16(0)
	// 6.Zero Female CS Inventory?
	p.EncodeUint16(0)
	// 7.Monster Battle?
	p.EncodeUint16(0)
	// 8.Arcane symbols ready for 5th job?
	p.EncodeUint16(0)
	// 9.Totems equip not work
	for _, totem := range equipInv.Totems {
		if totem.BagIndex < uint16(maple.Totem1) || totem.BagIndex > uint16(maple.Totem3) {
			continue
		}
		p.EncodeUint16(totem.BagIndex) // maple.EquipPart
		GWItemSlotBaseEncode(p, maple.ItemTypeEquip, totem)
	}
	p.EncodeUint16(0)
	// 10.Haku Fan equip
	p.EncodeUint16(0)
}

// Call by CharacterData::Decode
// VirtualEquipInventory::Decode
func VirtualEquipInventoryEncode(p *outPacket) {
	// For Android
	p.EncodeUint16(0)
	p.EncodeUint16(0)
}

// Call by CharacterData::Decode
// TODO GW_WildHunterInfo::Decode
func GWWildHunterInfoEncode(p *outPacket) {
	p.EncodeByte(0) // 10 * (getIdx() + 1)?
	for i := 0; i < 5; i++ {
		p.EncodeUint32(0) // adwCapturedMob
	}
}

// Call by CharacterData::Decode
// TODO ZeroInfo::Decode
func ZeroInfoEncode(p *outPacket, zeroFlag maple.ZeroFlag) {
	p.EncodeUint16(uint16(zeroFlag))
	if zeroFlag&maple.ZeroFlagBeta != 0 {
		p.EncodeByte(0)
	}
	if zeroFlag&maple.ZeroFlagSubHP != 0 {
		p.EncodeUint32(0)
	}
	if zeroFlag&maple.ZeroFlagSubMP != 0 {
		p.EncodeUint32(0)
	}
	if zeroFlag&maple.ZeroFlagSubSkin != 0 {
		p.EncodeByte(0)
	}
	if zeroFlag&maple.ZeroFlagSubHair != 0 {
		p.EncodeUint32(0)
	}
	if zeroFlag&maple.ZeroFlagSubFace != 0 {
		p.EncodeUint32(0)
	}
	if zeroFlag&maple.ZeroFlagSubMHP != 0 {
		p.EncodeUint32(0)
	}
	if zeroFlag&maple.ZeroFlagSubMMP != 0 {
		p.EncodeUint32(0)
	}
	if zeroFlag&maple.ZeroFlagDBCharZeroLinkCashPart != 0 {
		p.EncodeUint32(0)
	}
	if zeroFlag&maple.ZeroFlagHairColor != 0 {
		p.EncodeUint32(0) // nMixBaseHairColor
		p.EncodeUint32(0) // nMixAddHairColor
		p.EncodeUint32(0) // MixHairBaseProb
	}
}

// Call by CharacterData::Decode
// TODO BuyLimitData::Decode
func BuyLimitDataEncode(p *outPacket) {
	p.EncodeUint32(0)         // dwNPCID
	p.EncodeUint16(0)         // nItemIndex
	p.EncodeUint32(0)         // nItemID
	p.EncodeUint16(0)         // nCount
	p.EncodeFT(util.ZeroTime) // expirationTime
}

// Call by CharacterData::Decode
// TODO ReturnEffectInfo::Decode
func ReturnEffectInfoEncode(p *outPacket, item inventory.ItemSlotBundle) {
	condition := false
	p.EncodeBool(condition)
	if !condition {
		return
	}
	GWItemSlotBaseEncode(p, maple.ItemTypeBundle, item)
	p.EncodeUint32(item.ItemID) // nUsedUItemID  0,1,0,0?
}

// Call by CharacterData::Decode
// GW_DressUpInfo::Decode
func GWDressUpInfoEncode(p *outPacket, char *character.Character) {
	p.EncodeUint32(char.Look.Face)                     // nFace
	p.EncodeUint32(char.Look.Hair)                     // nHair
	p.EncodeUint32(0)                                  // nClothe
	p.EncodeByte(char.Look.SkinColor)                  // nSkin
	p.EncodeUint32(uint32(char.Look.MixBaseHairColor)) // nMixBaseHairColor
	p.EncodeUint32(uint32(char.Look.MixAddHairColor))  // nMixAddHairColor
	p.EncodeUint32(uint32(char.Look.MixHairBaseProb))  // nMixHairBaseProb
}

// Call by CharacterData::Decode
// FARM_POTENTIAL::Decode
func FarmPotentialEncode(p *outPacket) {
	p.EncodeInt32(0)
	p.EncodeInt32(0)          // dwMonsterID
	p.EncodeFT(util.ZeroTime) // potentialExpire
}

// Call by CharacterData::Decode
// GW_Core::Decode
func GWCoreEncode(p *outPacket) {
	p.EncodeUint32(0) // nCoreID
	p.EncodeUint32(0) // nLeftCount
}

// Call by CharacterData::Decode
// MemorialCubeInfo::Decode
func MemorialCubeInfoEncode(p *outPacket) {
	condition := false
	p.EncodeBool(condition)
	if !condition {
		return
	}
	// GWItemSlotBaseEncode(p)
	p.EncodeInt32(0) // nCubeItemID
	p.EncodeInt32(0) // nEItemPOS
}

// Call by CharacterData::Decode
// GW_LikePoint::Decode
func GWLikePointEncode(p *outPacket) {
	p.EncodeUint32(0)         // nPoint
	p.EncodeFT(util.ZeroTime) // ftIncTime
	p.EncodeUint32(0)         // nSeason
}

// Call by CharacterData::Decode
// RunnerGameRecord::Decode
func RunnerGameRecordEncode(p *outPacket, charID uint32) {
	p.EncodeUint32(charID)
	p.EncodeUint32(0)         // nLastScore
	p.EncodeUint32(0)         // nHighscore
	p.EncodeUint32(0)         // nRunnerPoint
	p.EncodeFT(util.ZeroTime) // tLastPlayed
	p.EncodeUint32(0)         // nTotalLeft
}

// Call by GW_CharacterStat::Decode
// CharacterData::DecodeTextEquipInfo
func CharacterDataEncodeTextEquipInfo(p *outPacket) {
	count := 0
	p.EncodeUint32(uint32(count))
	for i := 0; i < count; i++ {
		p.EncodeUint32(0)
		p.EncodeStr("")
	}
}

// Call by CharacterData::Decode
// CoreAura::Decode
func CoreAuraEncode(p *outPacket) {
	p.EncodeUint32(0) // aura id
	p.EncodeUint32(0) // master id
	p.EncodeUint32(0) // master level
	p.EncodeUint32(0) // aura level
	p.EncodeUint32(0) // total point
	// core aura attr(6 int)
	p.EncodeUint32(0) // weapon attack
	p.EncodeUint32(0) // dex
	p.EncodeUint32(0) // luk
	p.EncodeUint32(0) // magic attack
	p.EncodeUint32(0) // int
	p.EncodeUint32(0) // str
	// other(4 int)
	p.EncodeUint32(0x5)    // unk
	p.EncodeUint32(0x20)   // every attr max num
	p.EncodeUint32(0x12)   // unk
	p.EncodeUint32(0x44)   // unk
	p.EncodeFT(time.Now()) // expiration time
	p.EncodeBool(false)    // set true can't use?
}

// Call by CWvsContext::OnStatChanged
// GW_CharacterStat::DecodeChangeStat
func GWCharacterStatEncodeChangeStat(p *outPacket, mask maple.CharStat, char *character.Character, money uint64) {
	p.EncodeUint64(uint64(mask)) // v138 32->64
	if mask&maple.CS_Skin != 0 {
		p.EncodeByte(char.Look.SkinColor) // nSkin
	}
	if mask&maple.CS_Face != 0 {
		p.EncodeUint32(char.Look.Face) // nFace
	}
	if mask&maple.CS_Hair != 0 {
		p.EncodeUint32(char.Look.Hair) // nHair
	}
	if mask&maple.CS_Level != 0 {
		p.EncodeByte(char.Stat.Level) // nLevel
	}
	if mask&maple.CS_Job != 0 {
		p.EncodeUint16(char.Job)    // nJob
		p.EncodeUint16(char.SubJob) // nSubJob
	}
	if mask&maple.CS_STR != 0 {
		p.EncodeUint16(char.Stat.Str) // nSTR
	}
	if mask&maple.CS_DEX != 0 {
		p.EncodeUint16(char.Stat.Dex) // nDEX
	}
	if mask&maple.CS_INT != 0 {
		p.EncodeUint16(char.Stat.Int) // nINT
	}
	if mask&maple.CS_LUK != 0 {
		p.EncodeUint16(char.Stat.Luk) // nLUK
	}
	if mask&maple.CS_HP != 0 {
		p.EncodeUint32(char.Stat.HP) // nHP
	}
	if mask&maple.CS_MHP != 0 {
		p.EncodeUint32(char.Stat.MaxHP) // nMHP
	}
	if mask&maple.CS_MP != 0 {
		p.EncodeUint32(char.Stat.MP) // nMP
	}
	if mask&maple.CS_MMP != 0 {
		p.EncodeUint32(char.Stat.MaxMP) // nMMP
	}
	if mask&maple.CS_AP != 0 {
		p.EncodeUint16(char.Stat.AP) // nAP
	}
	if mask&maple.CS_ExtendSP != 0 {
		if maple.IsNotExtendSPJob(char.Job) {
			p.EncodeUint16(char.Skill.SP)
		} else {
			ExtendSPEncode(p, char)
		}
	}
	if mask&maple.CS_EXP != 0 {
		p.EncodeUint64(char.Stat.Exp) // nEXP64
	}
	if mask&maple.CS_POP != 0 {
		p.EncodeUint32(char.Stat.Pop) // nPOP
	}
	if mask&maple.CS_Money != 0 {
		p.EncodeUint64(money) // nMoney
	}
	if mask&maple.CS_Fatigue != 0 {
		p.EncodeByte(char.Trait.Fatigue) // nFatigue
	}
	if mask&maple.CS_CharismaEXP != 0 {
		p.EncodeUint32(maple.TraitTotalExp[char.Trait.CharismaLevel]) // nCharismaEXP
	}
	if mask&maple.CS_InsightEXP != 0 {
		p.EncodeUint32(maple.TraitTotalExp[char.Trait.InsightLevel]) // nInsightEXP
	}
	if mask&maple.CS_WillEXP != 0 {
		p.EncodeUint32(maple.TraitTotalExp[char.Trait.WillLevel]) // nWillEXP
	}
	if mask&maple.CS_CraftEXP != 0 {
		p.EncodeUint32(maple.TraitTotalExp[char.Trait.CraftLevel]) // nCraftEXP
	}
	if mask&maple.CS_SenseEXP != 0 {
		p.EncodeUint32(maple.TraitTotalExp[char.Trait.SenseLevel]) // nCharmEXP
	}
	if mask&maple.CS_CharmEXP != 0 {
		p.EncodeUint32(maple.TraitTotalExp[char.Trait.CharmLevel]) // nCharmEXP
	}
	if mask&maple.CS_DayLimit != 0 {
		// NonCombatStatDayLimit CInPacket::DecodeBuffer(21) Start
		p.EncodeUint16(char.Trait.CharismaExp)
		p.EncodeUint16(char.Trait.InsightExp)
		p.EncodeUint16(char.Trait.WillExp)
		p.EncodeUint16(char.Trait.CraftExp)
		p.EncodeUint16(char.Trait.SenseExp)
		p.EncodeUint16(char.Trait.CharmExp)
		p.EncodeByte(char.Trait.CharmByCashPR)
		p.EncodeFT(char.Trait.LastUpdateCharmByCashPR)
		// NonCombatStatDayLimit CInPacket::DecodeBuffer(21) End
	}
	if mask&maple.CS_AlbaActivity != 0 {
		p.EncodeByte(0)           // nAlbaActivityID
		p.EncodeFT(util.ZeroTime) // ftAlbaStartTime
		p.EncodeUint32(0)         // nAlbaDuration
		p.EncodeBool(false)       // bAlbaSpecialReward
	}
	if mask&maple.CS_CharacterCard != 0 {
		CharacterCardEncode(p, char)
	}
	if mask&maple.CS_PVP1 != 0 {
		p.EncodeUint32(char.PVP.Exp)   // nPvPExp
		p.EncodeByte(char.PVP.Grade)   // nPvPGrade
		p.EncodeUint32(char.PVP.Point) // nPvPPoint
	}
	if mask&maple.CS_PVP2 != 0 {
		p.EncodeByte(char.PVP.ModeLevel) // nPvPModeLevel
		p.EncodeByte(char.PVP.ModeType)  // nPvPModeType
	}
	if mask&maple.CS_EventPoint != 0 { // maybe mask<0?
		p.EncodeUint32(char.PVP.EventPoint) // nEventPoint
	}
	subMask := 0 // unk v138 new
	if subMask&1 != 0 {
		p.EncodeInt64(0)
	}
	if subMask&2 != 0 {
		p.EncodeInt32(0)
	}
	if subMask&4 != 0 {
		p.EncodeInt32(0)
	}
}

// Call by CWvsContext::OnStatChanged
// GW_CharacterStat::DecodeChangeMixHairStat
func GWCharacterStatEncodeChangeMixHairStat(p *outPacket, char *character.Character) {
	p.EncodeByte(char.Look.MixBaseHairColor) // nMixBaseHairColor
	p.EncodeByte(char.Look.MixAddHairColor)  // nMixAddHairColor
	p.EncodeByte(char.Look.MixHairBaseProb)  // nMixHairBaseProb
}
