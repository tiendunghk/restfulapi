package main

import (
	"fmt"
	"my-mod/database"
	"my-mod/handler"
	mymdw "my-mod/mymiddleware"
	rep "my-mod/repository"
	imp "my-mod/repository/repo_implement"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var userImp rep.UserRepo

var mySigningKey = []byte("mysecretkey")

func main() {
	fmt.Println("Xin chao cac ban hahahah")

	isLogined := middleware.JWT(mySigningKey)
	isAdmin := mymdw.IsAdmin
	result := database.ConnectdbClient()
	userImp = imp.NewUserRepo(result)
	handler.UserImp = userImp
	mymdw.UserImp = userImp
	e := echo.New()
	e.GET("/users", handler.GetAll, isLogined)
	e.POST("/add", handler.AddUser, isLogined, isAdmin)
	e.POST("/login", handler.Login, middleware.BasicAuth(mymdw.BasicAuth))

	//a := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":8080"))

}
