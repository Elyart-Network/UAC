package actions

import (
	"github.com/Elyart-Network/UAC/internal/drivers"
	"gorm.io/gorm"
)

var handler Handler

type Handler struct {
	Database *gorm.DB
	Redis    *drivers.RedisClient
}

func New(db *gorm.DB, rdb *drivers.RedisClient) {
	handler = Handler{db, rdb}
}
