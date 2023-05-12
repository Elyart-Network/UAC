package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

func GetEnv(Key string) string {
	_ = godotenv.Load()
	return os.Getenv(Key)
}

func SetEnvConf(Type string, ConfKey string) {
	ConfSplit := strings.Split(ConfKey, ".")
	var UpConfSplit []string
	for _, v := range ConfSplit {
		UpConfSplit = append(UpConfSplit, strings.ToUpper(v))
	}
	var EnvKey = strings.Join(UpConfSplit, "_")
	env := GetEnv(EnvKey)
	if env != "" {
		switch Type {
		case "string":
			Set(ConfKey, env)
		case "int":
			conv, err := strconv.Atoi(env)
			if err != nil {
				logrus.Errorf("[Config] Error converting string to int: %s", err)
			}
			Set(ConfKey, conv)
		case "bool":
			switch env {
			case "true":
				Set(ConfKey, true)
			case "false":
				Set(ConfKey, false)
			}
		case "array":
			trim := strings.TrimSpace(env)
			trim = strings.TrimPrefix(trim, "[")
			trim = strings.TrimSuffix(trim, "]")
			var envArray []string
			for _, v := range strings.Split(trim, ",") {
				trimSub := strings.TrimPrefix(v, "\"")
				trimSub = strings.TrimSuffix(trimSub, "\"")
				envArray = append(envArray, trimSub)
			}
			Set(ConfKey, envArray)
		}
	}
}

func EnvInit() {
	var dict = map[string][]string{
		"bool":   {},
		"int":    {},
		"string": {},
		"array":  {},
	}
	for key, value := range dict {
		for _, sub := range value {
			SetEnvConf(key, sub)
		}
	}
}

func SysTZ() string {
	tz := os.Getenv("TZ")
	if tz == "" {
		return "Asia/Shanghai"
	}
	return tz
}
