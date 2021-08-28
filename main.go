package main

import (
	domain "example-archi/app/domain/users"
	httplib "example-archi/library/http"
	"example-archi/route"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("Starting Program")
	e := echo.New()

	e.Validator = &httplib.CustomValidator{Validator: validator.New()}

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	log.Println("Connection to database....")
	dsn := "host=localhost user=example-user password=example-pass dbname=example-dbname port=5421 TimeZone=Asia/Jakarta"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect with DB")
	}

	table := []interface{}{
		&domain.UserModel{},
	}

	var dropTables bool
	envDropTables := os.Getenv("RECREATE_TABLES")
	dropTables, _ = strconv.ParseBool(envDropTables)

	if dropTables {
		log.Println("ENV Setting DROP_TABLES has found, droping tables")
		conn.Migrator().DropTable(table...)
	}

	err = conn.AutoMigrate(table...)
	if err != nil {
		panic("Failed to migrate db")
	}

	route.RegisterRoute(e, conn)

	e.Logger.Fatal(e.Start(":8182"))
}
