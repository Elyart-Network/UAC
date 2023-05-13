package actions

import (
	"context"
	"encoding/json"
	"time"
)

type Session struct {
	ctx context.Context
}

func NewSession(ctx context.Context) *Session {
	return &Session{ctx}
}

func (s *Session) New(flag string, data any) error {
	rdb := handler.Redis
	cacheKey := "session:" + flag
	cacheData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = rdb.Do(s.ctx, "JSON.SET", cacheKey, "$", cacheData).Err()
	if err != nil {
		return err
	}
	rdb.Expire(s.ctx, cacheKey, 7*24*time.Hour)
	return nil
}

func (s *Session) Get(flag string) (string, error) {
	rdb := handler.Redis
	cacheKey := "session:" + flag
	cacheRes, err := rdb.Do(s.ctx, "JSON.GET", cacheKey, "$").Result()
	if err != nil {
		return "", err
	}
	return cacheRes.(string), nil
}

func (s *Session) Update(flag string, data any) error {
	rdb := handler.Redis
	cacheKey := "session:" + flag
	cacheData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = rdb.Do(s.ctx, "JSON.SET", cacheKey, "$", cacheData).Err()
	if err != nil {
		return err
	}
	rdb.Expire(s.ctx, cacheKey, 7*24*time.Hour)
	return nil
}

func (s *Session) Delete(flag string) error {
	rdb := handler.Redis
	cacheKey := "session:" + flag
	err := rdb.Del(s.ctx, cacheKey).Err()
	return err
}
