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

	pageSize := 0
	pageNum := 0

	got, err := GetCardDescFromDtcgDB(pageSize, pageNum)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("卡片总数: %v", got.Count)
	log.Printf("总页: %v", got.PageTotal)
	log.Printf("当前页: %v", got.PageCurrent)
	for _, g := range got.Data {
		log.Println(g.ID, g.ScName)
	}
}
