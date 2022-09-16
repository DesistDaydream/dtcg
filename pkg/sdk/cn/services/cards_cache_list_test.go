package services

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetCardColor(t *testing.T) {
	got, err := GetCardColor()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	fmt.Println(got)
}

func TestGetCardLevel(t *testing.T) {
	got, err := GetCardLevel()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	fmt.Println(got)
}

func TestGetCardGetway(t *testing.T) {
	got, err := GetCardGetway()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	fmt.Println(got)
}
