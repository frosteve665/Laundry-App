package main

import (
	"laundry/config"
	"laundry/view"

	_ "github.com/lib/pq"
)

func main() {
	db := config.ConnectDb()
	defer db.Close()

	view.MainMenu(db)
}
