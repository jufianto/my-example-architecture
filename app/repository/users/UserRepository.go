package repository

import (
	domain "example-archi/app/domain/users"
	lib "example-archi/library"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &userRepository{Conn: Conn}
}

func (r *userRepository) Create(data *domain.UserModel) (err error) {
	result := r.Conn.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) FetchAll(limitOffset lib.LimitOffset) (result []domain.UserModel, err error) {
	ok := r.Conn.Select("username", "permission").Find(&result).Limit(limitOffset.Limit).Offset(limitOffset.Offset)
	if ok.Error != nil {
		return nil, ok.Error
	}
	log.Println("Total Rows", ok.RowsAffected)
	return result, nil
}

func (r *userRepository) CountAll() (totalRows uint64) {
	return
}

func (r *userRepository) Fetch(userName string) (result domain.UserModel, err error) {
	ok := r.Conn.First(&result, userName)
	if ok.Error != nil {
		return domain.UserModel{}, ok.Error
	}
	return result, nil
}
