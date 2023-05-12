package actions

import (
	"github.com/Elyart-Network/UAC/model"
)

func AddClient(client model.Clients) (int64, error) {
	db := handler.Database
	res := db.Create(&client)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetClient(uuid string) (model.Clients, error) {
	db := handler.Database
	var client model.Clients
	res := db.Where("uuid = ?", uuid).First(&client)
	if res.Error != nil {
		return client, res.Error
	}
	return client, nil
}

func UpdateClient(client model.Clients) (int64, error) {
	db := handler.Database
	res := db.Save(&client)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DeleteClient(uuid string) (int64, error) {
	db := handler.Database
	res := db.Where("uuid = ?", uuid).Delete(&model.Clients{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
