package actions

import (
	"github.com/Elyart-Network/UAC/model"
)

func AddCredential(credential model.Credentials) (int64, error) {
	db := handler.Database
	res := db.Create(&credential)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetCredentialUUID(cid string) (string, error) {
	db := handler.Database
	var credential model.Credentials
	res := db.Where("client_id = ?", cid).First(&credential)
	if res.Error != nil {
		return credential.UUID, res.Error
	}
	return credential.UUID, nil
}

func GetCredential(uuid string) (model.Credentials, error) {
	db := handler.Database
	var credential model.Credentials
	res := db.Where("uuid = ?", uuid).First(&credential)
	if res.Error != nil {
		return credential, res.Error
	}
	return credential, nil
}

func UpdateCredential(credential model.Credentials) (int64, error) {
	db := handler.Database
	res := db.Save(&credential)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DeleteCredential(uuid string) (int64, error) {
	db := handler.Database
	res := db.Where("uuid = ?", uuid).Delete(&model.Credentials{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
