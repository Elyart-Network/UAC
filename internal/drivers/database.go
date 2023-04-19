package drivers

import (
	"github.com/Elyart-Network/UAC/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Database(host string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(host), &gorm.Config{})
	if err != nil {
		log.Fatalln("Can't connect to handlers!\n", err)
	}
	err = db.AutoMigrate(&models.Clients{}, &models.Credentials{}, &models.Users{}, &models.Tokens{})
	if err != nil {
		log.Fatalln("Migrate handlers failed!\n", err)
	}
	return db
}
