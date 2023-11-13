package config

import "github.com/jinzhu/configor"

type Config struct {
	TgBot struct {
		DevChatID int64 `yaml:"DevChatID"`
	} `yaml:"tgbot"`
}	

// NewConfig init and return project config
func NewConfig(confPath string) (Config, error) {
	var c = Config{}
	err := configor.Load(&c, confPath)
	return c, err
}
