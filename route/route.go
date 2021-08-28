package route

import (
	controller "example-archi/app/delivery/http"
	repository "example-archi/app/repository/users"
	usecase "example-archi/app/usecase/users"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, conn *gorm.DB) {
	gUsers := e.Group("/v1/user")
	routeUsers(gUsers, conn)
}

func routeUsers(e *echo.Group, conn *gorm.DB) {
	repo := repository.NewUserRepository(conn)
	uc := usecase.NewUserUsecase(repo)
	c := controller.UserInfoController{Uc: uc}

	// Route
	e.GET("/store", c.SignUp)
}
