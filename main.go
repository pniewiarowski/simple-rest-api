package main

import (
	"github.com/joho/godotenv"
	"github.com/pniewiarowski/simple-rest-api/app"
	"github.com/pniewiarowski/simple-rest-api/database"
	"github.com/pniewiarowski/simple-rest-api/models"
	"log"
	"os"
	"strconv"
)

var Models = []interface{}{
	models.Book{},
	models.Author{},
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("create .env file first")
		return
	}

	db := os.Getenv("DATABASE")

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		log.Fatal("port in .env file should be a number")
		return
	}

	migration, err := strconv.ParseBool(os.Getenv("MIGRATION"))
	if err != nil {
		log.Fatal("migration in .env file should be a boolean")
		return
	}

	database.Setup(db)
	if migration {
		database.MakeMigration(Models)
	}

	app.Run(int(port))
}
