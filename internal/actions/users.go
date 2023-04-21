package actions

import "github.com/Elyart-Network/UAC/internal/models"

func AddUser(user models.Users) (int64, error) {
	db := handler.Database
	res := db.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func GetUserByID(id int) (models.Users, error) {
	db := handler.Database
	var user models.Users
	res := db.Where("id = ?", id).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func GetUserByName(username string) (models.Users, error) {
	db := handler.Database
	var user models.Users
	res := db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func GetUserByEmail(email string) (models.Users, error) {
	db := handler.Database
	var user models.Users
	res := db.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func UpdateUser(user models.Users) (int64, error) {
	db := handler.Database
	res := db.Save(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DeleteUser(id int) (int64, error) {
	db := handler.Database
	res := db.Where("id = ?", id).Delete(&models.Users{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
