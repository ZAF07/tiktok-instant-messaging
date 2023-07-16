package config

import (
	"errors"
	"fmt"
	"sync"
)

var appConfig *ApplicationConfig
var once sync.Once

type HTTPServiceConfig struct {
	Http `mapstructure:"http"`
	Rpc  `mapstructure:"rpc"`
}

type Http struct {
	Port         string   `mapstructure:"port"`
	ReadTimeout  int      `mapstructure:"readTimeout"`
	WriteTimeout int      `mapstructure:"writeTimeout"`
	Cors         []string `mapstructure:"cors_allow_origin"`
}

type Rpc struct {
	Port        string `mapstructure:"port"`
	Method      string `mapstructure:"method"`
	MaxConAge   int    `mapstructure:"max_con_age"`
	MaxConGrace int    `mapstructure:"max_con_grace"`
}

type ApplicationConfig struct {
	config HTTPServiceConfig
}

func InitAppConfig() *ApplicationConfig {
	once.Do(func() {
		serviceConfig := loadServiceConfig()
		appConfig = &ApplicationConfig{
			config: *serviceConfig,
		}
	})
	return appConfig
}

func GetConfig() (*ApplicationConfig, error) {
	if appConfig != nil {
		return appConfig, nil
	}
	return nil, errors.New("config is not yet initialised")
}

func (ac *ApplicationConfig) GetPortHTTP() string {
	return fmt.Sprintf(":%v", ac.config.Http.Port)
}

func (ac *ApplicationConfig) GetReadTimeoutHTTP() int {
	return ac.config.Http.ReadTimeout
}
func (ac *ApplicationConfig) GetWriteTimeoutHTTP() int {
	return ac.config.Http.WriteTimeout
}
func (ac *ApplicationConfig) GetCorsAllowOrigins() []string {
	return ac.config.Cors
}

// RPC SERVICE CONFIG
func (ac *ApplicationConfig) GetPortRPC() string {
	return ac.config.Rpc.Port
}
func (ac *ApplicationConfig) GetMethodRPC() string {
	return ac.config.Rpc.Method
}
func (ac *ApplicationConfig) GetMaxConAgeRPC() int {
	return ac.config.Rpc.MaxConAge
}
func (ac *ApplicationConfig) GetMaxConGraceRPC() int {
	return ac.config.Rpc.MaxConGrace
}
