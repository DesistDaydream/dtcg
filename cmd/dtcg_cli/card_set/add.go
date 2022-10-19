package cardset

import (
	"sort"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func AddCardSetCommand() *cobra.Command {
	AddCardSetCmd := &cobra.Command{
		Use:   "add",
		Short: "添加卡片集合",
		Run:   addCardSet,
	}

	return AddCardSetCmd
}

func addCardSet(cmd *cobra.Command, args []string) {
	var cardSets models.CardSets

	client := cdb.NewCdbClient(core.NewClient(""))
	series, err := client.GetSeries()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, serie := range series.Data {
		for _, pack := range serie.SeriesPack {
			if pack.Language == "chs" {
				d := &models.CardSet{
					SeriesID:        int(serie.SeriesID),
					SeriesName:      serie.SeriesName,
					Language:        pack.Language,
					PackCover:       pack.PackCover,
					PackEnName:      pack.PackEnName,
					PackID:          pack.PackID,
					PackJapName:     pack.PackJapName,
					PackName:        pack.PackName,
					PackPrefix:      pack.PackPrefix,
					PackReleaseDate: pack.PackReleaseDate,
					PackRemark:      pack.PackRemark,
				}

				cardSets.Data = append(cardSets.Data, *d)

			}
		}
	}
	sort.Slice(cardSets.Data, func(i, j int) bool {
		return cardSets.Data[i].PackID < cardSets.Data[j].PackID
	})

	for _, pack := range cardSets.Data {
		logrus.WithFields(logrus.Fields{
			"前缀":     pack.PackPrefix,
			"名称":     pack.PackName,
			"PackID": pack.PackID,
			"发布时间":   pack.PackReleaseDate,
		}).Infof("%v 中的卡包信息", pack.SeriesName)

		database.AddCardSet(&pack)
	}
}
