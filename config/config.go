package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Listen  string  `yaml:"listen"`
	Mysql   MySQL   `yaml:"mysql"`
	SQLite  SQLite  `yaml:"sqlite"`
	Moecard Moecard `yaml:"moecard"`
	JHS     JHS     `yaml:"jhs"`
}

type MySQL struct {
	Server   string `yaml:"server"`
	Password string `yaml:"password"`
}

type SQLite struct {
	FilePath string `yaml:"filePath"`
}

type Moecard struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Retry    int    `yaml:"retry"`
}

type JHS struct {
	UserName string `yaml: username`
}

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

	return &config, viper.ConfigFileUsed()
}
