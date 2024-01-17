package service

import (
	"strconv"
	"wallet/internal/dto"
	"wallet/internal/model"
	"wallet/internal/repository"
	"wallet/pkg/custom_error"
)

type TransactionService interface {
	GetTransactions(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error)
	TopUp(input *dto.TopUpRequestBody) (*model.Transaction, error)
	Transfer(input *dto.TransferRequestBody) (*model.Transaction, error)
	CountTransaction(userID int) (int64, error)
}

type transactionService struct {
	transactionRepository  repository.TransactionRepository
	walletRepository       repository.WalletRepository
	sourceOfFundRepository repository.SourceOfFundRepository
}

type TSConfig struct {
	TransactionRepository  repository.TransactionRepository
	WalletRepository       repository.WalletRepository
	SourceOfFundRepository repository.SourceOfFundRepository
}

func NewTransactionService(c *TSConfig) TransactionService {
	return &transactionService{
		transactionRepository:  c.TransactionRepository,
		walletRepository:       c.WalletRepository,
		sourceOfFundRepository: c.SourceOfFundRepository,
	}
}

func (s *transactionService) GetTransactions(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error) {
	transactions, err := s.transactionRepository.FindAll(userID, query)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) TopUp(input *dto.TopUpRequestBody) (*model.Transaction, error) {
	sourceOfFund, err := s.sourceOfFundRepository.FindById(input.SourceOfFundID)
	if err != nil {
		return &model.Transaction{}, err
	}
	if sourceOfFund.ID == 0 {
		return &model.Transaction{}, &custom_error.SourceOfFundNotFoundError{}
	}

	wallet, err := s.walletRepository.FindByUserId(int(input.User.ID))
	if err != nil {
		return &model.Transaction{}, err
	}
	if wallet.ID == 0 {
		return &model.Transaction{}, &custom_error.WalletNotFoundError{}
	}

	transaction := &model.Transaction{
		SourceOfFundID: &sourceOfFund.ID,
		SourceOfFund:   sourceOfFund,
		UserID:         input.User.ID,
		User:           *input.User,
		Wallet:         *wallet,
		DestinationID:  wallet.ID,
		Amount:         input.Amount,
		Description:    "Top Up from " + sourceOfFund.Name,
		Category:       "Top Up",
	}

	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	wallet.Balance = wallet.Balance + input.Amount
	wallet, err = s.walletRepository.Update(wallet)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *transactionService) CountTransaction(userID int) (int64, error) {
	totalTransactions, err := s.transactionRepository.Count(userID)
	if err != nil {
		return totalTransactions, err
	}

	return totalTransactions, nil
}

func (s *transactionService) Transfer(input *dto.TransferRequestBody) (*model.Transaction, error) {
	myWallet, err := s.walletRepository.FindByUserId(int(input.User.ID))
	if err != nil {
		return &model.Transaction{}, err
	}
	if myWallet.ID == 0 {
		return &model.Transaction{}, &custom_error.WalletNotFoundError{}
	}
	if myWallet.Balance < input.Amount {
		return &model.Transaction{}, &custom_error.InsufficientBallanceError{}
	}
	number := strconv.Itoa(input.WalletNumber)
	if myWallet.Number == number {
		return &model.Transaction{}, &custom_error.TransferToSameWalletError{}
	}

	destinationWallet, err := s.walletRepository.FindByNumber(number)
	if err != nil {
		return &model.Transaction{}, err
	}
	if destinationWallet.ID == 0 {
		return &model.Transaction{}, &custom_error.WalletNotFoundError{}
	}

	//create transaction for receiver
	transaction := &model.Transaction{
		UserID:        destinationWallet.User.ID,
		DestinationID: myWallet.ID,
		Amount:        input.Amount,
		Description:   input.Description,
		Category:      "Receive Money",
	}

	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	// create transaction for sender
	transaction = &model.Transaction{
		UserID:        input.User.ID,
		DestinationID: destinationWallet.ID,
		Amount:        input.Amount,
		Description:   input.Description,
		Category:      "Send Money",
	}

	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	myWallet.Balance = myWallet.Balance - input.Amount
	myWallet, err = s.walletRepository.Update(myWallet)
	if err != nil {
		return transaction, err
	}

	destinationWallet.Balance = destinationWallet.Balance + input.Amount
	_, err = s.walletRepository.Update(destinationWallet)
	if err != nil {
		return transaction, err
	}

	balance := uint(myWallet.Balance)
	transaction.SourceOfFundID = &balance
	transaction.User = *input.User
	transaction.Wallet = *destinationWallet

	return transaction, nil
}
