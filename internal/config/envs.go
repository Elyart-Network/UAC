package config

import (
	"os"
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
		split := strings.Split(Get("database.host").(string), ":")
		return split[0]
	}
	return host
}

func DBPort() string {
	port := os.Getenv("DB_PORT")
	if port == "" {
		full := Get("database.host").(string)
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
		return Get("database.username").(string)
	}
	return user
}

func DBPass() string {
	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		return Get("database.password").(string)
	}
	return pass
}

func DBName() string {
	name := os.Getenv("DB_NAME")
	if name == "" {
		return Get("database.name").(string)
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
