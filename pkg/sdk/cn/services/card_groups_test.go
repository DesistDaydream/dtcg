package services

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetCardPackage(t *testing.T) {
	cardPackageResp, err := GetCardGroups()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, p := range cardPackageResp.List {
		logrus.WithFields(logrus.Fields{
			"名称": p.Name,
			"ID": p.ID,
			"状态": p.State,
		}).Info("卡包信息")
	}
}
