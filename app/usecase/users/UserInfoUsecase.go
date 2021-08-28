package usecase

import (
	domain "example-archi/app/domain/users"
	lib "example-archi/library"
)

type userInfoUsecase struct {
	RepoUser domain.UserRepository
}

func NewUserUsecase(repoUser domain.UserRepository) domain.UserUsecase {
	return &userInfoUsecase{RepoUser: repoUser}
}

func (uc *userInfoUsecase) Create(data *domain.UserModel) error {
	// TODO: validate data
	// TODO: check existing data

	err := uc.RepoUser.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userInfoUsecase) GetAll(limitOffset lib.LimitOffset) ([]domain.UserModel, error) {
	result, err := uc.RepoUser.FetchAll(limitOffset)
	if err != nil {
		return nil, err
	}
	return result, nil
}
