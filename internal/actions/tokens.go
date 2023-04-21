package actions

import "github.com/Elyart-Network/UAC/internal/models"

func AddToken(token models.Tokens) (int64, error) {
	db := handler.Database
	res := db.Create(&token)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetToken(token string) (models.Tokens, error) {
	db := handler.Database
	var dbs models.Tokens
	res := db.Where("token = ?", token).First(&dbs)
	if res.Error != nil {
		return dbs, res.Error
	}
	return dbs, nil
}

func UpdateToken(token models.Tokens) (int64, error) {
	db := handler.Database
	res := db.Save(&token)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DeleteToken(token string) (int64, error) {
	db := handler.Database
	res := db.Where("token = ?", token).Delete(&models.Tokens{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
