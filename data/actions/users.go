package actions

import (
	"github.com/Elyart-Network/UAC/model"
)

func AddUser(user model.Users) (int64, error) {
	db := handler.Database
	res := db.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetUserByID(uuid string) (model.Users, error) {
	db := handler.Database
	var user model.Users
	res := db.Where("uuid = ?", uuid).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func GetUserByName(username string) (model.Users, error) {
	db := handler.Database
	var user model.Users
	res := db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func UpdateUser(user model.Users) (int64, error) {
	db := handler.Database
	res := db.Save(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DeleteUser(uuid string) (int64, error) {
	db := handler.Database
	res := db.Where("uuid = ?", uuid).Delete(&model.Users{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
