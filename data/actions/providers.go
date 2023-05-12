package actions

import (
	"github.com/Elyart-Network/UAC/model"
)

func AddProvider(provider model.Providers) (int64, error) {
	db := handler.Database
	res := db.Create(&provider)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetProvider(uuid string) (model.Providers, error) {
	db := handler.Database
	var provider model.Providers
	res := db.Where("uuid = ?", uuid).First(&provider)
	if res.Error != nil {
		return provider, res.Error
	}
	return provider, nil
}

func UpdateProvider(provider model.Providers) (int64, error) {
	db := handler.Database
	res := db.Save(&provider)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DeleteProvider(id int) (int64, error) {
	db := handler.Database
	res := db.Where("id = ?", id).Delete(&model.Providers{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
