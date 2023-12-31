package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

var serviceConfig *RPCServiceConfig

func loadServiceConfig() *RPCServiceConfig {
	parseConfigFile()
	return serviceConfig
}

func parseConfigFile() {
	var configFilePath string

	a := &RPCServiceConfig{}

	flag.StringVar(&configFilePath, "config", "config.yml", "Path to config file")
	flag.Parse()
	v := viper.New()
	v.SetConfigFile(configFilePath)
	if rErr := v.ReadInConfig(); rErr != nil {
		log.Printf("could not find config file: %v, error msg: %v", configFilePath, rErr)
	}

	if err := v.Unmarshal(&a); err != nil {
		log.Fatalf("error parsing config file : %v, error msg: %v", configFilePath, err)
	}

	serviceConfig = a
}
