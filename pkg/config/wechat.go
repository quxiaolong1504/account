package config

type WeChatConfig struct {
	AppID string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}