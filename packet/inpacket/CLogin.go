package inpacket

import (
	"goms/maple"
	"net"
)

// MachineId(16)=MAC(6)+UNK(10)
// GameRoomClient(4)?
// NexonIDTab(1)? 0(CMS)/1(KMS)/2(GMS)
// Username "5 0 97 100 109 105 110" admin
// Password "5 0 97 100 109 105 110" admin
type CheckLoginAuthInfo struct {
	inPacket
	MachineID      []byte
	GameRoomClient int32
	NexonIDTab     uint8
	Username       string
	Skip           int16
	Password       string
}

// CLogin::CheckLoginAuthInfo
func NewCheckLoginAuthInfo(data []byte) *CheckLoginAuthInfo {
	p := &CheckLoginAuthInfo{inPacket: newInPacket(data)}
	p.MachineID = p.DecodeBuffer(16)
	p.GameRoomClient = p.DecodeInt32()
	p.NexonIDTab = p.DecodeByte()
	p.Username = p.DecodeStr()
	p.Skip = p.DecodeInt16() // UNK
	p.Password = p.DecodeStr()
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
	p.Gender = p.DecodeBool()
	p.Username = p.DecodeStr()
	p.SecondPassword = "123456" // CMS138 can't input second password
	return p
}

type SelectWorldRequest struct {
	inPacket
	AuthType     uint8
	WorldID      uint8
	ChannelIndex uint8
	ClientIP     string
}

func NewSelectWorldRequest(data []byte) *SelectWorldRequest {
	p := &SelectWorldRequest{inPacket: newInPacket(data)}
	p.AuthType = p.DecodeByte()
	p.WorldID = p.DecodeByte()
	p.ChannelIndex = p.DecodeByte()
	p.ClientIP = net.IP(p.DecodeBuffer(4)).String()
	return p
}

type CheckDuplicateIDPacket struct {
	inPacket
	CharacterName string
}

func NewCheckDuplicateIDPacket(data []byte) *CheckDuplicateIDPacket {
	p := &CheckDuplicateIDPacket{inPacket: newInPacket(data)}
	p.CharacterName = p.DecodeLocalStr()
	return p
}

type CharPacket struct {
	inPacket
	CharacterName       string
	KeySettingType      uint32
	EventNewCharSaleJob uint32 // 0xFF FF FF FF
	JobClass            uint32
	Job                 uint16
	Gender              bool
	SkinColor           uint8
	Face                uint32
	Hair                uint32
	DefFaceAcc          uint32
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
	p.CharacterName = p.DecodeStr()
	p.KeySettingType = p.DecodeUint32()
	p.EventNewCharSaleJob = p.DecodeUint32()
	p.JobClass = p.DecodeUint32()
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
	if p.JobClass == maple.ClassDemon || p.JobClass == maple.ClassXenon || p.JobClass == maple.ClassBeastTamer {
		p.DefFaceAcc = items[index]
		index++
	} else if p.JobClass == maple.ClassHayato || p.JobClass == maple.ClassKanna {
		p.Hat = items[index]
		index++
	}
	// Ears and Tail
	if p.JobClass == maple.ClassBeastTamer {
		p.Hat = items[index]
		p.Cape = items[index+1]
		index += 2
	}
	// TopBottom or Overall
	if p.JobClass == maple.ClassAran || p.JobClass == maple.ClassEvan ||
		p.JobClass == maple.ClassEunWol || p.JobClass == maple.ClassMikhail {
		p.Top = items[index]
		p.Bottom = items[index+1]
		index += 2
	} else {
		p.Overall = items[index]
		index++
	}
	// Cape
	if p.JobClass == maple.ClassPhantom || p.JobClass == maple.ClassLuminous || p.JobClass == maple.ClassEunWol ||
		p.JobClass == maple.ClassKnightCygnus || p.JobClass == maple.ClassZero {
		p.Cape = items[index]
		index++
	}
	// Shoes
	p.Shoes = items[index]
	index++
	// Gloves
	if p.JobClass == maple.ClassHayato || p.JobClass == maple.ClassKanna {
		p.Gloves = items[index]
		index++
	}
	// Weapon
	p.Weapon = items[index]
	index++
	// SubWeapon
	if p.JobClass == maple.ClassDemon || p.JobClass == maple.ClassZero {
		p.SubWeapon = items[index]
	}
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

type DeleteCharPacket struct {
	inPacket
	SecondPassword string
	CharacterID    uint32
}

func NewDeleteCharPacket(data []byte) *DeleteCharPacket {
	p := &DeleteCharPacket{inPacket: newInPacket(data)}
	p.SecondPassword = p.DecodeStr()
	p.CharacterID = p.DecodeUint32()
	return p
}

type ReservedDeleteCharacterCancelStep struct {
	inPacket
	CharacterID uint32
}

func NewReservedDeleteCharacterCancelStep(data []byte) *ReservedDeleteCharacterCancelStep {
	p := &ReservedDeleteCharacterCancelStep{inPacket: newInPacket(data)}
	p.CharacterID = p.DecodeUint32()
	return p
}

type ReservedDeleteCharacterConfirmStep struct {
	inPacket
	CharacterID    uint32
	SecondPassword string
}

func NewReservedDeleteCharacterConfirmStep(data []byte) *ReservedDeleteCharacterConfirmStep {
	p := &ReservedDeleteCharacterConfirmStep{inPacket: newInPacket(data)}
	p.CharacterID = p.DecodeUint32()
	p.SecondPassword = p.DecodeStr()
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
