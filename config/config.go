package config

type Server struct {
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`          //Jwt 配置
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`          //zap 日志配置
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`    //redis 配置
	Email     Email     `mapstructure:"email" json:"email" yaml:"email"`    //邮箱配置
	Casbin    Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"` //	权限管理配置
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Snowflake Snowflake `mapstructure:"snowflake" json:"snowflake" yaml:"snowflake"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"` //mysql 配置
	// oss
	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`

	Minio Minio `mapstructure:"minio" json:"minio" yaml:"minio"` //minio开源对象存储

	WebSocket Websocket `mapstructure:"websocket" json:"websocket" yaml:"websocket"` //websocket配置
}
