package cardset

import (
	"os"
	"strconv"

	"github.com/DesistDaydream/dtcg/internal/database"
	"github.com/olekukonko/tablewriter"
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"包名", "编号", "发布时间"})

	for _, cardSet := range allCardSets.Data {
		table.Append([]string{cardSet.SetPrefix, strconv.Itoa(cardSet.SetID), cardSet.SetReleaseDate})
	}

	table.Render()
}
