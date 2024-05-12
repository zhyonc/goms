package config

import "goms/maple/world"

type ChannelConfig struct {
	Logger        LoggerConfig
	DB            DBConfig
	ChannelServer ChannelServerConfig
}

func NewChannelConfig(path string) ChannelConfig {
	var conf ChannelConfig
	if !loadConf(path, &conf) {
		conf = defaultChannelConfig()
		saveConf(path, conf)
	}
	return conf
}

func defaultChannelConfig() ChannelConfig {
	return ChannelConfig{
		Logger:        defaultLoggerConfig(),
		DB:            defaultDBConfig(),
		ChannelServer: defaultChannelServerConfig(),
	}
}

type ChannelServerConfig struct {
	WorldID               world.WorldID
	IP                    string
	Ports                 []int
	UDPAddr               string
	MaxCapacityPerChannel uint16
}

func defaultChannelServerConfig() ChannelServerConfig {
	return ChannelServerConfig{
		WorldID:               world.Alia,
		IP:                    "127.0.0.1",
		Ports:                 []int{8501, 8502, 8503, 8504, 8505},
		UDPAddr:               "127.0.0.1:8500",
		MaxCapacityPerChannel: 1000,
	}
}
