package usecase

import (
	"errors"
	"log"

	"github.com/higordenomar/codepix/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository         model.PixKeyRepositoryInterface
}

func (t *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo, pixKeyKindTo, description string) (*model.Transaction, error) {
	account, err := t.PixRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := t.PixRepository.FindKeyById(pixKeyTo, pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)

	if err != nil {
		return nil, err
	}

	t.TransactionRepository.Save(transaction)

	if transaction.ID == "" {
		return nil, errors.New("unable to process this transaction")
	}

	return transaction, nil
}

func (t *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)

	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed

	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionCompleted

	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Error(transactionId, description string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = description

	err = t.TransactionRepository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
