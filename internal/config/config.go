package config

import (
	"encoding/json"
	"os"
	"sync"
)

type config struct {
	Port     int    `json:"port"`
	Address  string `json:"address"`
	Mongo    string `json:"mongo"`
	Database string `json:"database"`
}

var (
	c    *config
	once sync.Once
)

func GetConfig() *config {
	once.Do(func() {
		data, err := os.ReadFile("config.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &c)
		if err != nil {
			panic(err)
		}
	})
	return c
}
