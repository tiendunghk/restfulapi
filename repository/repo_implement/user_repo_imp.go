package repo_implement

import (
	"context"
	"fmt"
	"my-mod/models"
	rep "my-mod/repository"

	"firebase.google.com/go/v4/db"
	"github.com/google/uuid"
)

type UserRepoImp struct {
	Db *models.MyDatabase
}

func NewUserRepo(db *models.MyDatabase) rep.UserRepo {
	return &UserRepoImp{
		Db: db,
	}
}
func (r *UserRepoImp) Select() ([]models.User, error) {
	var (
		client *db.Client      = r.Db.Client
		ctx    context.Context = r.Db.Ctx
	)

	ref := client.NewRef("Users")
	var results map[string]models.User
	if err := ref.Get(ctx, &results); err != nil {
		fmt.Println(err.Error())
	}
	output := make([]models.User, 0)
	for _, v := range results {
		output = append(output, v)
	}
	return output, nil
}
func (r *UserRepoImp) Insert(user models.User) error {
	var (
		client *db.Client      = r.Db.Client
		ctx    context.Context = r.Db.Ctx
	)
	var path string = "Users"
	a, _ := uuid.NewRandom()
	path = path + "/" + a.String()
	ref := client.NewRef(path)
	ref.Set(ctx, user)
	return nil
}
