package outpacket

import (
	"goms/mongodb/model/social"
)

// Call by CWvsContext::OnPartyResult
// PARTYDATA::Decode
func PartyDataEncode(p *outPacket, soc *social.Social) {
	// CInPacket::DecodeBuffer(0xFAu)
	for _, member := range soc.Party.Members {
		p.EncodeUint32(member.CharacterID)
		p.EncodeLocalName(member.Name, 13)
		p.EncodeInt32(member.Job)
		p.EncodeInt32(member.SubJob)
		p.EncodeInt32(member.Level)
		p.EncodeInt32(0) // channelIndex
		p.EncodeInt32(0) // isOnline?
		p.EncodeInt32(0) // v138 new
	}
	p.EncodeUint32(soc.Party.LeaderID)
	// adwFieldID CInPacket::DecodeBuffer(0x18u)
	for _, member := range soc.Party.Members {
		p.EncodeUint32(member.MapID) // ?
	}
	// aTownPortal CInPacket::DecodeBuffer(0x78u)
	for _, member := range soc.Party.Members {
		p.EncodeUint32(member.TownID)
		p.EncodeUint32(member.TargetID)
		p.EncodeUint32(member.SkillID)
		p.EncodeInt32(member.PosX)
		p.EncodeInt32(member.PosY)
	}
	p.EncodeBool(soc.Party.Appliable) // bAppliable
	p.EncodeBool(false)               // v138 new
	p.EncodeStr(soc.Party.Name)       // sPartyName
}
