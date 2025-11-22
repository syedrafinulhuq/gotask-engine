package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerPort int `json:"server_port"`
}

func LoadConfig(path string) (*Config, error) {
	var file *os.File
	var err error
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err = json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
