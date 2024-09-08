package main

import (
	"github.com/rachitamaharjan/poll/db"
	"github.com/rachitamaharjan/poll/env"
	"github.com/rachitamaharjan/poll/migrations"
	"github.com/rachitamaharjan/poll/routes"
)

func init() {
	env.LoadEnvVariables()
	db.ConnectToDB()
	migrations.SyncDB()
}

func main() {
	r := routes.SetupRouter()
	r.Run()
}
