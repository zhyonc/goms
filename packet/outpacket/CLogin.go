package outpacket

import (
	"fmt"
	"goms/maple"
	"goms/mongodb/model"
	"goms/mongodb/model/character"
	"goms/mongodb/model/inventory"
	"goms/opcode"
	"goms/util"
	"strconv"
	"time"
)

// Only see in CMS
func NewRSAKey() []byte {
	p := newOutPacket(opcode.CLogin_OnRSAKey)
	num := util.GetRandomNum(4, 7)
	p.EncodeStr("MapLogin" + strconv.Itoa(num))
	p.EncodeUint32(util.Time2YMDH(time.Now()))
	return p.buf
}

// CLogin::OnCheckPasswordResult
func NewCheckPasswordResultBanned(account *model.Account) []byte {
	p := newOutPacket(opcode.CLogin_OnCheckPasswordResult)
	p.EncodeByte(byte(maple.TipBanned))
	p.EncodeByte(0) // nBlockReason: CMS unused
	// Banned time
	if account.IsForeverBanned {
		p.EncodeInt64(-1) // dtUnblockDate
	} else {
		p.EncodeFT(account.TempBannedExpireDate) // dtUnblockDate
	}
	return p.buf
}

// CLogin::OnCheckPasswordResult
func NewCheckPasswordResult(result maple.Tip, account *model.Account) []byte {
	p := newOutPacket(opcode.CLogin_OnCheckPasswordResult)
	p.EncodeByte(byte(result))
	if result != maple.TipSuccess {
		return p.buf
	}
	p.EncodeStr(account.Username)   // sID
	p.EncodeUint32(account.ID)      // dwAccountId
	p.EncodeBool(account.Gender)    // nGender
	p.EncodeByte(account.GradeCode) // nGradeCode
	var specialEffect uint32 = 0
	switch account.GMLevel {
	case 1:
		specialEffect |= 1 << 4 // 0x10
	case 2:
		specialEffect |= 1 << 5 // 0x20
	case 3:
		specialEffect |= 1 << 13 // 0x2000
	case 4:
		specialEffect |= 1 << 4
		specialEffect |= 1 << 5
		specialEffect |= 1 << 13
	}
	p.EncodeUint32(specialEffect)
	p.EncodeUint32(0)                 // m_nAge
	p.EncodeByte(account.PurchaseExp) // nPurchaseExp
	if account.ChatUnblockDate.After(time.Now()) {
		p.EncodeBool(true)                  // nChatBlockReason
		p.EncodeFT(account.ChatUnblockDate) // chat block expire time
	} else {
		p.EncodeBool(false)
		p.EncodeFT(util.ZeroTime)
	}
	p.EncodeByte(0)  // UNK
	p.EncodeInt64(0) // AccountType: m_bManagerAccount/m_bTesterAccount/m_bSubTesterAccount/m_nVIPGrade
	p.EncodeStr("")  // Username mask?
	enableClass := len(maple.EnableClassCreation) > 0
	p.EncodeBool(enableClass)
	if enableClass {
		p.EncodeBool(true)
		for i := 0; i <= int(maple.ClassCount); i++ {
			_, ok := maple.EnableClassCreation[uint32(i)]
			p.EncodeBool(ok)
			p.EncodeUint16(1)
		}
	}
	p.EncodeBool(false)
	p.EncodeUint32(account.ShiningStarCount) // nShiningStarCount
	p.EncodeBool(false)                      // IsBeginningUser?
	p.EncodeStr("")
	p.EncodeStr(account.Username) // Show username on sdo login tip
	p.EncodeBool(true)            // Set false will show fill adult info tip
	p.EncodeBool(true)            // Is Adult?
	p.EncodeByte(0)
	return p.buf
}

// CLogin::OnChooseGender
func NewChooseGender(username string) []byte {
	p := newOutPacket(opcode.CLogin_OnChooseGender)
	p.EncodeStr(username)
	return p.buf
}

// CLogin::OnGenderSetResult
func NewGenderSetResult(ok bool) []byte {
	p := newOutPacket(opcode.CLogin_OnGenderSetResult)
	if !ok {
		p.EncodeBool(true) // Show failed tip
		return p.buf
	}
	p.EncodeBool(false)
	p.EncodeStr("")
	p.EncodeStr("")
	return p.buf
}

