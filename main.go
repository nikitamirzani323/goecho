package main

import (
	"goecho/db"
	"goecho/routes"
)

func main() {
	db.Init()
	app := routes.Init()
	app.Logger.Fatal(app.Start(":8089"))
}
