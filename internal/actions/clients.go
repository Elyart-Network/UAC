package actions

import "github.com/Elyart-Network/UAC/internal/models"

func AddClient(client models.Clients) (int64, error) {
	db := handler.Database
	res := db.Create(&client)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetClient(id int) (models.Clients, error) {
	db := handler.Database
	var client models.Clients
	res := db.Where("id = ?", id).First(&client)
	if res.Error != nil {
		return client, res.Error
	}
	return client, nil
}

func UpdateClient(client models.Clients) (int64, error) {
	db := handler.Database
	res := db.Save(&client)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DeleteClient(id int) (int64, error) {
	db := handler.Database
	res := db.Where("id = ?", id).Delete(&models.Clients{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
