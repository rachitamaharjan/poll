package migrations

import (
	"github.com/rachitamaharjan/poll/db"
	"github.com/rachitamaharjan/poll/models"
)

func SyncDB() {
	db.DB.AutoMigrate(&models.Poll{}, &models.Option{})
}
