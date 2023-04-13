package config

type Websocket struct {
	WorkerPoolSize int `mapstructure:"workerPoolSize" json:"workerPoolSize" yaml:"workerPoolSize"`
	MaxWorkTaskLen int `mapstructure:"maxWorkTaskLen" json:"maxWorkTaskLen" yaml:"maxWorkTaskLen"`
	MaxConnLen     int `mapstructure:"maxConnLen" json:"maxConnLen" yaml:"maxConnLen"`
}
