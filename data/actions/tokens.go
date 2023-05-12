package actions

import (
	"github.com/Elyart-Network/UAC/model"
)

func AddToken(token model.Tokens) (int64, error) {
	db := handler.Database
	res := db.Create(&token)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetToken(uuid string) (model.Tokens, error) {
	db := handler.Database
	var dbs model.Tokens
	res := db.Where("uuid = ?", uuid).First(&dbs)
	if res.Error != nil {
		return dbs, res.Error
	}
	return dbs, nil
}

func UpdateToken(token model.Tokens) (int64, error) {
	db := handler.Database
	res := db.Save(&token)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DeleteToken(uuid string) (int64, error) {
	db := handler.Database
	res := db.Where("uuid = ?", uuid).Delete(&model.Tokens{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
