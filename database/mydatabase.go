package database

import (
	"context"
	"log"
	models "my-mod/models"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func ConnectdbClient() *models.MyDatabase {
	opt := option.WithCredentialsFile("cr.json")

	ct := context.Background()
	config := &firebase.Config{
		DatabaseURL: "https://testnotibackground.firebaseio.com",
	}

	app, err := firebase.NewApp(ct, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	c, err := app.Database(ct)
	if err != nil {
		log.Fatal(err)
	}

	var out = new(models.MyDatabase)
	out.Client = c
	out.Ctx = ct
	return out
}
