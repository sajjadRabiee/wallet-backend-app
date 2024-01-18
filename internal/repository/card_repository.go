package repository

import (
	"gorm.io/gorm"
	"wallet/internal/model"
)

type CardRepository interface {
	FindByUserId(id int) ([]model.Card, error)
}

type cardRepository struct {
	db *gorm.DB
}

type CDConfig struct {
	DB *gorm.DB
}

func NewCardRepository(c *CDConfig) CardRepository {
	return &cardRepository{
		db: c.DB,
	}
}

func (c *cardRepository) FindByUserId(id int) ([]model.Card, error) {
	var card []model.Card
	if err := c.db.Where("user_id = ?", id).Find(&card).Error; err != nil {
		return nil, err
	}
	return card, nil
}
