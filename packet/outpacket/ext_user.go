package outpacket

import (
	"goms/util"
	"time"
)

// Call by CStage::OnSetField
// Call by CUserRemote::Init
// CUser::StarPlanetRank::Decode
func CUserStarPlanetRankEncode(p *outPacket) {
	result := false
	p.EncodeBool(result)
	if result {
		p.EncodeInt32(0) // nRoundID
		index := 0
		p.EncodeByte(byte(index))
		if index >= 10 {
			for i := 0; i < 10; i++ {
				p.EncodeInt32(0) // *(anRanking - 10)
				p.EncodeInt32(0) // *anRanking
				p.EncodeInt32(0) //  anRanking[10] time?
			}
		} else {
			p.EncodeInt32(0) // anPoint
			p.EncodeInt32(0) // anRanking
			p.EncodeInt32(0) // atLastCheckRank time?
		}
		p.EncodeFT(time.Now()) // ftShiningStarExpiredTime
		p.EncodeUint32(0)      // nShiningStarPickedCount
		p.EncodeUint32(0)      // nRoundStarPoint
	}
}

// Call by CUserRemote::Init
// CUser::DecodeStarPlanetTrendShopLook
func CUserEncodeStarPlanetTrendShopLook(p *outPacket) {
	count := 0
	p.EncodeUint32(uint32(count)) // will break REMOTE_AVATAR_MODIFIED if count!=0?
	for i := 0; i < count; i++ {
		p.EncodeUint32(0)
	}
}

// Call by CStage::OnSetField
// Call by CUserRemote::Init
// CUser::DecodeTextEquipInfo
func CUserEncodeTextEquipInfo(p *outPacket) {
	result := 0
	p.EncodeInt32(int32(result))
	if result > 0 {
		for i := 0; i < result; i++ {
			p.EncodeInt32(0) // first
			p.EncodeStr("")  // sText
		}
	}
}

// Call by CStage::OnSetField
// Call by CUserRemote::Init
// CUser::DecodeFreezeHotEventInfo
func CUserEncodeFreezeHotEventInfo(p *outPacket, accountType uint8, accountID uint32) {
	// FreezeAndHotEventInfo::Decode
	p.EncodeByte(accountType)
	p.EncodeUint32(accountID)
}

// Call by CStage::OnSetField
// Call by CUserRemote::Init
// CUser::DecodeEventBestFriendInfo
func CUserEncodeEventBestFriendInfo(p *outPacket) {
	p.EncodeInt32(0) // m_dwEventBestFriendAID
}

// Call by CStage::OnSetField
// Call by CUserRemote::Init
// CUser::DecodeSundayMaple
func CUserEncodeSundayMaple(p *outPacket) {
	count := 0
	p.EncodeUint32(uint32(count))
	for i := 0; i < count; i++ {
		p.EncodeUint32(0)
	}
}

// Call by CField::OnSetQuickMoveInfo
// QuickMoveInfo::Decode
func QuickMoveInfoEncode(p *outPacket) {
	p.EncodeInt32(0)          // nQmiID
	p.EncodeInt32(0)          // dwTemplateID
	p.EncodeInt32(0)          // nCode
	p.EncodeInt32(0)          // nLevelMin
	p.EncodeStr("")           // sMessage
	p.EncodeFT(util.ZeroTime) // ftStart
	p.EncodeFT(util.ZeroTime) // ftEnd
}
