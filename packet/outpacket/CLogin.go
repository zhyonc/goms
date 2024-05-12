package outpacket

import (
	"fmt"
	"goms/maple"
	"goms/maple/class"
	"goms/maple/exp"
	"goms/maple/job"
	"goms/maple/part"
	"goms/maple/tip"
	"goms/maple/world"
	"goms/mongodb/model"
	"goms/mongodb/model/character"
	"goms/opcode"
	"goms/util"
	"time"
)

func NewChooseGender() []byte {
	p := newOutPacket(uint16(opcode.OnChooseGender))
	p.EncodeBool(true) // Open gender UI
	return p.buf
}

func NewGenderSetResult(isSet bool) []byte {
	p := newOutPacket(uint16(opcode.OnGenderSetResult))
	// false will show "try again" tip
	// true will send CheckLoginAuthInfo again
	p.EncodeBool(isSet)
	return p.buf
}

func NewLoginAuthFailed(result tip.Tip) []byte {
	p := newOutPacket(uint16(opcode.CheckLoginAuthInfo))
	p.EncodeUint16(uint16(result))
	p.EncodeUint32(0)
	return p.buf
}

func NewLoginAuthBanned(account *model.Account) []byte {
	p := newOutPacket(uint16(opcode.OnCheckPasswordResult))
	p.EncodeByte(byte(tip.Banned))
	p.EncodeString("")
	p.EncodeBool(true) // true show banned tip or false show default tip
	// Banned time
	if account.IsForeverBanned {
		p.EncodeInt64(-1)
	} else {
		p.EncodeInt64(util.UnixToFT(account.TempBannedExpireTime)) // windows file time
	}
	p.EncodeLocalString(account.BannedReason) // additional reason
	return p.buf
}

func NewLoginAuthSuccess(account *model.Account) []byte {
	p := newOutPacket(uint16(opcode.OnAccountInfoResult))
	p.EncodeByte(byte(tip.Success))
	p.EncodeString("")
	p.EncodeInt32(0)
	p.EncodeBool(account.IsGM) // AccountType
	p.EncodeInt32(0)           // GM Stuff AccoutType mask
	p.EncodeInt32(0)           // UNK
	p.EncodeInt32(0)           // UNK
	p.EncodeInt32(0)           // 0x25 maybe Age?
	p.EncodeByte(0)            // 0x03 pBlockReason
	if account.ChatUnblockTime.After(time.Now()) {
		p.EncodeBool(true)                                    // enable chat block
		p.EncodeInt64(util.UnixToFT(account.ChatUnblockTime)) // chat block expire time
	} else {
		p.EncodeBool(false) // disable chat block
		p.EncodeInt64(0)
	}
	p.EncodeBool(false) // what type block?
	p.EncodeInt64(0)    // UNK block expire time?
	p.EncodeString(account.Username)
	p.EncodeString("") // getCensoredNxLoginID()
	p.EncodeString("") // UNK
	enableClass := len(class.EnableClassCreation) > 0
	p.EncodeBool(enableClass)
	if enableClass {
		p.EncodeByte(job.JobOrder)
		for i := 0; i < 28; i++ {
			_, ok := class.EnableClassCreation[class.ClassID(i)]
			p.EncodeBool(ok)
			p.EncodeUint16(1)
		}
		for i := 1000; i < 1004; i++ {
			_, ok := class.EnableClassCreation[class.ClassID(i)]
			p.EncodeBool(ok)
			p.EncodeUint16(1)
		}
	}
	p.EncodeByte(byte(account.GradeCode)) // player grade code
	p.EncodeInt32(255)                    // CLogin.nShiningStarCount
	p.EncodeBool(false)                   // IsBeginningUser?
	return p.buf
}

func NewWorldInformation(worldID world.WorldID, state world.WorldState, desc string, channelPorts []int, ballons []world.Ballon) []byte {
	p := newOutPacket(uint16(opcode.OnWorldInformation))
	worldName := world.NameMap[worldID]
	p.EncodeByte(byte(worldID))
	p.EncodeLocalString(worldName)
	p.EncodeInt32(0) //UNK
	p.EncodeInt32(0) //UNK
	p.EncodeByte(0)  //UNK
	p.EncodeByte(0)  //UNK
	p.EncodeByte(byte(state))
	p.EncodeLocalString(desc)
	p.EncodeByte(byte(len(channelPorts)))
	for i := 1; i <= len(channelPorts); i++ {
		p.EncodeLocalString(fmt.Sprintf("%s-%d", worldName, i))
		p.EncodeInt32(10)           // current online player num every channel
		p.EncodeByte(byte(worldID)) // world ID
		p.EncodeByte(byte(i - 1))   // channel index
		p.EncodeBool(false)         // is Adult channel
	}
	p.EncodeInt16(int16(len(ballons))) // ballons length
	for _, ballon := range ballons {
		p.EncodeUint16(ballon.NX)
		p.EncodeUint16(ballon.NY)
		p.EncodeLocalString(ballon.Message)
	}
	p.EncodeInt32(0)    // some offset
	p.EncodeBool(false) // connect with star planet stuff, not interested
	return p.buf
}

