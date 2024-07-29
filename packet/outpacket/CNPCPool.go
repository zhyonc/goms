package outpacket

import (
	"goms/nx"
	"goms/opcode"
)

// CNpcPool::OnNpcEnterField
func NewNpcEnterField(life *nx.Life) []byte {
	p := newOutPacket(opcode.CNpcPool_OnNpcEnterField)
	CNpcPoolEncodeInitPacket(&p, life)
	return p.buf
}

// CNpcPool::OnNpcChangeController
func NewNpcChangeController(isLocalNpc bool, life *nx.Life) []byte {
	p := newOutPacket(opcode.CNpcPool_OnNpcChangeController)
	p.EncodeBool(isLocalNpc)
	p.EncodeUint32(life.GetID()) // objectID: use lifeID as objectID
	if isLocalNpc {
		CNpcPoolSetLocalNpc(&p, life)
	}
	return p.buf
}
