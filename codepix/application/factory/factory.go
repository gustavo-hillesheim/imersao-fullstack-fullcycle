package factory

import (
	"github.com/gustavo-hillesheim/imersao-full-cycle/codepix-go/application/usecase"
	"github.com/gustavo-hillesheim/imersao-full-cycle/codepix-go/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := &repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := &repository.TransactionRepositoryDb{Db: database}
	return usecase.TransactionUseCase{
		TransactionRepository: transactionRepository,
		PixKeyRepository:      pixRepository,
	}
}
