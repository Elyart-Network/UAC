package actions

import "gorm.io/gorm"

var handler Handler

type Handler struct {
	Database *gorm.DB
}

func New(db *gorm.DB) {
	handler = Handler{db}
}
