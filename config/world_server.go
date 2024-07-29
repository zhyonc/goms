package config

import (
	"goms/maple"
	"time"
)

type WorldServerConfig struct {
	Addr                     string
	UDPXORKey                string
	ChannelPorts             []int
	OnlineLimitPerChannel    uint32
	WorldID                  maple.WorldID
	WorldTag                 maple.WorldTag
	WorldExpRate             uint16
	WorldDropRete            uint16
	WorldBallons             []maple.Ballon
	WorldRecommendMsg        string
	DeleteCharWaitTime       uint32
	RenameCharEventStartDate time.Time
	RenameCharEventEndDate   time.Time
}

func defaultWorldConfig() WorldServerConfig {
	return WorldServerConfig{
		Addr:                     "127.0.0.1:8500",
		UDPXORKey:                "654321",
		ChannelPorts:             []int{8585, 8586, 8587, 8588, 8589},
		WorldID:                  maple.BlueSnail,
		WorldTag:                 maple.NewWorld,
		WorldExpRate:             1,
		WorldDropRete:            1,
		WorldRecommendMsg:        "Popular World",
		RenameCharEventStartDate: time.Now(),
		RenameCharEventEndDate:   time.Now().Add(168 * time.Hour),
		OnlineLimitPerChannel:    500, // Max is 1200?
	}
}
