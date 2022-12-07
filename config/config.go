package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Listen string `yaml:"listen"`
	Mysql  MySQL  `yaml:"mysql"`
	SQLite SQLite `yaml:"sqlite"`
	DtcgDB DtcgDB `yaml:"dtcgDB"`
}

type MySQL struct {
	Server   string `yaml:"server"`
	Password string `yaml:"password"`
}

type SQLite struct {
	FilePath string `yaml:"filePath"`
}

type DtcgDB struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Token    string `yaml:"token"`
}

func NewConfig(path, name string) *Config {
	var config Config
	viper.AddConfigPath(path)
	viper.AddConfigPath("/etc/dtcg")
	viper.AddConfigPath("./config")
	viper.SetConfigName(name)
	viper.SetConfigName("my_dtcg.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("读取配置文件失败: %v", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Fatalf("解析配置文件失败: %v", err)
	}

	return &config
}