func NewWorldInformationEnd() []byte {
	p := newOutPacket(uint16(opcode.OnWorldInformation))
	p.EncodeByte(255)
	p.EncodeBool(false)
	p.EncodeBool(false)
	p.EncodeInt32(0)
	p.EncodeInt32(0)
	return p.buf
}

func NewSelectWorldButton(result tip.Tip, worldID uint8) []byte {
	p := newOutPacket(uint16(opcode.OnSelectWorldButton))
	p.EncodeByte(byte(result))
	p.EncodeString("")
	p.EncodeInt32(int32(worldID))
	p.EncodeInt32(int32(world.Normal))
	return p.buf
}

func NewSelectWorldResult(result tip.Tip, chars []*character.Character, worldID world.WorldID, disableCreateChar bool, renameCharEventStartTime time.Time, renameCharEventEndTime time.Time) []byte {
	p := newOutPacket(uint16(opcode.OnSelectWorldResult))
	p.EncodeByte(byte(result))
	p.EncodeString("")
	if result != tip.Success {
		return p.buf
	}
	p.EncodeBool(false)                   // ture -> Recv opcode 198 show LAB server desc
	p.EncodeInt32(1)                      // V181 str->int
	p.EncodeInt32(1)                      // V181 str->int
	p.EncodeInt32(1)                      // V240 new
	p.EncodeInt32(4)                      // char slots default 4?
	p.EncodeBool(worldID == world.Reboot) // isReboot?
	p.EncodeBool(disableCreateChar)       // true -> cant create character
	// Separate characters type
	deletedChars := make([]*character.Character, 0)
	remainChars := make([]*character.Character, 0)
	renamedChars := make([]*character.Character, 0)
	burningCharLength := 0
	for _, char := range chars {
		if char.IsDeleted {
			deletedChars = append(deletedChars, char)
		} else {
			remainChars = append(remainChars, char)
			if char.IsRenamed {
				renamedChars = append(renamedChars, char)
			}
			if char.IsBurning {
				burningCharLength++
			}
		}
	}
	// Deleted Characters
	p.EncodeInt32(int32(len(deletedChars))) // deleting char length
	p.EncodeFT(time.Now())                  // FT zero time
	for _, char := range deletedChars {
		p.EncodeUint32(char.ID)
		p.EncodeFT(char.FutureDeleteTime)
	}
	// Characters position order
	p.EncodeBool(true) // disable order
	p.EncodeInt32(int32(len(remainChars)))
	for _, char := range remainChars {
		p.EncodeInt32(int32(char.ID))
	}
	// Characters entity
	p.EncodeByte(byte(len(remainChars)))
	for _, char := range remainChars {
		AvatarData(&p, char)
	}
	p.EncodeByte(3)     // bLoginOpt SPWMode=3
	p.EncodeBool(false) // bQuerySSNOnCreateNewCharacter
	p.EncodeByte(0)     // 1?
	p.EncodeUint32(52)  // remaining empty character slots that can be created
	p.EncodeInt32(0)    // buy character slot card
	p.EncodeInt32(255)  // nEventNewCharJob
	// Rename char Event
	if renameCharEventEndTime.Before(time.Now()) {
		p.EncodeBool(false)
	} else {
		p.EncodeBool(true)
		p.EncodeInt64(util.UnixToFT(renameCharEventStartTime))
		p.EncodeInt64(util.UnixToFT(renameCharEventEndTime))
		p.EncodeInt32(int32(len(renamedChars)))
		for _, char := range renamedChars {
			p.EncodeInt32(int32(char.ID)) // buy rename card
		}
	}
	p.EncodeInt64(time.Now().Unix())      // FT zero time
	p.EncodeBool(false)                   // V160 new
	p.EncodeBool(false)                   // true -> apply rename?
	p.EncodeByte(byte(burningCharLength)) // currently burning character count
	p.EncodeBool(false)                   // V160 new
	p.EncodeInt32(0)                      // UNK
	// p.EncodeBool(false)                // V244 new
	p.EncodeInt32(0)                 // V153 new
	p.EncodeInt32(0)                 // V160 0 -> 5
	p.EncodeInt64(time.Now().Unix()) // FT zero time
	return p.buf
}

func NewCheckSPWExistResult() []byte {
	p := newOutPacket(uint16(opcode.OnCheckSPWExistResult))
	p.EncodeByte(3) // SPWMode
	p.EncodeByte(0)
	return p.buf
}

