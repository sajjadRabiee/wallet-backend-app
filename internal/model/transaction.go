package model

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint `gorm:"primaryKey"`
	SourceOfFundID *uint
	SourceOfFund   *SourceOfFund `gorm:"foreignKey:SourceOfFundID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID         uint
	User           User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DestinationID  uint
	Wallet         Wallet `gorm:"foreignKey:DestinationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount         int
	Description    string
	Category       string
	Type           string `gorm:"-"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (t *Transaction) AfterFind(tx *gorm.DB) (err error) {
	if strings.Contains(t.Category, "Receive") || strings.Contains(t.Category, "Top") {
		t.Type = "+"
	}
	if strings.Contains(t.Category, "Send") || strings.Contains(t.Category, "Withdraw") {
		t.Type = "-"
	}
	return nil
}
