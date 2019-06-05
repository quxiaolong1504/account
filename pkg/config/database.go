package config

import "fmt"

type MySQLConfig struct {
	Engine   string `toml:"engine"`
	Host     string `toml:"host"`
	Port     uint16  `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
	SubFix	 string `toml:"subFix"`
	MaxOpenConns uint64 `toml:"MaxOpenConns"`
}

func (m *MySQLConfig) Uri () string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.User, m.Password, m.Host, m.Port,m.Name, m.SubFix)
}


type DBConfig struct {
	Master *MySQLConfig   `toml:"master"`
	Slaves []*MySQLConfig `toml:"slaves"`
	Redis *RedisConfig 	`toml:"redis"`
}
