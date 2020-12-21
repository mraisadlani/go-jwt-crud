package common

import (
	"encoding/json"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"os"
)

type Configuration struct {
	Port string `json:"port"`

	LogFilename string `json:"logFilename"`
	LogMaxSize int `json:"logMaxSize"`
	LogMaxBackups int `json:"logMaxBackups"`
	LogMaxAge int `json:"logMaxAge"`

	DbUser string `json:"DBUser"`
	DbPassword string `json:"DBPass"`
	DbHost string `json:"DBHost"`
	DbPort string `json:"DBPort"`
	DbName string `json:"DBName"`
	DbDriver string `json:"DBDrive"`

	KeySecret string `json:"keySecret"`
}


var (
	Config *Configuration
)

func LoadConfig() error {
	file, err := os.Open("/home/vanilla/Vanilla/Golang/go-jwt-crud/api/config/config.json")

	if err != nil {
		return err
	}

	Config = new(Configuration)
	err = json.NewDecoder(file).Decode(&Config)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(&lumberjack.Logger{
		Filename: Config.LogFilename,
		MaxSize: Config.LogMaxSize,
		MaxBackups: Config.LogMaxBackups,
		MaxAge: Config.LogMaxAge,
	})

	log.SetLevel(log.DebugLevel)

	log.SetFormatter(&log.JSONFormatter{})

	return nil
}
