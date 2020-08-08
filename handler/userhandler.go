package handler

import (
	"fmt"
	"my-mod/models"
	rep "my-mod/repository"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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

var mySigningKey = []byte("mysecretkey")

func Login(c echo.Context) error {

	username := c.Get("username")
	admin := c.Get("admin")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "login thành công",
		"token":  tokenString,
	})
}
