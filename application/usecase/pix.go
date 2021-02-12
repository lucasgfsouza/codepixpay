package usecase

import (
	"github.com/codeedu/imersao/codepix-go/domain/model"
)

type PixUseCase struct {

	PixKeyRepository model.PixKeyRepositoryInterface

}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, erro) {

	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil{
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, errors.New(text: "Unable to create new key at moment")
	}

	return pixKey, nil

}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {

	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}

	return pixKey, nil

}