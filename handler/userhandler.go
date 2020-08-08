package handler

import (
	"my-mod/models"
	rep "my-mod/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

var UserImp rep.UserRepo

func GetAll(c echo.Context) error {
	users, _ := UserImp.Select()
	return c.JSON(http.StatusOK, users)
}
func AddUser(c echo.Context) error {
	req := new(models.User)
	c.Bind(req)
	UserImp.Insert(*req)
	return c.JSON(http.StatusOK, echo.Map{
		"status": "Insert thanh cong",
	})
}
