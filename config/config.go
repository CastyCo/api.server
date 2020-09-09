package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type ConfMap struct {
	App struct {
		Version string `yaml:"version"`
		Debug   bool   `yaml:"debug"`
		Env     string `yaml:"env"`
	} `yaml:"app"`
	Grpc struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"grpc"`
	Http struct {
		AccessControlAllowOrigin string `yaml:"access_control_allow_origin"`
	} `yaml:"http"`
	Secrets struct {
		SentryDsn      string `yaml:"sentry_dsn"`
		HcaptchaSecret string `yaml:"hcaptcha_secret"`
	} `yaml:"secrets"`
	StoragePath string `yaml:"storage_path"`
}

var Map = new(ConfMap)

func Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open config file: %v", err)
	}
	if err := yaml.NewDecoder(file).Decode(&Map); err != nil {
		return fmt.Errorf("could not decode config file: %v", err)
	}
	log.Printf("ConfigMap Loaded: [version: %s]", Map.App.Version)
	return nil
}