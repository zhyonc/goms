package config

import "goms/logger"

type LoggerConfig struct {
	LogLevel          logger.LogLevel
	LogBackupPath     string
	LogBackupInterval uint16 // Seconds
}

func defaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		LogLevel:          logger.Debug,
		LogBackupPath:     "./temp/",
		LogBackupInterval: 3600,
	}
}
