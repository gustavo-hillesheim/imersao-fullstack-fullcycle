package repository

import (
	"fmt"

	"github.com/gustavo-hillesheim/imersao-full-cycle/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	Save(transaction *model.Transaction) error
	Find(id string) (*model.Transaction, error)
}

func (r *TransactionRepositoryDb) Create(transaction *model.Transaction) error {
	err := r.Db.Create(transaction).Error
	return err
}

func (r *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := r.Db.Save(transaction).Error
	return err
}

func (r *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction *model.Transaction
	r.Db.Preload("AccountFrom.Bank", &transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("No transaction was found")
	}
	return transaction, nil
}