// CLogin::OnWorldInformation
func NewWorldInformation(worldID maple.WorldID, tag maple.WorldTag, channelPorts []int, OnlineLimitPerChannel uint32, gaugePx []uint32, ballons []maple.Ballon) []byte {
	p := newOutPacket(opcode.CLogin_OnWorldInformation)
	worldName := maple.WorldNameMap[worldID]
	p.EncodeUint16(uint16(worldID)) // nWorldID
	p.EncodeLocalStr(worldName)
	p.EncodeByte(byte(tag))
	p.EncodeStr("")     // World desc? must "" otherwise crash
	p.EncodeUint16(100) // UNK must 100 otherwise crash
	p.EncodeUint16(100) // UNK must 100 otherwise crash
	p.EncodeByte(byte(len(channelPorts)))
	p.EncodeUint32(OnlineLimitPerChannel)
	for i := 0; i < len(channelPorts); i++ {
		p.EncodeLocalStr(fmt.Sprintf("%s-%d", worldName, i+1))
		p.EncodeUint32(gaugePx[i])  // nGaugePx current online player num every channel
		p.EncodeByte(byte(worldID)) // nWorldID
		p.EncodeByte(byte(i))       // nChannelID channel index
		p.EncodeBool(false)         // bAdultChannel is Adult channel
	}
	p.EncodeUint16(uint16(len(ballons))) // ballons length
	for _, ballon := range ballons {
		p.EncodeUint16(ballon.NX)
		p.EncodeUint16(ballon.NY)
		p.EncodeLocalStr(ballon.Message)
	}
	p.EncodeInt32(0)    // m_uOffset
	p.EncodeBool(false) // m_nState connect with star planet stuff
	return p.buf
}

// CLogin::OnWorldInformation
func NewWorldInformationEnd() []byte {
	p := newOutPacket(opcode.CLogin_OnWorldInformation)
	p.EncodeInt16(-1)
	p.EncodeFT(time.Now())
	p.EncodeBool(false) // m_bNotActiveAccountDlgFocus
	p.EncodeBool(false) // true will call CTerminateException make client exit
	return p.buf
}

// CLogin::OnRecommendWorldMessage
func NewRecommendWorldMessage(worldCount int, worldID maple.WorldID, desc string) []byte {
	p := newOutPacket(opcode.CLogin_OnRecommendWorldMessage)
	p.EncodeByte(byte(worldCount))
	for i := 0; i < worldCount; i++ {
		p.EncodeInt32(int32(worldID))
		p.EncodeStr(desc)
	}
	return p.buf
}

// CLogin::OnWorldStatus
func NewWorldStatus(status maple.WorldStatus) []byte {
	p := newOutPacket(opcode.CLogin_OnWorldStatus)
	p.EncodeByte(byte(status))
	return p.buf
}

// CLogin::OnSetClientKey
func NewSetClientKey() []byte {
	// Not sure how to use, maybe 4->5->6?
	p := newOutPacket(opcode.CLogin_OnSetClientKey)
	p.EncodeUint64(0)
	return p.buf
}

// CLogin::OnSetPhysicalWorldID
func NewSetPhysicalWorldID(worldID maple.WorldID) []byte {
	//  Not sure how to use, maybe 4->5->6?
	p := newOutPacket(opcode.CLogin_OnSetPhysicalWorldID)
	p.EncodeUint32(uint32(worldID))
	return p.buf
}

// CLogin::OnSelectWorldResult
func NewSelectWorldResultFailed(result maple.Tip) []byte {
	p := newOutPacket(opcode.CLogin_OnSelectWorldResult)
	p.EncodeByte(byte(result))
	p.EncodeStr("")
	return p.buf
}

// CLogin::OnSelectWorldResult
func NewSelectWorldResultSuccess(worldID maple.WorldID,
	charOrderIDs []uint32, reservedChars []*character.Character, chars []*character.Character, invMap map[uint32]*inventory.Inventory,
	renameCharIDs []uint32, renameCharEventStartDate time.Time, renameCharEventEndDate time.Time, burningCharLength int) []byte {
	p := newOutPacket(opcode.CLogin_OnSelectWorldResult)
	p.EncodeByte(byte(maple.TipSuccess))
	p.EncodeStr("normal") // sWorldType reboot?
	p.EncodeInt32(4)      // m_nTrunkSlotCount
	p.EncodeBool(false)   // m_bBurningEventBlock
	// Deleted Characters
	p.EncodeInt32(int32(len(reservedChars))) // reserved char length
	p.EncodeFT(time.Now())                   // ftReservedDate
	for _, char := range reservedChars {
		p.EncodeUint32(char.ID)
		p.EncodeFT(char.ReservedDate) // ftReservedDate: client will add one day base on ftReservedDate
	}
	// Characters position order
	p.EncodeBool(true) // m_bIsEditedList
	p.EncodeInt32(int32(len(charOrderIDs)))
	for _, charID := range charOrderIDs {
		p.EncodeUint32(charID) // p_m_aCharacterSelectList
	}
	// Characters entity
	p.EncodeByte(byte(len(chars)))
	for i := 0; i < len(chars); i++ {
		char := chars[i]
		AvatarDataEncode(&p, char, invMap[char.ID])
		p.EncodeByte(byte(i % 4)) // every 4 char as 1 row
	}
	p.EncodeByte(3)                                // bLoginOpt SPWMode=3
	p.EncodeBool(false)                            // bQuerySSNOnCreateNewCharacter
	p.EncodeUint32(uint32(maple.MaxCharacterSlot)) // m_nSlotCount: remaining empty character slots that can be created
	p.EncodeInt32(0)                               // m_nBuyCharCount: buy character slot card
	p.EncodeInt32(0xFF)                            // m_nEventNewCharJob
	p.EncodeInt32(0)
	p.EncodeInt32(0)
	p.EncodeInt32(0)
	p.EncodeByte(0)
	p.EncodeByte(0)
	p.EncodeInt32(0)
	enableClass := len(maple.EnableClassCreation) > 0
	p.EncodeBool(enableClass)
	if enableClass {
		p.EncodeBool(true)
		for i := 0; i <= int(maple.ClassCount); i++ {
			_, ok := maple.EnableClassCreation[uint32(i)]
			p.EncodeBool(ok)
			p.EncodeUint16(1)
		}
	}
	return p.buf
}

