package main

import (
	"fmt"
	"my-mod/database"
	"my-mod/handler"
	rep "my-mod/repository"
	imp "my-mod/repository/repo_implement"
	"os"

	"github.com/labstack/echo/v4"
)

var userImp rep.UserRepo

func main() {
	fmt.Println("Xin chao cac ban hahahah")

	result := database.ConnectdbClient()
	userImp = imp.NewUserRepo(result)
	handler.UserImp = userImp
	e := echo.New()
	e.GET("/users", handler.GetAll)
	e.POST("/add", handler.AddUser)

	os.Setenv("POST", "8080")
	a := os.Getenv("POST")

	e.Logger.Fatal(e.Start(":" + a))

}
