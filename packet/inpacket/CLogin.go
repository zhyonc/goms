package inpacket

import (
	"goms/maple/class"
	"goms/util"
	"net"
)

type CheckLoginAuthInfo struct {
	inPacket
	MAC      string
	Username string
	Password string
}

// MAC(6)
// UNK(10)
// NexonTab(6) "0,0,0,0,?,0" ? is 0(TMS)/1(KMS)/2(GMS)
// Username "5 0 97 100 109 105 110" admin
// Password "5 0 97 100 109 105 110" admin
func NewCheckLoginAuthInfo(data []byte) *CheckLoginAuthInfo {
	p := &CheckLoginAuthInfo{inPacket: newInPacket(data)}
	p.MAC = util.Bytes2MAC(p.DecodeBytes(6))
	p.Skip(16)
	p.Username = p.DecodeString()
	p.Password = p.DecodeString()
	return p
}

type SelectWorldButton struct {
	inPacket
	WorldID uint8
}

func NewSelectWorldButton(data []byte) *SelectWorldButton {
	p := &SelectWorldButton{inPacket: newInPacket(data)}
	_ = p.DecodeByte() // idk
	p.WorldID = p.DecodeByte()
	return p
}

type SelectWorldRequest struct {
	inPacket
	AuthType     uint8
	WorldID      uint8
	ChannelIndex uint8
	IsReLogin    bool
	AuthInfo     string
	MAC          string
	ClientIP     string
	CPU          string
	OS           string
}

func NewSelectWorldRequest(data []byte) *SelectWorldRequest {
	p := &SelectWorldRequest{inPacket: newInPacket(data)}
	p.AuthType = p.DecodeByte()
	p.WorldID = p.DecodeByte()
	p.ChannelIndex = p.DecodeByte() + 1
	p.IsReLogin = p.DecodeBool()
	if p.IsReLogin {
		p.Skip(1) // 1
		p.AuthInfo = p.DecodeString()
		p.MAC = util.Bytes2MAC(p.DecodeBytes(6))
		p.Skip(10) // UNK
		p.Skip(4)  // 0 0 0 0
		p.Skip(1)  // 0
	}
	p.ClientIP = net.IP(p.DecodeBytes(4)).String()
	p.CPU = p.DecodeString()
	p.OS = p.DecodeString()
	return p
}

type CheckSPWRequest struct {
	inPacket
	PICEncodeType int32
	PIC           string
	CharID        int32
	InVisible     bool
	MAC           string
	HWID          string
}

func NewCheckSPWRequest(data []byte) *CheckSPWRequest {
	p := &CheckSPWRequest{inPacket: newInPacket(data)}
	p.PICEncodeType = p.DecodeInt32()
	p.PIC = p.DecodeString()
	p.CharID = p.DecodeInt32()
	p.InVisible = p.DecodeBool()
	p.MAC = p.DecodeString()
	p.HWID = p.DecodeString()
	return p
}

type CheckDuplicateIDPacket struct {
	inPacket
	CharacterName string
}

func NewCheckDuplicateIDPacket(data []byte) *CheckDuplicateIDPacket {
	p := &CheckDuplicateIDPacket{inPacket: newInPacket(data)}
	p.CharacterName = p.DecodeString()
	return p
}

type CharPacket struct {
	inPacket
	CharacterName       string
	KeySettingType      uint32
	EventNewCharSaleJob uint32 // 0xFF FF FF FF
	Class               uint32
	Job                 uint16
	Gender              bool
	SkinColor           uint8
	Face                uint32
	Hair                uint32
	SpecialFace         uint32
	Hat                 uint32
	Top                 uint32
	Bottom              uint32
	Overall             uint32
	Cape                uint32
	Shoes               uint32
	Gloves              uint32
	Weapon              uint32
	SubWeapon           uint32
}

