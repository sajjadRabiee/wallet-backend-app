package dto

import (
	"wallet/internal/model"
)

type CardsBody struct {
	CardNumber string `json:"card_number" binding:"required"`
}

type UserRequestParams struct {
	UserID int `uri:"id" binding:"required"`
}

type UserRequestQuery struct {
	Name        string `form:"name"`
	PhoneNumber string `form:"phone_number"`
}

type UserResponseBody struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func FormatUser(user *model.User) UserResponseBody {
	formattedUser := UserResponseBody{}
	formattedUser.ID = user.ID
	formattedUser.Name = user.Name
	formattedUser.PhoneNumber = user.PhoneNumber
	return formattedUser
}

func FormatUsers(authors []*model.User) []UserResponseBody {
	formattedUsers := []UserResponseBody{}
	for _, user := range authors {
		formattedUser := FormatUser(user)
		formattedUsers = append(formattedUsers, formattedUser)
	}
	return formattedUsers
}

type UserDetailResponse struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	PhoneNumber string         `json:"phone_number"`
	Wallet      WalletResponse `json:"wallet"`
}

func FormatUserDetail(user *model.User, wallet *model.Wallet) UserDetailResponse {
	formattedUser := UserDetailResponse{}
	formattedUser.ID = user.ID
	formattedUser.Name = user.Name
	formattedUser.PhoneNumber = user.PhoneNumber
	formattedUser.Wallet = FormatWallet(wallet)
	return formattedUser
}
