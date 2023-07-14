package config

type AppConfig struct {
	http       `mapstructure:"http" json:"http"`
	rpcservice `mapstructure:"rpcservice" json:"rpcservice"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		http:       http{port: ":8080"},
		rpcservice: rpcservice{port: "999", method: "rpc", keepalive: keepalive{max_con_age: 1, max_con_grace: 1}},
	}

}

type http struct {
	port         string `mapstructure:"port"`
	readTimeout  int    `mapstructure:"readTimeout"`
	writeTimeout int    `mapstructure:"writeTimeout"`
}

func (h *http) GetHTTPPort() string {
	return h.port
}

func (h *http) GetReadTimeout() int {
	return h.readTimeout
}

func (h *http) GetWriteTimeout() int {
	return h.writeTimeout
}

type rpcservice struct {
	port      string `mapstructure:"port"`
	method    string `mapstructure:"method"`
	keepalive `mapstructure:"keepalive"`
}

type keepalive struct {
	max_con_age   int `mapstructure:"max_con_age"`
	max_con_grace int `mapstructure:"max_con_grace"`
}

func (r *rpcservice) GetRPCServicePort() string {
	return r.port
}

func (r *rpcservice) GetRPCServiceMaxConAge() int {
	return r.keepalive.max_con_age
}
func (r *rpcservice) GetRPCServiceMaxConGrace() int {
	return r.keepalive.max_con_grace
}
