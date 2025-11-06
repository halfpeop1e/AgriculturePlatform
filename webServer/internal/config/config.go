package config

import "github.com/BurntSushi/toml"

type MainConfig struct {
	AppName string `toml:"appName"`
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
}
type PostgreConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Db       string `toml:"db"`
}

type StaticSrcConfig struct {
	StaticAvatarPath string `toml:"staticAvatarPath"`
	StaticFilePath   string `toml:"staticFilePath"`
}

type JwtConfig struct {
	JwtSecret     string `toml:"jwtSecret"`
	JwtExpireTime int64  `toml:"jwtExpireTime"`
}

type Config struct {
	MainConfig      `toml:"mainConfig"`
	PostgreConfig   `toml:"mysqlConfig"`
	StaticSrcConfig `toml:"staticSrcConfig"`
	JwtConfig       `toml:"jwtConfig"`
}

var config *Config

func LoadConfig() error {
	if _, err := toml.DecodeFile("./configs/config.toml", config); err != nil {
		return err
	}
	return nil
}

func GetConfig() *Config {
	if config == nil {
		config = new(Config)
		_ = LoadConfig()
	}
	return config
}
