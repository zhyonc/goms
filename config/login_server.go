package config

type LoginServerConfig struct {
	TCPAddr              string
	UDPAddr              string
	UDPXORKey            string
	IsTestMode           bool
	IsAutoRegister       bool
	EnableBcryptPassword bool
}

func defaultLoginServerConfig() LoginServerConfig {
	return LoginServerConfig{
		TCPAddr:              "127.0.0.1:8484",
		UDPAddr:              "127.0.0.1:8484",
		UDPXORKey:            "123456",
		IsTestMode:           true,
		IsAutoRegister:       false,
		EnableBcryptPassword: false,
	}
}
