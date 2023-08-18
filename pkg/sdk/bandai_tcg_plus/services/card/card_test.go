package card

import (
	"fmt"
	"log"
	"testing"

	"github.com/DesistDaydream/dtcg/pkg/sdk/bandai_tcg_plus/core"
)

var client *CardClient

func initTest() {
	client = NewCardClient(core.NewClient(""))
}

func TestCardClient_GetCardMetadata(t *testing.T) {
	initTest()

	got, err := client.GetCardMetadata("2")
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, cardSet := range got.Success.CardSetList {
		fmt.Println(cardSet.ID, cardSet.Name, cardSet.Number)
	}
}

func TestCardClient_GetCardList(t *testing.T) {
	initTest()
	got, err := client.GetCardList("592")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(got)

}