// CLogin::OnCreateCharStep
func NewCreateCharStep() []byte {
	p := newOutPacket(opcode.CLogin_OnCreateCharStep)
	p.EncodeByte(2) // 4-2-2 call
	p.EncodeBool(false)
	return p.buf
}

// CLogin::OnCheckDuplicatedIDResult
func NewCheckDuplicatedIDResult(nickname string, isDuplicateNickname bool) []byte {
	p := newOutPacket(opcode.CLogin_OnCheckDuplicatedIDResult)
	p.EncodeStr(nickname)
	p.EncodeBool(isDuplicateNickname)
	return p.buf
}

// CLogin::OnCreateNewCharacterResult
func NewCreateNewCharacterResult(result maple.Tip, char *character.Character, inv *inventory.Inventory) []byte {
	p := newOutPacket(opcode.CLogin_OnCheckSPWOnCreateNewCharacterResult)
	p.EncodeByte(byte(result))
	if result != maple.TipSuccess {
		return p.buf
	}
	AvatarDataEncode(&p, char, inv)
	return p.buf
}

// CLogin::OnReservedDeleteCharacterResult
func NewReservedDeleteCharacterResult(characterID uint32, result maple.Tip, reservedDate time.Time) []byte {
	p := newOutPacket(opcode.CLogin_OnReservedDeleteCharacterResult)
	p.EncodeUint32(characterID)
	p.EncodeByte(byte(result))
	if result != maple.TipSuccess {
		// ignore other result CLoginUtilDlg::Error
		return p.buf
	}
	p.EncodeFT(time.Now())   // ftServerTime
	p.EncodeFT(reservedDate) // ftReservedDate: client will add one day base on ftReservedDate
	return p.buf
}

// CLogin::OnReservedDeleteCharacterCancelResult
func NewReservedDeleteCharacterCancelResult(characterID uint32, result maple.Tip) []byte {
	p := newOutPacket(opcode.CLogin_OnReservedDeleteCharacterCancelResult)
	p.EncodeUint32(characterID)
	p.EncodeByte(byte(result))
	return p.buf
}

// CLogin::OnDeleteCharacterResult
func NewDeleteCharacterResult(characterID uint32, result maple.Tip) []byte {
	p := newOutPacket(opcode.CLogin_OnDeleteCharacterResult)
	p.EncodeUint32(characterID)
	p.EncodeByte(byte(result))
	if result != maple.TipSuccess {
		return p.buf
	}
	p.EncodeBool(false)
	p.EncodeByte(0)
	return p.buf
}

// CLogin::OnSelectCharacterResult
func NewSelectCharacterResult(result maple.Tip, ip4 []byte, port int, characterID uint32) []byte {
	p := newOutPacket(opcode.CLogin_OnSelectCharacterResult)
	p.EncodeByte(byte(result)) // error code
	p.EncodeByte(0)            // dwCharacterID?
	if result != maple.TipSuccess {
		return p.buf
	}
	p.EncodeBuffer(ip4)               // NexonIP
	p.EncodeUint16(uint16(port))      // ChannelPort 8585
	p.EncodeUint32(0)                 // uChatIp 0.0.0.0
	p.EncodeUint16(0)                 // uChatPort 0
	p.EncodeInt32(int32(characterID)) // dwCharacterID?
	p.EncodeByte(0)                   // bAuthenCode
	p.EncodeInt32(0)                  // ulArgument
	return p.buf
}
