package models

import (
	"context"

	"firebase.google.com/go/v4/db"
)

type MyDatabase struct {
	Client *db.Client
	Ctx    context.Context
}

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
