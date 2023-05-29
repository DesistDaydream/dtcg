package jp

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestImageHandler_GetCardSets(t *testing.T) {
	i := &ImageHandler{
		Lang:      "jp_hk",
		DirPrefix: "",
	}
	i.GetCardSets()

}

func TestImageHandler_GetImagesURL(t *testing.T) {
	i := &ImageHandler{
		Lang:      "jp_hk",
		DirPrefix: "",
	}
	got, err := i.GetImagesURL("")
	if err != nil {
		logrus.Fatal(err)
	}
	for _, g := range got {
		logrus.WithFields(logrus.Fields{
			"名称": g,
		}).Info("")
	}
}
