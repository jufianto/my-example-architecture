package domain

import (
	lib "example-archi/library"

	"github.com/lib/pq"
)

type UserModel struct {
	UserName   string         `json:"username" gorm:"primaryKey;column:username" validate:"required"`
	Password   string         `json:"-" validate:"required"`
	Permission pq.StringArray `json:"permission" gorm:"type:text[]"`
}

type UserRequestBody struct {
	UserName   string   `json:"username"`
	Password   string   `json:"password"`
	Permission []string `json:"permission"`
}

type UserRepository interface {
	Create(data *UserModel) error
	FetchAll(limitOffset lib.LimitOffset) (result []UserModel, err error)
}

type UserUsecase interface {
	Create(data *UserModel) error
	//FetchAll(limitOffset lib.LimitOffset) (result []UserModel, err error)
}
