package cards

import "github.com/spf13/cobra"

type CardsFlags struct {
	CardVersionID int
}

var cardsFlags CardsFlags

func CreateCommand() *cobra.Command {
	cardVersionCmd := &cobra.Command{
		Use:   "cards",
		Short: "卡牌",
		// PersistentPreRun: productsPersistentPreRun,
	}

	cardVersionCmd.AddCommand(
		ListCmd(),
		GetCmd(),
	)

	cardVersionCmd.PersistentFlags().IntVar(&cardsFlags.CardVersionID, "cid", 3187, "类型")

	return cardVersionCmd
}
