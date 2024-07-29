package outpacket

import (
	"goms/maple"
	"goms/mongodb/model/inventory"
)

const familiarMask uint32 = 0x74000000

// Call by CStage::Decode
// FamiliarSystem::Decode
func FamiliarSystemEncode(inv *inventory.Inventory, bCharacterData bool) []byte {
	// CInPacket::RawAppendBuffer
	p := newEmptyOutPacket()
	if inv.FamiliarInv.SummonedCard.ID != 0 {
		p.EncodeUint32(3) // flag?
		p.EncodeUint32(0x257B3124)
		p.EncodeUint32(familiarMask | 2)
		p.EncodeUint32(0x44)
		FamiliarCardEncode(&p, inv.ID, inv.FamiliarInv.SummonedCard, true)
		p.EncodeBuffer([]byte{189, 1, 198, 3}) // BD 01 C6 03?
		p.EncodeUint32(2000)
		p.EncodeUint32(2000)
	} else {
		p.EncodeUint32(2) // flag?
	}
	p.EncodeUint32(1392187010)
	p.EncodeUint32(familiarMask | 1)
	if bCharacterData {
		cardsLength := len(inv.FamiliarInv.Cards)
		if cardsLength > 0 {
			p.EncodeUint32(uint32(cardsLength*56 + 14))
		} else {
			p.EncodeUint32(8)
		}
		if inv.FamiliarInv.SummonedCard.ID != 0 {
			p.EncodeUint32(2)
		} else {
			p.EncodeUint32(0)
		}
		if cardsLength > 0 {
			p.EncodeByte(byte(cardsLength * 2))
			for _, card := range inv.FamiliarInv.Cards {
				FamiliarCardEncode(&p, inv.ID, card, false)
			}
			p.EncodeUint16(2)
			p.EncodeByte(6)
			p.EncodeUint16(0)
		}
	}
	p.EncodeUint16(0)
	p.EncodeUint16(0)
	p.EncodeUint32(1980835063)
	p.EncodeUint32(familiarMask)
	p.EncodeUint32(8)
	p.EncodeUint32(familiarMask | 1)
	p.EncodeUint32(0)
	return p.buf
}

func FamiliarCardEncode(p *outPacket, characterID uint32, card inventory.Familiar, bUNK bool) {
	p.EncodeUint32(characterID)
	p.EncodeUint32(0)
	p.EncodeBuffer([]byte{2, 80, 32, 100}) // 02 50 20 64?
	p.EncodeUint32(card.ID)
	p.EncodeLocalName(card.Name, maple.FamiliarCardNameLength)
	p.EncodeUint16(0)
	p.EncodeByte(0)
	p.EncodeUint16(card.Level)
	p.EncodeUint16(card.Skill)
	p.EncodeUint16(131)
	p.EncodeUint32(card.Exp)
	p.EncodeUint16(0)
	p.EncodeUint16(card.Option1)
	p.EncodeUint16(card.Option2)
	p.EncodeUint16(card.Option3)
	p.EncodeByte(0)
	if bUNK {
		p.EncodeByte(32)
		p.EncodeByte(100)
	} else {
		p.EncodeByte(0)
		p.EncodeByte(0)
	}
	p.EncodeBuffer([]byte{140, 254, 102, 21}) // 8C FE 66 15?
}
