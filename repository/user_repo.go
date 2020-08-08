package repository

import (
	models "my-mod/models"
)

type UserRepo interface {
	Select() ([]models.User, error)
	Insert(user models.User) error
}
