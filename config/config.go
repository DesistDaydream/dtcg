package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Listen string `yaml:"listen"`
	Mysql  MySQL  `yaml:"mysql"`
	SQLite SQLite `yaml:"sqlite"`
}

type MySQL struct {
	Server   string `yaml:"server"`
	Password string `yaml:"password"`
}

type SQLite struct {
	FilePath string `yaml:"file_path"`
}

func NewConfig() *Config {
	var config Config
	viper.AddConfigPath("/etc/dtcg")
	viper.AddConfigPath("./config")
	viper.SetConfigName("my_dtcg.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("读取配置文件失败: %v", err)
	}
	viper.Unmarshal(&config)

	return &config
}
