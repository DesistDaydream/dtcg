package main

import (
	"github.com/DesistDaydream/dtcg/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	logFlags := logging.LoggingFlags{}
	logFlags.AddFlags()
	pflag.Parse()
	if err := logging.LogInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}
}
