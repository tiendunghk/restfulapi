package main

import (
	"fmt"
	"my-mod/database"
	"my-mod/handler"
	"my-mod/models"
	rep "my-mod/repository"
	imp "my-mod/repository/repo_implement"
	"net/http"

	"github.com/labstack/echo/v4"
)

var userImp rep.UserRepo

func getAll(c echo.Context) error {
	users, _ := userImp.Select()
	return c.JSON(http.StatusOK, users)
}
func addUser(c echo.Context) error {
	req := new(models.User)
	c.Bind(req)
	userImp.Insert(*req)
	return c.JSON(http.StatusOK, "Insert thanh cong")
}
func main() {
	fmt.Println("Xin chao cac ban hahahah")

	result := database.ConnectdbClient()
	userImp = imp.NewUserRepo(result)
	handler.UserImp = userImp
	e := echo.New()
	e.GET("/users", handler.GetAll)
	e.POST("/add", handler.AddUser)

	e.Logger.Fatal(e.Start(":8080"))
}
