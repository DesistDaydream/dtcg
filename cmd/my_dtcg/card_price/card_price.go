package cardprice

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/sirupsen/logrus"
)

func AddCardPrice() {
	cardsDesc, err := database.ListCardDescFromDtcgDB()
	if err != nil {
		logrus.Errorf("%v", err)
	}

	for _, cardDesc := range cardsDesc.Data {
		logrus.Infoln(cardDesc.CardID)
	}
}
