package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port        string `json:"port"`
	MongoDbName string `json:"mongo_db_name"`
	MongoURI    string `json:"mongo_uri"`
}

var cfg Config

func Init(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	return cfg
}