func NewCheckDuplicatedIDResult(nickname string, isDuplicateNickname bool) []byte {
	p := newOutPacket(uint16(opcode.OnCheckDuplicatedIDResult))
	p.EncodeString(nickname)
	p.EncodeBool(isDuplicateNickname)
	return p.buf
}

func NewCreateNewCharacterResult(result tip.Tip, char *character.Character) []byte {
	p := newOutPacket(uint16(opcode.OnCreateNewCharacterResult))
	p.EncodeByte(byte(result))
	if result != tip.Success {
		if result == tip.ServerLimitCreate {
			p.EncodeUint32(uint32(char.WorldID))
		}
		return p.buf
	}
	p.EncodeUint32(uint32(char.WorldID))
	AvatarData(&p, char)
	p.EncodeBool(false)
	p.EncodeInt32(0)
	return p.buf
}

func AvatarData(p *outPacket, char *character.Character) {
	// sub_7AF860 start
	CharacterStat(p, char) // sub_7A9300
	p.EncodeInt32(0)
	p.EncodeInt32(0)
	p.EncodeInt64(0)
	p.EncodeInt32(0)
	// p.EncodeInt32(0) // V244 new
	AvatarLook(p, char) // sub_7870F0
	// sub_7AF860 end
}

func CharacterStat(p *outPacket, char *character.Character) {
	p.EncodeUint32(char.ID)
	p.EncodeUint32(char.ID)
	p.EncodeUint32(uint32(char.WorldID))
	p.EncodeLocalStringBuf(char.Name, 15)
	p.EncodeBool(char.Gender)
	p.EncodeByte(0) // UNK
	p.EncodeByte(char.Look.SkinColor)
	p.EncodeUint32(char.Look.Face)
	p.EncodeUint32(char.Look.Hair)
	p.EncodeByte(char.Look.MixBaseHairColor)
	p.EncodeByte(char.Look.MixAddHairColor)
	p.EncodeByte(char.Look.MixHairBaseProb)
	p.EncodeUint32(char.Stat.Level)
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
	charJob := job.JobID(char.Job)
	if charJob == job.BeastTamer || charJob == job.PinkBean || charJob == job.Manager || charJob == job.GM || charJob == job.SuperGM {
		p.EncodeUint16(char.Skill.SP)
	} else {
		// Extend SP
		p.EncodeByte(byte(len(char.Skill.ExtendSPs)))
		for _, extendSP := range char.Skill.ExtendSPs {
			p.EncodeByte(extendSP.JobLevel)
			p.EncodeUint32(extendSP.JobSP)
		}
	}
	p.EncodeUint64(char.Stat.Exp)
	p.EncodeUint32(char.Stat.Fame)
	p.EncodeUint32(char.WeaponPoint)  // Zero job use WeaponPoint
	p.EncodeUint64(char.GachExp)      // Other job use GachExp
	p.EncodeUint64(char.PlayTimeUnix) // playtime in seconds?
	p.EncodeUint32(char.MapID)
	p.EncodeByte(char.Portal)
	p.EncodeUint16(char.SubJob)
	if char.Look.SpecialFace > 0 {
		p.EncodeUint32(char.Look.SpecialFace)
	}
	p.EncodeByte(0)   // V229 new
	p.EncodeUint64(0) // V229 new
	// CharacterTrait
	p.EncodeUint16(char.Trait.Fatigue)
	p.EncodeUint32(util.Time2YMDH(char.Trait.LastFatigueUpdateTime)) // YYYYMMDDHH
	p.EncodeUint32(exp.TraitTotalExp[char.Trait.CharismaLevel])
	p.EncodeUint32(exp.TraitTotalExp[char.Trait.InsightLevel])
	p.EncodeUint32(exp.TraitTotalExp[char.Trait.WillLevel])
	p.EncodeUint32(exp.TraitTotalExp[char.Trait.CraftLevel])
	p.EncodeUint32(exp.TraitTotalExp[char.Trait.SenseLevel])
	p.EncodeUint32(exp.TraitTotalExp[char.Trait.CharmLevel])
	// Current exp
	p.EncodeUint32(char.Trait.CharismaExp)
	p.EncodeUint32(char.Trait.InsightExp)
	p.EncodeUint32(char.Trait.WillExp)
	p.EncodeUint32(char.Trait.CraftExp)
	p.EncodeUint32(char.Trait.SenseExp)
	p.EncodeUint32(char.Trait.CharmExp)
	p.EncodeByte(0)  // CharmByCashPR?
	p.EncodeInt64(0) // LastUpdateCharmByCashPR?
	// CharacterPVP
	p.EncodeUint32(char.PVP.Exp)
	p.EncodeByte(char.PVP.Rank)
	p.EncodeUint32(char.PVP.Point)
	p.EncodeByte(char.PVP.ModeLevel)
	p.EncodeByte(char.PVP.ModeType)
	p.EncodeUint32(char.PVP.EventPoint)
	p.EncodeInt32(0) // UNK
	p.EncodeInt64(0)
	// CharacterBurning start sub_75A5C0
	p.EncodeFT(char.BurningStartTime)
	p.EncodeFT(char.BurningEndTime)
	if char.BurningType > 0 {
		p.EncodeUint32(maple.BurningMinLevel)
		p.EncodeUint32(maple.BurningMaxLevel[maple.BurningType(char.BurningType)])
	} else {
		p.EncodeInt32(0)
		p.EncodeInt32(0)
	}
	p.EncodeInt32(0) // UNK
	p.EncodeByte(char.BurningType)
	// CharacterBurning end
	// TMS Catching Fish and Big Arena
	p.EncodeBool(false) // V160 new
	p.EncodeBuf(make([]byte, 25))
	p.EncodeByte(0)
	p.EncodeByte(0)
	p.EncodeByte(0)
	p.EncodeByte(0)
	p.EncodeByte(0)
	p.EncodeInt32(0)
}

