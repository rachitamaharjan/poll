package main

import (
	"github.com/rachitamaharjan/leave-management-system/db"
	"github.com/rachitamaharjan/leave-management-system/env"
	"github.com/rachitamaharjan/leave-management-system/migrations"
	"github.com/rachitamaharjan/leave-management-system/routes"
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
