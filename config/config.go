package config

import "time"

type ServerConfig struct {
	Name           string        `mapstructure:"name" json:"name"`
	Host           string        `mapstructure:"host" json:"host"`
	Port           int           `mapstructure:"port" json:"port"`
	Tags           []string      `mapstructure:"tags" json:"tags"`
	RunMode        string        `mapstructure:"run_mode" json:"run_mode"`
	LogLevel       string        `mapstructure:"log_level" json:"log_level"`
	ReadTimeout    time.Duration `mapstructure:"read_timeout" json:"read_timeout"`
	WriteTimeout   time.Duration `mapstructure:"write_timeout" json:"write_timeout"`
	ContextTimeout time.Duration `mapstructure:"context_timeout" json:"context_timeout"`

	JWTInfo    JWTConfig    `mapstructure:"jwt" json:"jwt"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
	JaegerInfo JaegerConfig `mapstructure:"jaeger" json:"jaeger"`

	LecturerSrvInfo LecturerSrvConfig `mapstructure:"lecturer_srv" json:"lecturer_srv"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Database string `mapstructure:"database" json:"database"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

type JWTConfig struct {
	SignatureKey string `mapstructure:"key" json:"key"`
	ExpiresAt    int64  `mapstructure:"expire" json:"expire"`
	Issuer       string `mapstructure:"issuer" json:"issuer"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type JaegerConfig struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type LecturerSrvConfig struct {
	Name string `mapstructure:"name" josn:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
