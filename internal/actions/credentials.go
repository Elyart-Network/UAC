package actions

import "github.com/Elyart-Network/UAC/internal/models"

func AddCredential(credential models.Credentials) (int64, error) {
	db := handler.Database
	res := db.Create(&credential)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetCredential(cid int) (models.Credentials, error) {
	db := handler.Database
	var credential models.Credentials
	res := db.Where("client_id = ?", cid).First(&credential)
	if res.Error != nil {
		return credential, res.Error
	}
	return credential, nil
}

func UpdateCredential(credential models.Credentials) (int64, error) {
	db := handler.Database
	res := db.Save(&credential)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DeleteCredential(id int) (int64, error) {
	db := handler.Database
	res := db.Where("id = ?", id).Delete(&models.Credentials{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
