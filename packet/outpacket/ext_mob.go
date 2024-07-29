package outpacket

import "goms/util"

// Call by CharacterData::Decode
// MonsterBattle_MobInfo::Decode
func MonsterBattleMobInfoEncode(p *outPacket) {
	p.EncodeInt32(0) // dwMobID
	p.EncodeInt32(0) // dwLevel
	p.EncodeInt32(0) // dwExp
	p.EncodeInt32(0) // nMonsterAbility1
	p.EncodeInt32(0) // nMonsterAbility2
	p.EncodeInt32(0) // nMonsterAbility3
	p.EncodeByte(0)  // nMobType
	p.EncodeInt32(0) // dwMainCharacterID
	p.EncodeInt32(0) // nSlotIndex
}

// Call by CharacterData::Decode
// TODO GW_MonsterBattleLadder_UserInfo::Decode
func GWMonsterBattleLadderUserInfoEncode(p *outPacket) {
	p.EncodeUint32(0)         // nWorldID
	p.EncodeUint32(0)         // dwAccountID
	p.EncodeUint32(0)         // dwCharacterID
	p.EncodeUint32(0)         // nPoint
	p.EncodeUint32(0)         // nWin
	p.EncodeUint32(0)         // nDefeat
	p.EncodeUint32(0)         // nDraw
	p.EncodeUint32(0)         // nCountBattle
	p.EncodeUint32(0)         // nCountPVP
	p.EncodeFT(util.ZeroTime) // ftBPUpdateTime
	p.EncodeFT(util.ZeroTime) // ftPVPUpdateTime
	p.EncodeFT(util.ZeroTime) // ftRegisterDate
}

// Call by CharacterData::Decode
// TODO GW_MonsterBattleRankInfo::Decode
func GWMonsterBattleRankInfoEncode(p *outPacket) {
	p.EncodeByte(0)   // nType
	p.EncodeInt32(0)  // nRank
	p.EncodeByte(0)   // nWorldID
	p.EncodeUint32(0) // dwAccountID
	p.EncodeUint32(0) // dwCharacterID
	p.EncodeUint32(0) // nPoint
	p.EncodeUint32(0) // dwMobID1
	p.EncodeUint32(0) // dwMobID2
	p.EncodeUint32(0) // dwMobID3
	p.EncodeStr("")   // sCharacterName
}

// Call by CharacterData::Decode
// ReadLeafInfo::Decode
func ReadLeafInfoEncode(p *outPacket, accountID, characterID uint32) {
	// sub_751EF0 RedLeafInfo::Decode
	ids := []uint32{9410165, 9410166, 9410167, 9410168, 9410198}
	p.EncodeUint32(accountID)
	p.EncodeUint32(characterID)
	p.EncodeUint32(uint32(len(ids))) // ids count 5
	// CInPacket::DecodeBuffer(0x28)
	for i := 0; i < len(ids); i++ {
		p.EncodeUint32(0) // unk
		p.EncodeUint32(ids[i])
	}
}
