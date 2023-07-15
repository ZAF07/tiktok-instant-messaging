package config

import "sync"

var appConfig *ApplicationConfig
var once sync.Once

type RPCServiceConfig struct {
	Rpc       `mapstructure:"rpc"`
	Datastore `mapstructure:"datastore"`
}

type Rpc struct {
	Port        string `mapstructure:"port"`
	Method      string `mapstructure:"method"`
	MaxConAge   int    `mapstructure:"max_con_age"`
	MaxConGrace int    `mapstructure:"max_con_grace"`
}

type Datastore struct {
	Environment string `mapstructure:"environment"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	User        string `mapstructure:"user"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	SSL         string `mapstructure:"ssl"`
	MaxConns    int    `mapstructure:"max_conns"`
}

type ApplicationConfig struct {
	config RPCServiceConfig
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

// DATASTORE
func (ac *ApplicationConfig) GetDatastoreName() string {
	return ac.config.Datastore.Name
}

func (ac *ApplicationConfig) GetDatastoreEnv() string {
	return ac.config.Datastore.Environment
}

func (ac *ApplicationConfig) GetDatastorePort() string {
	return ac.config.Datastore.Port
}

func (ac *ApplicationConfig) GetDatastoreHost() string {
	return ac.config.Datastore.Host
}

func (ac *ApplicationConfig) GetDatastoreUser() string {
	return ac.config.Datastore.User
}

func (ac *ApplicationConfig) GetDatastorePassword() string {
	return ac.config.Datastore.Password
}

func (ac *ApplicationConfig) GetDatastoreSSL() string {
	return ac.config.Datastore.SSL
}

func (ac *ApplicationConfig) GetDatastoreMaxCons() int {
	return ac.config.Datastore.MaxConns
}
