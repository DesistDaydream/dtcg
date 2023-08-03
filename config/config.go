package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Listen         string  `yaml:"listen"`
	TokenExpiresAt string  `yaml:"tokenExpiresAt"` // Token 过期时间
	Mysql          MySQL   `yaml:"mysql"`
	SQLite         SQLite  `yaml:"sqlite"`
	Moecard        Moecard `yaml:"moecard"`
	JHS            JHS     `yaml:"jhs"`
}

type MySQL struct {
	Server   string `yaml:"server"`
	Password string `yaml:"password"`
}

type SQLite struct {
	FilePath string `yaml:"filePath"`
}

type Moecard struct {
	Retry int `yaml:"retry"`
}

type JHS struct {
	AutoUpdateTokenDuration string `yaml:"autoUpdateTokenDuration"` // 每次更新集换社 Token 的间隔时间
}

var Conf *Config

func NewConfig(path, name string) (*Config, string) {
	logrus.Debugf("检查手动指定的配置文件信息: %s/%s", path, name)

	if name == "" {
		name = "my_dtcg.yaml"
	}

	var config Config
	viper.AddConfigPath(path)
	viper.AddConfigPath("/etc/dtcg")
	viper.AddConfigPath("./config")
	viper.SetConfigName(name)
	// viper.SetConfigName("my_dtcg.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("读取配置文件失败: %v", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Fatalf("解析配置文件失败: %v", err)
	}

	logrus.Debugf("读取到的配置文件绝对路径: %v", viper.ConfigFileUsed())

	Conf = &config

	fmt.Println(Conf)

	return &config, viper.ConfigFileUsed()
}
