package config

type Snowflake struct {
	StartTime string `mapstructure:"start_time" json:"start_time" yaml:"start_time"`
	MachineID int64  `mapstructure:"machine_id" json:"machine_id" yaml:"machine_id"`
}
