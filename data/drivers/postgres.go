package drivers

import (
	"github.com/Elyart-Network/UAC/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDSN struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSL      string
	TimeZone string
}

func Postgres(dsn PostgresDSN) *gorm.DB {
	c := "host=" + dsn.Host + " port=" + dsn.Port + " user=" + dsn.User + " dbname=" + dsn.Name + " password=" + dsn.Password + " sslmode=" + dsn.SSL + " TimeZone=" + dsn.TimeZone
	db, err := gorm.Open(postgres.Open(c), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("[Postgres] Can't connect to handlers! %s", err)
	}
	err = db.AutoMigrate(&model.Clients{}, model.Providers{}, &model.Credentials{}, &model.Users{}, &model.Tokens{})
	if err != nil {
		logrus.Fatalf("[Postgres] Migrate handlers failed! %s", err)
	}
	return db
}
