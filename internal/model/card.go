package model

const cardTableName = "card"

type Card struct {
	ID         uint   `json:"id"          gorm:"primaryKey"`
	CardNumber string `json:"card_number" gorm:"column:card_number"`
	UserID     uint   `json:"user_id"     gorm:"column:user_id"`
}

func (Card) TableName() string {
	return cardTableName
}
