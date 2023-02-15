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
	JHS    JHS    `yaml:"jhs"`
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
	Retry    int    `yaml:"retry"`
}

type JHS struct {
	Token string `yaml:"token"`
}

func NewConfig(path, name string) (*Config, string) {
	logrus.Debugf("检查手动指定的配置文件信息: %s/%s", path, name)

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

	logrus.Debugf("读取到的配置文件绝对路径: %v", viper.ConfigFileUsed())

	return &config, viper.ConfigFileUsed()
}
