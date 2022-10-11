package database

import (
	"log"
	"testing"
)

func TestListCardDesc(t *testing.T) {
	i := &DBInfo{
		FilePath: "my_dtcg.db",
	}
	InitDB(i)

	got, err := ListCardDesc()
	if err != nil {
		log.Fatalln(err)
	}

	for _, g := range got.Data {
		log.Println(g.Name)
	}
}

func TestListCardDescFromDtcgDB(t *testing.T) {
	i := &DBInfo{
		FilePath: "my_dtcg.db",
	}
	InitDB(i)

	got, err := ListCardDescFromDtcgDB()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(got.Count)

	// for _, g := range got.List {
	// 	log.Println(g.ScName)
	// }
}
