package controller

import (
	domain "example-archi/app/domain/users"
	lib1 "example-archi/library"
	lib "example-archi/library/encrypt"
	httplib "example-archi/library/http"

	"github.com/labstack/echo"
)

type UserInfoController struct {
	Uc domain.UserUsecase
}

func (h *UserInfoController) SignUp(c echo.Context) error {
	req := new(domain.UserRequestBody)
	if err := c.Bind(req); err != nil {
		return httplib.SendFailedBindEcho(c, err)
	}
	userModel := domain.UserModel{
		UserName:   req.UserName,
		Password:   req.Password,
		Permission: []string{"USER"},
	}

	if userModel.Password != "" {
		userModel.Password = lib.EncryptSHA512(userModel.Password)
	}

	if err := c.Validate(userModel); err != nil {
		return err
	}

	err := h.Uc.Create(&userModel)
	if err != nil {
		userModel.Password = "upss, di terencrypt dong"
		return httplib.SendFailedCreate(c, err, userModel)
	}
	responseData := httplib.ResponseData{
		Data: userModel,
	}
	return c.JSON(200, responseData)
}

func (h *UserInfoController) GetAll(c echo.Context) error {
	req := new(lib1.LimitOffset)
	if err := c.Bind(req); err != nil {
		return httplib.SendFailedBindEcho(c, err)
	}
	return nil
}
