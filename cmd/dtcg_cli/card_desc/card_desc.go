package carddesc

import (
	"fmt"
	"strings"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cardDescCmd := &cobra.Command{
		Use:   "card-desc",
		Short: "控制卡片描述信息",
		// PersistentPreRun: vpcPersistentPreRun,
	}

	cardDescCmd.AddCommand(
		AddCardDescCommand(),
		UpdateCardDescCommand(),
	)

	return cardDescCmd
}

func getWantAddCardDesc() []models.CardSet {
	var wantCardSets []models.CardSet

	allCardSets, err := database.ListCardSets()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, cardSet := range allCardSets.Data {
		logrus.WithFields(logrus.Fields{
			"名称": cardSet.SetPrefix,
		}).Infof("卡包信息")
	}

	fmt.Printf("请选择需要下载图片的卡包，多个卡包用逗号分隔(使用 all 下载所有): ")

	var userInput string

	fmt.Scanln(&userInput)

	userInputCardSets := strings.Split(userInput, ",")
	for _, name := range userInputCardSets {
		for _, cardInfo := range allCardSets.Data {
			if cardInfo.SetPrefix == name {
				wantCardSets = append(wantCardSets, cardInfo)
			}
		}
	}

	switch userInput {
	case "all":
		return allCardSets.Data
	default:
		return wantCardSets
	}
}
