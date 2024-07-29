package outpacket

import (
	"goms/nx"
)

// Call by CNpcPool::OnNpcEnterField
// Call by CNpcPool::OnNpcLeaveField
// CNpcPool::DecodeInitPacket
func CNpcPoolEncodeInitPacket(p *outPacket, life *nx.Life) {
	p.EncodeUint32(life.GetID()) // objectID: use lifeID as objectID
	p.EncodeUint32(life.GetID()) // dwTemplateID is lifeID
	CNpcInit(p, life)
}

// Call by CNpcPool::DecodeInitPacket
// Call by CNpcPool::SetLocalNpc
// CNpc::Init
func CNpcInit(p *outPacket, life *nx.Life) {
	p.EncodeInt16(life.X)                      // m_ptPos
	p.EncodeInt16(life.Y)                      // m_ptPos
	p.EncodeByte(0)                            // m_bMove
	p.EncodeBool(!life.F)                      // m_nMoveAction !isFlip?
	p.EncodeInt16(life.Fh)                     // Foothold
	p.EncodeInt16(life.Rx0)                    // rgHorz.low
	p.EncodeInt16(life.Rx1)                    // rgHorz.high
	p.EncodeBool(!life.Hide)                   // m_bEnabled miniMap?
	p.EncodeInt32(0)                           // unk
	p.EncodeUint32(0)                          // nItemID
	p.EncodeByte(0)                            // m_nPresentTimeState
	p.EncodeBuffer([]byte{255, 255, 255, 255}) // m_tPresentTime
	m_nNoticeBoardType := 0
	p.EncodeUint32(uint32(m_nNoticeBoardType)) // m_nNoticeBoardType
	if m_nNoticeBoardType == 1 {
		p.EncodeUint32(0) // m_nNoticeBoardValue
	}
	p.EncodeInt32(0) // tAlpha
	p.EncodeStr("")  // sLocalRepeatEffect
	condition := false
	p.EncodeBool(condition)
	if condition {
		CScreenInfoEncode(p)
	}
}

// Call by CNpc::Init
// CScreenInfo::Decode
func CScreenInfoEncode(p *outPacket) {
	p.EncodeByte(0)  // nType
	p.EncodeInt32(0) // unk
}

// Call by CNpcPool::OnNpcChangeController
// CNpcPool::SetLocalNpc
func CNpcPoolSetLocalNpc(p *outPacket, life *nx.Life) {
	p.EncodeUint32(life.GetID()) // dwTemplateID
	CNpcInit(p, life)
}
