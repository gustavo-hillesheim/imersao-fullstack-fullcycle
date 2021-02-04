package usecase

import (
	"errors"
	"log"

	"github.com/gustavo-hillesheim/imersao-full-cycle/codepix-go/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepository
	PixKeyRepository      model.PixKeyRepository
}

func (usecase *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {
	account, err := usecase.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := usecase.PixKeyRepository.FindByKeyAndKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)
	if err != nil {
		return nil, err
	}

	usecase.TransactionRepository.Create(transaction)
	if transaction.ID == "" {
		return nil, errors.New("Unable to process this transaction")
	}
	return transaction, nil
}

func (usecase *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := usecase.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err = usecase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (usecase *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := usecase.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted
	err = usecase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
