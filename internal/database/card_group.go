package database

type CardGroup struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func AddCardGroup(cardGroup *CardGroup) {
	db.Create(cardGroup)
}
