package config

type DBConfig struct {
	DBURI  string
	DBName string
}

func defaultDBConfig() DBConfig {
	return DBConfig{
		DBURI:  "mongodb://localhost:27017/?replicaSet=rs0",
		DBName: "goms",
	}
}
