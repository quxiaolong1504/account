package config

import "fmt"

type RedisConf struct {
	Host string `toml:"host"`
	Port uint16 `toml:"port"`
	DB uint8 `toml:"db"`
	Password string `toml:"password"`
}

func (m *RedisConf) Addr() string {
	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}

type RedisConfig struct {
	Cache *RedisConf `toml:"cache"`
}