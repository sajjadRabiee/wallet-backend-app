package service

import (
	"wallet/internal/dto"
	"wallet/internal/model"
	"wallet/internal/repository"
	"wallet/pkg/custom_error"
)

type WalletService interface {
	GetWalletByUserId(input *dto.WalletRequestBody) (*model.Wallet, error)
	CreateWallet(user *model.User) (*model.Wallet, error)
}

type walletService struct {
	userRepository   repository.UserRepository
	walletRepository repository.WalletRepository
}

type WSConfig struct {
	UserRepository   repository.UserRepository
	WalletRepository repository.WalletRepository
}

func NewWalletService(c *WSConfig) WalletService {
	return &walletService{
		userRepository:   c.UserRepository,
		walletRepository: c.WalletRepository,
	}
}

func (s *walletService) GetWalletByUserId(input *dto.WalletRequestBody) (*model.Wallet, error) {
	wallet, err := s.walletRepository.FindByUserId(input.UserID)
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (s *walletService) CreateWallet(user *model.User) (*model.Wallet, error) {
	wallet, err := s.walletRepository.FindByUserId(int(user.ID))
	if err != nil {
		return &model.Wallet{}, err
	}
	if wallet.ID != 0 {
		return &model.Wallet{}, &custom_error.WalletAlreadyExistsError{}
	}

	wallet = &model.Wallet{
		UserID:  user.ID,
		Number:  user.PhoneNumber,
		Balance: 0,
	}

	newWallet, err := s.walletRepository.Save(wallet)
	if err != nil {
		return newWallet, err
	}

	return newWallet, nil
}
