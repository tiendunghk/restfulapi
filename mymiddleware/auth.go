package mymiddleware

import (
	rep "my-mod/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

var UserImp rep.UserRepo

func BasicAuth(username string, password string, c echo.Context) (bool, error) {
	users, _ := UserImp.Select()

	for _, user := range users {
		if strconv.Itoa(user.Id) == username && user.Password == password {
			c.Set("username", username)
			c.Set("admin", user.IsAdmin)
			return true, nil
		}
	}
	return false, nil
}
