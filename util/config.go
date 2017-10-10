package util

import (
	"encoding/json"
	"os"

	"github.com/zcytop/Gomq/log"
)

// PortsConfig configs ports
type PortsConfig struct {
	HTTP string `json:"http"`
	UDP  string `json:"udp"`
}

// Config is the common configuration
type Config struct {
	Ports PortsConfig `json:"ports"`
}

// GetConfig get the configuration from config flle
func GetConfig() Config {
	file, err := os.Open("/home/cy/go/src/github.com/zcytop/Gomq/config.json")
	decoder := json.NewDecoder(file)
	if err != nil {
		log.Fatal("Config file error, ", err)
	}
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Some error occurs when decode config file, ", err)
	}
	return config
}
