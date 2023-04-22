package config

import (
	"os"
	"strconv"
	"strings"
)

func SysTZ() string {
	tz := os.Getenv("TZ")
	if tz == "" {
		return "Asia/Shanghai"
	}
	return tz
}

func Docker() bool {
	docker := os.Getenv("DOCKER_MODE")
	if docker == "" || docker == "false" {
		return false
	}
	return true
}

func DBHost() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		split := strings.Split(Get("postgres.host").(string), ":")
		return split[0]
	}
	return host
}

func DBPort() string {
	port := os.Getenv("DB_PORT")
	if port == "" {
		full := Get("postgres.host").(string)
		if !strings.Contains(full, ":") {
			return "5432"
		}
		split := strings.Split(full, ":")
		return split[1]
	}
	return port
}

func DBUser() string {
	user := os.Getenv("DB_USER")
	if user == "" {
		return Get("postgres.username").(string)
	}
	return user
}

func DBPass() string {
	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		return Get("postgres.password").(string)
	}
	return pass
}

func DBName() string {
	name := os.Getenv("DB_NAME")
	if name == "" {
		return Get("postgres.name").(string)
	}
	return name
}

func DBSSL() string {
	ssl := os.Getenv("DB_SSL")
	if ssl == "" {
		return "disable"
	}
	return ssl
}

func RedisHosts() []string {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		cfg := Get("redis.hosts").([]interface{})
		var hosts = make([]string, len(cfg))
		for key, value := range cfg {
			hosts[key] = value.(string)
		}
		return hosts
	}
	return []string{host}
}

func RedisMaster() string {
	master := os.Getenv("REDIS_MASTER")
	if master == "" {
		return Get("redis.master").(string)
	}
	return master
}

func RedisUser() string {
	user := os.Getenv("REDIS_USER")
	if user == "" {
		return Get("redis.username").(string)
	}
	return user
}

func RedisPass() string {
	pass := os.Getenv("REDIS_PASSWORD")
	if pass == "" {
		return Get("redis.password").(string)
	}
	return pass
}

func RedisDB() int {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if db == 0 || err != nil {
		return 0
	}
	return db
}
