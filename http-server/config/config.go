package config

import (
	"sync"
)

var appConfig *ApplicationConfig
var once sync.Once

type HTTPServiceConfig struct {
	Http       http `mapstructure:"http"`
	Rpcservice `mapstructure:"rpcservice" json:"rpcservice"`
}

type http struct {
	Port         string `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"readTimeout"`
	WriteTimeout int    `mapstructure:"writeTimeout"`
}

type Rpcservice struct {
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

func (ac *ApplicationConfig) GetPortHTTP() string {
	return ac.config.Http.Port
}

func (ac *ApplicationConfig) GetReadTimeoutHTTP() int {
	return ac.config.Http.ReadTimeout
}
func (ac *ApplicationConfig) GetWriteTimeoutHTTP() int {
	return ac.config.Http.WriteTimeout
}

// RPC SERVICE CONFIG
func (ac *ApplicationConfig) GetPortRPC() string {
	return ac.config.Rpcservice.Port
}
func (ac *ApplicationConfig) GetMethodRPC() string {
	return ac.config.Rpcservice.Method
}
func (ac *ApplicationConfig) GetMaxConAgeRPC() int {
	return ac.config.Rpcservice.MaxConAge
}
func (ac *ApplicationConfig) GetMaxConGraceRPC() int {
	return ac.config.Rpcservice.MaxConGrace
}
