package config

type JWT struct {
	SigningKey      string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
	ExpiresTime     int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"`
	BufferTime      int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`
	UserIdName      string `mapstructure:"user_id_name" json:"user_id_name" yaml:"user_id_name"`
	TokenHeaderName string `mapstructure:"token_header_name" yaml:"token_header_name" json:"token_header_name"`
}
