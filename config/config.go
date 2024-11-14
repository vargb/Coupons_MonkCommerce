package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	User     string `json:"user"`
	Dbname   string `json:"dbname"`
	Sslmode  string `json:"sslmode"`
}

type Config struct {
	Server   Server   `json:"server"`
	Postgres Postgres `json:"postgres"`
}

var logger *logrus.Logger

func GetLogger() *logrus.Logger {
	return logger
}

func InitConfig() (*Config, error) {
	loggy := logrus.New()
	logger = loggy
	loggy.Info("entering config...")
	jsonFile, err := os.Open("C:\\Users\\Dell\\monkCommerce\\src\\monkCommerce\\config\\config.json")
	if err != nil {
		loggy.Error("error in opening json file", err)
		return nil, err
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		loggy.Error("error in opening json file", err)
		return nil, err
	}

	var config *Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		loggy.Error("error in opening json file", err)
		return nil, err
	}
	return config, nil
}
