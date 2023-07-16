package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var serviceConfig *HTTPServiceConfig

func loadServiceConfig() *HTTPServiceConfig {
	parseConfigFile()
	return serviceConfig
}

func parseConfigFile() {
	fmt.Println("PARSING 🚨")
	var configFilePath string

	a := &HTTPServiceConfig{}

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
