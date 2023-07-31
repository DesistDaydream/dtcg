package cardset

import (
	"sort"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/handler"
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

	series, err := handler.H.MoecardServices.Cdb.GetSeries()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, serie := range series.Data {
		for _, pack := range serie.SeriesPack {
			if pack.Language == "chs" {
				d := &models.CardSet{
					SeriesID:       int(serie.SeriesID),
					SeriesName:     serie.SeriesName,
					Language:       pack.Language,
					SetCover:       pack.PackCover,
					SetEnName:      pack.PackEnName,
					SetID:          pack.PackID,
					SetJapName:     pack.PackJapName,
					SetName:        pack.PackName,
					SetPrefix:      pack.PackPrefix,
					SetReleaseDate: pack.PackReleaseDate,
					SetRemark:      pack.PackRemark,
				}

				cardSets.Data = append(cardSets.Data, *d)

			}
		}
	}
	sort.Slice(cardSets.Data, func(i, j int) bool {
		return cardSets.Data[i].SetID < cardSets.Data[j].SetID
	})

	for _, pack := range cardSets.Data {
		logrus.WithFields(logrus.Fields{
			"前缀":     pack.SetPrefix,
			"名称":     pack.SetName,
			"PackID": pack.SetID,
			"发布时间":   pack.SetReleaseDate,
		}).Infof("%v 中的卡包信息", pack.SeriesName)

		database.AddCardSet(&pack)
	}
}
