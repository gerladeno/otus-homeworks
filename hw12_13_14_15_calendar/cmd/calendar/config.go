package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gerladeno/otus_homeworks/hw12_13_14_15_calendar/cmd"
)

type Config struct {
	Logger  cmd.LoggerConf
	Storage cmd.StorageConf
	HTTP    HTTPConf
	GRPC    GRPCConf
}

type HTTPConf struct {
	Port int `json:"port"`
}

type GRPCConf struct {
	Port    int    `json:"port"`
	Network string `json:"network"`
}

func NewConfig(path string) Config {
	if path == "" {
		path = filepath.Join("configs", "config.json")
	}
	configJSON, err := ioutil.ReadFile(path)
	if err != nil {
		return defaultConfig()
	}
	config := Config{}
	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		return defaultConfig()
	}
	return config
}

func defaultConfig() Config {
	log.Print("failed to config properly, using default settings...")
	return Config{
		Logger:  cmd.LoggerConf{Level: "Debug", Path: "stdout"},
		Storage: cmd.StorageConf{Remote: false},
		HTTP:    HTTPConf{Port: 3000},
		GRPC:    GRPCConf{Port: 3005, Network: "tcp"},
	}
}
