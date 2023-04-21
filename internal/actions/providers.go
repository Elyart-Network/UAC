package actions

import "github.com/Elyart-Network/UAC/internal/models"

func AddProvider(provider models.Providers) (int64, error) {
	db := handler.Database
	res := db.Create(&provider)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetProvider(id int) (models.Providers, error) {
	db := handler.Database
	var provider models.Providers
	res := db.Where("id = ?", id).First(&provider)
	if res.Error != nil {
		return provider, res.Error
	}
	return provider, nil
}

func UpdateProvider(provider models.Providers) (int64, error) {
	db := handler.Database
	res := db.Save(&provider)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DeleteProvider(id int) (int64, error) {
	db := handler.Database
	res := db.Where("id = ?", id).Delete(&models.Providers{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
