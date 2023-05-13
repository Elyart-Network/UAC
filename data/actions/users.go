package actions

import (
	"context"
	"encoding/json"
	"github.com/Elyart-Network/UAC/model"
	"time"
)

type User struct {
	ctx context.Context
}

func NewUser(ctx context.Context) *User {
	return &User{ctx}
}

func (u *User) Insert(user model.Users) (int64, error) {
	// Insert into Database
	db := handler.Database
	res := db.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}

	// Insert into Cache
	rdb := handler.Redis
	cacheKey := "user:" + user.UUID
	cacheData, err := json.Marshal(user)
	if err != nil {
		return 0, err
	}
	err = rdb.Do(u.ctx, "JSON.SET", cacheKey, "$", cacheData).Err()
	rdb.Expire(u.ctx, cacheKey, 10*time.Minute)
	return res.RowsAffected, err
}

func (u *User) Get(uuid string) (model.Users, error) {
	// Lookup user in cache
	rdb := handler.Redis
	cacheKey := "user:" + uuid
	cacheRes, _ := rdb.Do(u.ctx, "JSON.GET", cacheKey, "$").Result()
	if cacheRes != nil {
		var sliceRes []model.Users
		err := json.Unmarshal([]byte(cacheRes.(string)), &sliceRes)
		return sliceRes[0], err
	}

	// Select user from database
	db := handler.Database
	var user model.Users
	res := db.Where("uuid = ?", uuid).First(&user)
	if res.Error != nil {
		return user, res.Error
	}

	// Add user to cache and return
	cacheData, err := json.Marshal(user)
	if err != nil {
		return user, err
	}
	err = rdb.Do(u.ctx, "JSON.SET", cacheKey, "$", cacheData).Err()
	rdb.Expire(u.ctx, cacheKey, 10*time.Minute)
	return user, err
}

func (u *User) ID(username string) (string, error) {
	db := handler.Database
	var user model.Users
	res := db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		return user.UUID, res.Error
	}
	return user.UUID, nil
}

func (u *User) Update(user model.Users) (int64, error) {
	// Update database
	db := handler.Database
	res := db.Where("uuid = ?", user.UUID).Updates(&user)
	if res.Error != nil {
		return 0, res.Error
	}

	// Update cache
	rdb := handler.Redis
	cacheKey := "user:" + user.UUID
	cacheData, err := json.Marshal(user)
	if err != nil {
		return 0, err
	}
	rdb.Del(u.ctx, cacheKey)
	err = rdb.Do(u.ctx, "JSON.SET", cacheKey, "$", cacheData).Err()
	return res.RowsAffected, err
}

func (u *User) Delete(uuid string) (int64, error) {
	// Delete from database
	db := handler.Database
	res := db.Where("uuid = ?", uuid).Delete(&model.Users{})
	if res.Error != nil {
		return 0, res.Error
	}

	// Delete from cache
	rdb := handler.Redis
	cacheKey := "user:" + uuid
	rdb.Del(u.ctx, cacheKey)
	return res.RowsAffected, nil
}
