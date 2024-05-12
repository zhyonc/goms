package config

type LoginConfig struct {
	Logger      LoggerConfig
	DB          DBConfig
	LoginServer LoginServerConfig
	WorldList   []*WorldConfig
}

func NewLoginConfig(path string) LoginConfig {
	var conf LoginConfig
	if !loadConf(path, &conf) {
		conf = defaultLoginConfig()
		saveConf(path, conf)
	}
	return conf
}

func defaultLoginConfig() LoginConfig {
	return LoginConfig{
		Logger:      defaultLoggerConfig(),
		DB:          defaultDBConfig(),
		LoginServer: defaultLoginServerConfig(),
		WorldList:   defaultWorldConfigList(),
	}
}

type LoginServerConfig struct {
	TCPAddr              string
	UDPAddr              string
	EnableAutoRegister   bool
	EnableBcryptPassword bool
}

func defaultLoginServerConfig() LoginServerConfig {
	return LoginServerConfig{
		TCPAddr:              "127.0.0.1:8484",
		UDPAddr:              "127.0.0.1:8484",
		EnableAutoRegister:   true,
		EnableBcryptPassword: true,
	}
}
