package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Configs struct {
	App *App    `yaml:"server"`
	DB  *DBConf `yaml:"db"`
}

type DBConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	TimeOut  int    `yaml:"timeout"`
}

type App struct {
	AppPort         string `yaml:"port"`
	AppShutdownTime int    `yaml:"shutdown_time"`
}

func New() (*Configs, error) {
	yamlBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	config := &Configs{}
	err = yaml.Unmarshal(yamlBytes, config)
	if err != nil {
		log.Fatal(err)
	}
	return config, err
}
