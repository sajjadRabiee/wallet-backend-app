package service

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"wallet/internal/dto"
	"wallet/internal/model"
	"wallet/internal/repository"
	"wallet/pkg/custom_error"
	"wallet/pkg/utils"
)

type AuthService interface {
	Attempt(input *dto.LoginRequestBody) (*model.User, error)
	ForgotPass(input *dto.ForgotPasswordRequestBody) (*model.PasswordReset, error)
	ResetPass(input *dto.ResetPasswordRequestBody) (*model.PasswordReset, error)
}

type authService struct {
	userRepository          repository.UserRepository
	passwordResetRepository repository.PasswordResetRepository
}

type ASConfig struct {
	UserRepository          repository.UserRepository
	PasswordResetRepository repository.PasswordResetRepository
}

func NewAuthService(c *ASConfig) AuthService {
	return &authService{
		userRepository:          c.UserRepository,
		passwordResetRepository: c.PasswordResetRepository,
	}
}

func (s *authService) Attempt(input *dto.LoginRequestBody) (*model.User, error) {
	user, err := s.userRepository.FindByPhoneNumber(input.PhoneNumber)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, &custom_error.UserNotFoundError{}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, &custom_error.IncorrectCredentialsError{}
	}

	return user, nil
}

func (s *authService) ForgotPass(input *dto.ForgotPasswordRequestBody) (*model.PasswordReset, error) {

	user, err := s.userRepository.FindByPhoneNumber(input.PhoneNumber)
	if err != nil {
		return &model.PasswordReset{}, err
	}

	if user.ID == 0 {
		return &model.PasswordReset{}, &custom_error.UserNotFoundError{}
	}

	passwordReset, err := s.passwordResetRepository.FindByUserId(int(user.ID))
	if err != nil {
		return &model.PasswordReset{}, err
	}

	passwordReset.UserID = user.ID
	passwordReset.Token = utils.GenerateString(10)
	passwordReset.ExpiredAt = time.Now().Add(time.Minute * 15)

	passwordReset, err = s.passwordResetRepository.Save(passwordReset)
	passwordReset.User = *user

	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}

func (s *authService) ResetPass(input *dto.ResetPasswordRequestBody) (*model.PasswordReset, error) {
	passwordReset, err := s.passwordResetRepository.FindByToken(input.Token)
	if err != nil {
		return passwordReset, err
	}

	if passwordReset.User.PhoneNumber == "" {
		return passwordReset, &custom_error.ResetTokenNotFound{}
	}

	if input.Password != input.ConfirmPassword {
		return passwordReset, &custom_error.PasswordNotSame{}
	}

	user := &passwordReset.User
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return passwordReset, err
	}
	user.Password = string(passwordHash)

	_, err = s.userRepository.Update(user)
	if err != nil {
		return passwordReset, err
	}

	passwordReset, err = s.passwordResetRepository.Delete(passwordReset)
	if err != nil {
		return passwordReset, err
	}

	return passwordReset, nil
}
