package outpacket

import (
	"goms/mongodb/model/character"
	"goms/opcode"
	"time"
)

func NewSetField(channelIndex uint8, char *character.Character, usingBuffProtector bool) []byte {
	p := newOutPacket(uint16(opcode.OnSetField))
	p.EncodeUint32(uint32(channelIndex))
	p.EncodeByte(0)   // dev
	p.EncodeInt32(0)  // old driver ID
	p.EncodeByte(1)   // characterData ? 1 : 2
	p.EncodeUint32(0) // FieldType unused
	p.EncodeByte(0)   // V196 new
	p.EncodeInt32(0)  // field.width
	p.EncodeInt32(0)  // field.height
	load := true
	p.EncodeBool(load)
	nNotifierCheck := 0
	p.EncodeUint16(0) // nNotifierCheck
	if nNotifierCheck > 0 {
		// pBlockReasonIter
		p.EncodeString("")
		// sNotifierMessage
		for i := 0; i < nNotifierCheck; i++ {
			p.EncodeString("")
		}
	}
	if load {
		p.EncodeUint32(0) // seed1
		p.EncodeUint32(0) // seed2
		p.EncodeUint32(0) // seed3
		// There need CharacterData::Decode
		bUNK1 := false
		p.EncodeBool(bUNK1)
		if bUNK1 {
			bUNK2 := false
			p.EncodeBool(bUNK2)
			if bUNK2 {
				p.EncodeInt32(0)
			}
			p.EncodeUint64(0)
			p.EncodeByte(0)
			p.EncodeUint64(0)
		}
		// goto LABEL_65
	} else {
		// LABEL_45
		usingBuffProtector := false
		p.EncodeBool(usingBuffProtector)
		// LABEL_46
		p.EncodeInt32(0)             // field.ID
		p.EncodeByte(0)              // portal
		p.EncodeUint32(char.Stat.HP) //HP
		bUNK := false
		p.EncodeBool(bUNK)
		if bUNK {
			p.EncodeInt32(0)
			p.EncodeInt32(0)
		}
		p.EncodeInt32(0)
	}
	// LABEL_65
	p.EncodeBool(false)    // CWvsContext::SetWhiteFadeInOut
	p.EncodeBool(false)    // bChatBlockReason
	p.EncodeFT(time.Now()) // ChatBlockTime
	p.EncodeInt32(0)
	// CustomField
	isCustomField := false
	p.EncodeBool(isCustomField)
	if isCustomField {
		p.EncodeInt32(0)   // partyBonusExpRate
		p.EncodeString("") // BGM
		p.EncodeInt32(0)   // bgFieldID
	}
	p.EncodeBool(false) // CWvsContext::OnInitPvPStat true-> Enable TMS PVP Map toggle
	p.EncodeBool(false) // bCanNotifyAnnouncedQuest
	p.EncodeBool(false) // isSeparatedSpJob?
	p.EncodeInt32(0)
	p.EncodeBool(false)
	p.EncodeFT(time.Now())
	// MapID
	count := 0
	p.EncodeByte(byte(count))
	if count > 0 {
		for i := 0; i < count; i++ {
			p.EncodeString("")
		}
	}
	//CUser::EncodeTextEquipInfo
	count1 := 0
	p.EncodeInt32(int32(count1))
	if count1 > 0 {
		p.EncodeInt32(0)
		p.EncodeString("")
	}
	// CUser::EncodeFreezeHotEventInfo
	p.EncodeByte(0)
	p.EncodeInt32(0)
	// CUser::EventBestFriendInfo
	p.EncodeInt32(0)
	// MenuItem
	setMenuItem := true
	p.EncodeBool(setMenuItem)
	if setMenuItem {
		for i := 0; i < 5; i++ {
			p.EncodeInt32(0)
		}
		p.EncodeString("") // say hello
	}
	// Sunday Event
	bSundayMaple := false
	p.EncodeBool(bSundayMaple)
	if bSundayMaple {
		p.EncodeString("UI/UIWindowEvent.img/sundayMaple")
		p.EncodeString("") // EventDesc
		p.EncodeString("") // Date
		p.EncodeInt32(0)
		p.EncodeInt32(0)
	}
	p.EncodeInt32(0)
	p.EncodeByte(0)
	p.EncodeInt32(0)
	p.EncodeByte(0)
	return p.buf
}
