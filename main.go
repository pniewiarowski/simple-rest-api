package main

import (
	"github.com/pniewiarowski/simple-rest-api/app"
	"github.com/pniewiarowski/simple-rest-api/database"
	"github.com/pniewiarowski/simple-rest-api/env"
	"github.com/pniewiarowski/simple-rest-api/models"
)

var Models = []interface{}{
	models.Book{},
	models.Author{},
	models.Auth{},
}

func main() {
	env.Load(".env")

	db := env.GetDb()
	migration := env.GetMigration()
	port := env.GetPort()

	database.Setup(db)
	if migration {
		database.MakeMigration(Models)
	}

	app.Run(int(port))
}
