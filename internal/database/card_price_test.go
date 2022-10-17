package database

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetCardPrice(t *testing.T) {
	dbInfo := &DBInfo{
		FilePath: "my_dtcg.db",
	}
	InitDB(dbInfo)
	got, err := GetCardPrice("2210")
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Info(got)
}
