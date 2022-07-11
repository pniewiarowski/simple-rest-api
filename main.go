package main

import (
	"flag"
	"fmt"
	"github.com/pniewiarowski/simple-rest-api/app"
	"github.com/pniewiarowski/simple-rest-api/database"
	"github.com/pniewiarowski/simple-rest-api/models"
)

var Models = []interface{}{
	models.Book{},
}

func main() {
	var db string
	var port int
	var migration bool

	flag.StringVar(&db, "database", "database", "Name for application database.")
	flag.IntVar(&port, "port", 7070, "Port for application server.")
	flag.BoolVar(&migration, "migration", true, "Should application make model migration.")
	flag.Parse()

	fmt.Printf("Current run application settings...\n")
	fmt.Printf("Database\t->\t%s.db\n", db)
	fmt.Printf("Port    \t->\t%d\n", port)
	fmt.Printf("Migrate \t->\t%t\n", migration)

	database.Setup(db)
	if migration {
		database.MakeMigration(Models)
	}

	app.Run(port)
}
