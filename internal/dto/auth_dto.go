package dto

import "wallet/internal/model"

type LoginRequestBody struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required,min=5"`
}

type RegisterRequestBody struct {
	Name        string `json:"name" binding:"required,alphanum"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required,min=5"`
}

type ForgotPasswordRequestBody struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type ResetPasswordRequestBody struct {
	Token           string `json:"token" binding:"required"`
	Password        string `json:"password" binding:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=5"`
}

type ForgotPasswordResponseBody struct {
	PhoneNumber string `json:"phone_number"`
	Token       string `json:"token"`
}

type LoginResponseBody struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	WalletNumber string `json:"wallet"`
	Token        string `json:"token"`
}

func FormatLogin(user *model.User, wallet *model.Wallet, token string) LoginResponseBody {
	return LoginResponseBody{
		ID:           user.ID,
		Name:         user.Name,
		PhoneNumber:  user.PhoneNumber,
		WalletNumber: wallet.Number,
		Token:        token,
	}
}

func FormatForgotPassword(passwordReset *model.PasswordReset) ForgotPasswordResponseBody {
	return ForgotPasswordResponseBody{
		PhoneNumber: passwordReset.User.PhoneNumber,
		Token:       passwordReset.Token,
	}
}
