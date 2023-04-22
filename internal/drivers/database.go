package drivers

import (
	"github.com/Elyart-Network/UAC/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBDSN struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSL      string
	TimeZone string
}

func Database(dsn DBDSN) *gorm.DB {
	c := "host=" + dsn.Host + " port=" + dsn.Port + " user=" + dsn.User + " dbname=" + dsn.Name + " password=" + dsn.Password + " sslmode=" + dsn.SSL + " TimeZone=" + dsn.TimeZone
	db, err := gorm.Open(postgres.Open(c), &gorm.Config{})
	if err != nil {
		log.Fatalln("Can't connect to handlers!\n", err)
	}
	err = db.AutoMigrate(&models.Clients{}, models.Providers{}, &models.Credentials{}, &models.Users{}, &models.Tokens{})
	if err != nil {
		log.Fatalln("Migrate handlers failed!\n", err)
	}
	return db
}
