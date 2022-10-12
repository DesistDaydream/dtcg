package database

import (
	"log"
	"testing"
)

func TestListCardGroups(t *testing.T) {
	i := &DBInfo{
		FilePath: "my_dtcg.db",
	}
	InitDB(i)

	got, err := ListCardGroupsFromOfficial()
	if err != nil {
		log.Fatalln(err)
	}

	for _, g := range got.Data {
		log.Println(g.Name)
	}
}
