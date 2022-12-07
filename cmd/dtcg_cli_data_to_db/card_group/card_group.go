package cardgroup

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"

	databasepkg "github.com/DesistDaydream/dtcg/pkg/database"
	serviceofficial "github.com/DesistDaydream/dtcg/pkg/sdk/cn/services"
	"github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/core"
	servicesdb "github.com/DesistDaydream/dtcg/pkg/sdk/dtcg_db/services/cdb"
	"github.com/sirupsen/logrus"
)

func AddCardGroupFromOfficial(wirteToJSON bool) {
	cardPackageResp, err := serviceofficial.GetCardGroups()
	if err != nil {
		logrus.Fatalln(err)
	}

	sort.Slice(cardPackageResp.List, func(i, j int) bool {
		return cardPackageResp.List[i].CreateTime < cardPackageResp.List[j].CreateTime
	})

	if wirteToJSON {
		jsonByte, _ := json.Marshal(cardPackageResp)

		fileName := filepath.Join("cards", "card_package.json")
		os.WriteFile(fileName, jsonByte, 0666)
	}

	for _, cardGroup := range cardPackageResp.List {
		g := &databasepkg.CardGroupFromOfficial{
			OfficialID: cardGroup.ID,
			Name:       cardGroup.Name,
			Image:      cardGroup.Image,
			State:      cardGroup.State,
			Position:   cardGroup.Position,
			CreateTime: cardGroup.CreateTime,
			UpdateTime: cardGroup.UpdateTime,
		}
		databasepkg.AddCardGroupFromOfficial(g)
	}
}

func AddCardGroupFromDtcgDB() {
	var cardGroupsFromDtcgDB databasepkg.CardGroupsFromDtcgDB

	client := servicesdb.NewCdbClient(core.NewClient("", 1))
	series, err := client.GetSeries()
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, serie := range series.Data {
		for _, pack := range serie.SeriesPack {
			if pack.Language == "chs" {
				d := &databasepkg.CardGroupFromDtcgDB{
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

				cardGroupsFromDtcgDB.Data = append(cardGroupsFromDtcgDB.Data, *d)

			}
		}
	}
	sort.Slice(cardGroupsFromDtcgDB.Data, func(i, j int) bool {
		return cardGroupsFromDtcgDB.Data[i].PackID < cardGroupsFromDtcgDB.Data[j].PackID
	})

	for _, pack := range cardGroupsFromDtcgDB.Data {
		logrus.WithFields(logrus.Fields{
			"前缀":     pack.PackPrefix,
			"名称":     pack.PackName,
			"PackID": pack.PackID,
			"发布时间":   pack.PackReleaseDate,
		}).Infof("%v 中的卡包信息", pack.SeriesName)

		databasepkg.AddCardGroupFromDtcgDB(&pack)
	}
}