func AvatarLook(p *outPacket, char *character.Character) {
	p.EncodeBool(char.Gender)
	p.EncodeByte(char.Look.SkinColor)
	p.EncodeUint32(char.Look.Face)
	p.EncodeUint32(char.Class)
	p.EncodeByte(0) // mega? 0:1
	p.EncodeUint32(char.Look.Hair)
	equips := part.ApplyEquipPart(char.EquipInventory.Equip)
	cashEquips := part.ApplyEquipPart(char.EquipInventory.CashEquip)
	// Equip start
	for partID, itemID := range equips {
		cashItemID, ok := cashEquips[partID]
		if ok && cashItemID > 0 {
			itemID = cashItemID // cash item will replace equip item look
		}
		p.EncodeByte(byte(partID))
		p.EncodeUint32(itemID)
	}
	p.EncodeByte(255) // Equip end
	// Equip of item inventory start
	p.EncodeByte(255) // Equip of item inventory end
	// v199 - new item loop, ignores bodypart restrictions
	p.EncodeByte(255) // end
	// Totem start
	totems := char.EquipInventory.Totems
	for i := 0; i < len(totems); i++ {
		if totems[i] == 0 {
			continue
		}
		p.EncodeByte(byte(i))
		p.EncodeUint32(totems[i])
	}
	p.EncodeByte(255)                                    // Totem end
	p.EncodeUint32(char.EquipInventory.CashEquip.Weapon) // Cash Weapon ID
	p.EncodeUint32(char.EquipInventory.Equip.Weapon)     // Weapon ID
	p.EncodeUint32(char.EquipInventory.Equip.SubWeapon)  // Sub Weapon ID
	p.EncodeInt32(0)                                     // Elf Ear ID
	p.EncodeInt32(0)                                     // V220 new
	p.EncodeBool(false)                                  // nEar!=0
	// p.EncodeByte(0) // V244 new
	for i := 0; i < len(char.EquipInventory.Pet); i++ {
		p.EncodeUint32(char.EquipInventory.Pet[i].ID) // pet ID
	}
	if char.Look.SpecialFace > 0 {
		p.EncodeUint32(char.Look.SpecialFace)
	}
	p.EncodeByte(char.Look.MixBaseHairColor)
	p.EncodeByte(char.Look.MixHairBaseProb) // mix hair percent
	p.EncodeInt32(0)                        // V230 new
	p.EncodeBuf(make([]byte, 5))
}

func NewDeleteCharacterResult(characterID uint32, result tip.Tip) []byte {
	p := newOutPacket(uint16(opcode.OnDeleteCharacterResult))
	p.EncodeUint32(characterID)
	p.EncodeByte(byte(result))
	if result != tip.Success {
		return p.buf
	}
	p.EncodeBool(false)
	p.EncodeByte(0)
	return p.buf
}

func NewSelectCharacterResult(result tip.Tip, ip4 []byte, port int, characterID uint32, isInvisibleOnline bool) []byte {
	p := newOutPacket(uint16(opcode.OnSelectCharacterResult))
	p.EncodeByte(byte(result))
	p.EncodeString("")
	p.EncodeByte(0) // error code
	if result != tip.Success {
		return p.buf
	}
	p.EncodeBuf(ip4)             // ip
	p.EncodeUint16(uint16(port)) // port
	p.EncodeUint32(characterID)
	p.EncodeInt32(1) // V181 str->int
	p.EncodeInt32(1) // V181 str->int
	p.EncodeInt32(1) // V240 new
	p.EncodeInt32(0)
	p.EncodeBool(isInvisibleOnline) // bAuthenCode
	p.EncodeInt32(0)                // ulArgument
	p.EncodeByte(20)                // selected channel
	p.EncodeUint32(1000)            // every channel max player
	return p.buf
}
