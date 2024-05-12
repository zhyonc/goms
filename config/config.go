package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func loadConf(path string, conf any) bool {
	_, err := toml.DecodeFile(path, conf)
	if err != nil {
		log.Printf("Error load config from %s\n: %v", path, err)
		return false
	}
	return true
}

func saveConf(path string, conf any) {
	file, err := os.Create(path)
	if err != nil {
		log.Printf("Error create config to %s\n", path)
		return
	}
	defer file.Close()
	encoder := toml.NewEncoder(file)
	err = encoder.Encode(conf)
	if err != nil {
		log.Printf("Error encode config: %v\n", err)
		return
	}
	log.Printf("Save config ok")
}
