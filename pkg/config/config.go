package config

import (
	"github.com/BurntSushi/toml"
	"github.com/quxiaolong/account/pkg/utils/logger"
	"os"
	"path/filepath"
)

type Config struct {
	DataBase *DBConfig `toml:"database"`
	Redis *RedisConfig `toml:"redis"`
	Auth *AuthConfig `toml:"auth"`
}

var Conf *Config

func init() {
	Conf = &Config{}
}

func (c *Config) Load(confDir string) {
	c.initDefault()

	if confDir != "" {
		configs := []string{"base.toml", "local.toml", "product.toml"}
		for _, name := range(configs) {
			c.loadFromToml(confDir, name)
		}
		return
	}
	logger.Logger.Warningf("confDif: {} is empty! use default config! ", confDir)
}

func (c *Config) loadFromToml(confDir, name string) {
	path := filepath.Join(confDir, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Logger.Warningf("Skip config: %s", path)
		return
	}

	if _, err := toml.DecodeFile(path, &c); err != nil {
		logger.Logger.Warningf("bad config %s: %s", path, err.Error())
	}
}

func (c *Config) initDefault() {
	// init default config db, server, etc
	logger.Logger.Infof("init default config.")
}
