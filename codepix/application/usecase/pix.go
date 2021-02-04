package usecase

import (
	"fmt"

	"github.com/gustavo-hillesheim/imersao-full-cycle/codepix-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepository
}

func (usecase *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := usecase.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	usecase.PixKeyRepository.Create(pixKey)
	if pixKey.ID == "" {
		return nil, fmt.Errorf("Unable to create new key at the moment")
	}
	return pixKey, nil
}

func (usecase *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := usecase.PixKeyRepository.FindByKeyAndKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, err
}