func NewCharPacket(data []byte) *CharPacket {
	p := &CharPacket{inPacket: newInPacket(data)}
	p.CharacterName = p.DecodeLocalString()
	p.KeySettingType = p.DecodeUint32()
	p.EventNewCharSaleJob = p.DecodeUint32()
	p.Class = p.DecodeUint32()
	p.Job = p.DecodeUint16()
	p.Gender = p.DecodeBool()
	p.SkinColor = p.DecodeByte()
	itemLength := p.DecodeByte()
	items := make([]uint32, itemLength)
	for i := 0; i < int(itemLength); i++ {
		items[i] = p.DecodeUint32()
	}
	p.Face = items[0]
	p.Hair = items[1]
	// SpecialFace or Hat
	index := 2
	classID := class.ClassID(p.Class)
	if classID == class.HoYoung || classID == class.Ark ||
		classID == class.Demon || classID == class.Xenon || classID == class.BeastTamer {
		p.SpecialFace = items[index]
		index++
	} else if classID == class.Pathfinder || classID == class.Hayato || classID == class.Kanna {
		p.Hat = items[index]
		index++
	}
	// Ears and Tail
	if classID == class.BeastTamer {
		p.Hat = items[index]
		p.Cape = items[index+1]
		index += 2
	}
	// TopBottom or Overall
	if classID == class.Aran || classID == class.Evan ||
		classID == class.EunWol || classID == class.Mikhail {
		p.Top = items[index]
		p.Bottom = items[index+1]
		index += 2
	} else {
		p.Overall = items[index]
		index++
	}
	// Cape
	if classID == class.HoYoung || classID == class.Phantom ||
		classID == class.Luminous || classID == class.EunWol ||
		classID == class.KnightCygnus || classID == class.Zero {
		p.Cape = items[index]
		index++
	}
	// Shoes
	p.Shoes = items[index]
	index++
	// Gloves
	if classID == class.Hayato || classID == class.Kanna {
		p.Gloves = items[index]
		index++
	}
	// Weapon
	p.Weapon = items[index]
	index++
	// SubWeapon
	if classID == class.Demon || classID == class.Zero {
		p.SubWeapon = items[index]
	}
	return p
}

type DeleteCharPacket struct {
	inPacket
	SecondPassword string
	CharacterID    uint32
}

func NewDeleteCharPacket(data []byte) *DeleteCharPacket {
	p := &DeleteCharPacket{inPacket: newInPacket(data)}
	p.SecondPassword = p.DecodeString()
	p.CharacterID = p.DecodeUint32()
	return p
}

type GenderSetRequest struct {
	inPacket
	Username       string
	SecondPassword string
	Gender         bool
}

func NewGenderSetRequest(data []byte) *GenderSetRequest {
	p := &GenderSetRequest{inPacket: newInPacket(data)}
	p.Username = p.DecodeString()
	p.SecondPassword = p.DecodeString()
	p.Gender = p.DecodeBool()
	return p
}

type BackToLoginScreen struct {
	inPacket
	Username string
}

func NewBackToLoginScreen(data []byte) *BackToLoginScreen {
	p := &BackToLoginScreen{inPacket: newInPacket(data)}
	p.Username = p.DecodeString()
	return p
}

type ChangeCharOrderRequest struct {
	inPacket
	AccountID    uint32
	CharacterIDs []uint32
}

func NewChangeCharOrderRequest(data []byte) *ChangeCharOrderRequest {
	p := &ChangeCharOrderRequest{inPacket: newInPacket(data)}
	p.AccountID = p.DecodeUint32()
	p.Skip(1)
	characterLength := int(p.DecodeUint32())
	p.CharacterIDs = make([]uint32, 0)
	for i := 0; i < characterLength; i++ {
		p.CharacterIDs = append(p.CharacterIDs, p.DecodeUint32())
	}
	return p
}

type SelectCharacterRequest struct {
	inPacket
	CharacterID       uint32
	IsInvisibleOnline bool
}

func NewSelectCharacterRequest(data []byte) *SelectCharacterRequest {
	p := &SelectCharacterRequest{inPacket: newInPacket(data)}
	p.CharacterID = p.DecodeUint32()
	p.IsInvisibleOnline = p.DecodeBool()
	return p
}
