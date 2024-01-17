package model

const cardTableName = "card"

type Card struct {
	ID         uint   `json:"id" gorm:"column: id"`
	CardNumber string `json:"card_number" gorm:"card_number"`
	UserID     string `json:"user_id" gorm:"column: user_id"`
}

func (Card) TableName() string {
	return cardTableName
}
