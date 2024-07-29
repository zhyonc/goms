package manager

import (
	"goms/network"
	"log/slog"
)

type questManager struct {
	client network.ChannelClient
}

func NewQuestManager(client network.ChannelClient) QuestManager {
	m := &questManager{
		client: client,
	}
	return m
}

// Dispose implements QuestManager.
func (m *questManager) Dispose() {
	slog.Debug("Dispose")
}
