package models

type CardList struct {
	Success CardListData `json:"success"`
}

type CardListData struct {
	Code  int    `json:"code"`
	Cards []Card `json:"cards"`
	Total string `json:"total"`
}

type Card struct {
	ID         int    `json:"id"`
	ImageURL   string `json:"image_url"`
	CardNumber string `json:"card_number"`
	CardName   string `json:"card_name"`
}
