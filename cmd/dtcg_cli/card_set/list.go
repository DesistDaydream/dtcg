package cardset

import (
	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func ListCardSetsCommand() *cobra.Command {
	ListCardSetsCmd := &cobra.Command{
		Use:   "list",
		Short: "添加卡片集合",
		Run:   listCardSets,
	}

	return ListCardSetsCmd
}

func listCardSets(cmd *cobra.Command, args []string) {
	allCardSets, err := database.ListCardSets()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, cardSet := range allCardSets.Data {
		logrus.WithFields(logrus.Fields{
			"名称": cardSet.SetPrefix,
			"编号": cardSet.SetID,
		}).Infof("卡包信息")
	}
}
