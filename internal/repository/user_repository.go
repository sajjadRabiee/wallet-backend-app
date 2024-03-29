package repository

import (
	"gorm.io/gorm"
	"wallet/internal/model"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id int) (*model.User, error)
	FindByName(name string) ([]*model.User, error)
	FindByPhoneNumber(phoneNumber string) (*model.User, error)
	Save(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	SaveCard(card *model.Card) (*model.Card, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		db: c.DB,
	}
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) FindById(id int) (*model.User, error) {
	var user *model.User

	err := r.db.Where("id =?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByName(name string) ([]*model.User, error) {
	var users []*model.User

	err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) FindByPhoneNumber(phoneNumber string) (*model.User, error) {
	var user *model.User

	err := r.db.Where("phone_number = ?", phoneNumber).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Save(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) SaveCard(card *model.Card) (*model.Card, error) {
	err := r.db.Create(&card).Error
	if err != nil {
		return card, err
	}

	return card, nil
}

func (r *userRepository) Update(user *model.User) (*model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
