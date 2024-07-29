package manager

import (
	"goms/network"
	"goms/packet/outpacket"
	"log/slog"
)

type npcManager struct {
	client network.ChannelClient
}

func NewNPCManager(client network.ChannelClient) NPCManager {
	m := &npcManager{
		client: client,
	}
	return m
}

// Dispose implements NPCManager.
func (m *npcManager) Dispose() {
	slog.Debug("Dispose")
}

// SendNext implements NPCManager.
func (m *npcManager) SendNext(text string) {
	slog.Debug("New script message", "text", text)
	sm := m.client.GetScriptMessage()
	if sm == nil {
		return
	}
	sm.Text = text
	sm.Prev = false
	sm.Next = true
	m.client.SendPacket(outpacket.NewScriptMessage(sm))
}

// SendNextPrev implements NPCManager.
func (m *npcManager) SendNextPrev(text string) {
	slog.Debug("New script message", "text", text)
	sm := m.client.GetScriptMessage()
	if sm == nil {
		return
	}
	sm.Text = text
	sm.Prev = true
	sm.Next = true
	m.client.SendPacket(outpacket.NewScriptMessage(sm))
}

// SendPrev implements NPCManager.
func (m *npcManager) SendPrev(text string) {
	slog.Debug("New script message", "text", text)
	sm := m.client.GetScriptMessage()
	if sm == nil {
		return
	}
	sm.Text = text
	sm.Prev = true
	sm.Next = false
	m.client.SendPacket(outpacket.NewScriptMessage(sm))
}

// Warp implements NPCManager.
func (m *npcManager) Warp(mapID int32, portalID int32) {

}
