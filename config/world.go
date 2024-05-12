package config

import (
	"goms/maple/world"
	"time"
)

type WorldConfig struct {
	ID                       world.WorldID
	State                    world.WorldState
	Desc                     string
	Ballons                  []world.Ballon
	DisableCreateChar        bool
	DisableSortChar          bool
	DeleteCharWaitTime       uint32
	RenameCharEventStartTime time.Time
	RenameCharEventEndTime   time.Time
	ChannelIP                string
	ChannelPorts             []int
}

func defaultWorldConfigList() []*WorldConfig {
	worldConfigList := []*WorldConfig{
		{
			ID:                       world.Alia,
			State:                    world.Full,
			Desc:                     "Popular World",
			RenameCharEventStartTime: time.Now(),
			RenameCharEventEndTime:   time.Now().Add(168 * time.Hour),
			ChannelIP:                "127.0.0.1",
			ChannelPorts:             []int{8501, 8502, 8503, 8504, 8505},
		},
		{
			ID:                       world.Reboot,
			State:                    world.New,
			Desc:                     "New World",
			RenameCharEventStartTime: time.Now(),
			RenameCharEventEndTime:   time.Now().Add(168 * time.Hour),
			ChannelIP:                "127.0.0.1",
			ChannelPorts:             []int{8601, 8602, 8603, 8604, 8605},
		},
	}
	return worldConfigList
}
