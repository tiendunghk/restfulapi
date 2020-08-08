package mymiddleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		admin := claims["admin"].(bool)
		if admin {
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}
