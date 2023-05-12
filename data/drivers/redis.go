package drivers

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisDSN struct {
	Hosts    []string
	Master   string
	Username string
	Password string
	DB       int
}

type RedisClient struct {
	redis.UniversalClient
}

func Redis(dsn RedisDSN) *RedisClient {
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      dsn.Hosts,
		MasterName: dsn.Master,
		Username:   dsn.Username,
		Password:   dsn.Password,
		DB:         dsn.DB,
	})
	ctx := context.Background()
	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("[Redis] Redis ping error: %s", err)
	}
	logrus.Debugf("[Redis] Redis ping: %s", ping)
	return &RedisClient{rdb}
}
