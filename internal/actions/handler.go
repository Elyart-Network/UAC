package actions

import (
	"context"
	"github.com/Elyart-Network/UAC/internal/drivers"
	"gorm.io/gorm"
	"log"
)

var handler Handler

type Handler struct {
	Database *gorm.DB
	Redis    *drivers.RedisClient
}

func New(db *gorm.DB, rdb *drivers.RedisClient) {
	handler = Handler{db, rdb}
	Test(&handler)
}

func Test(handler *Handler) {
	ctx := context.Background()
	result, err := handler.Redis.Ping(ctx).Result()
	if err != nil {
		log.Panicln("Redis ping error:", err)
		return
	}
	log.Println("Redis ping:", result)
}
