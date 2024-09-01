package migrations

import (
	"github.com/rachitamaharjan/leave-management-system/db"
	"github.com/rachitamaharjan/leave-management-system/models"
)

func SyncDB() {
	db.DB.AutoMigrate(&models.User{}, &models.Poll{}, &models.Option{})
}
