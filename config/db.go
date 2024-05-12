package config

type DBConfig struct {
	DBURI     string
	DBName    string
	DBTimeout uint8
}

func defaultDBConfig() DBConfig {
	return DBConfig{
		DBURI:     "mongodb://localhost:27017",
		DBName:    "goms",
		DBTimeout: 10,
	}
}
