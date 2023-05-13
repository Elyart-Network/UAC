package actions

import (
	"github.com/Elyart-Network/UAC/model"
)

type Token struct{}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) Insert(token model.Tokens) (int64, error) {
	db := handler.Database
	res := db.Create(&token)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func (t *Token) Get(uuid string) (model.Tokens, error) {
	db := handler.Database
	var dbs model.Tokens
	res := db.Where("uuid = ?", uuid).First(&dbs)
	if res.Error != nil {
		return dbs, res.Error
	}
	return dbs, nil
}

func (t *Token) Update(token model.Tokens) (int64, error) {
	db := handler.Database
	res := db.Save(&token)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func (t *Token) Delete(uuid string) (int64, error) {
	db := handler.Database
	res := db.Where("uuid = ?", uuid).Delete(&model.Tokens{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
