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

	got, err := ListCardGroups()
	if err != nil {
		log.Fatalln(err)
	}

	for _, g := range got {
		log.Println(g.Name)
	}
}
