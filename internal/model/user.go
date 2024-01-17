package model

type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	PhoneNumber  string
	Password     string
	NationalCode string `gorm:"column:national_code"`
}

func (User) TableName() string {
	return "users"
}
