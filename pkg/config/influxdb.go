package config

import "fmt"

type InfluxDBConfig struct {
	Protocol string `toml:"protocol"`
	Host string `toml:"host"`
	Port int16 `toml:"port"`
	UserName string `toml:"username"`
	Password string `toml:"password"`
}

func (i *InfluxDBConfig) Uri () string {
	return fmt.Sprintf("%s://%s:%d", i.Protocol, i.Host, i.Port)
}