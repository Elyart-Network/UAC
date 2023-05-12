package data

import (
	"github.com/Elyart-Network/UAC/config"
	"github.com/Elyart-Network/UAC/data/actions"
	"github.com/Elyart-Network/UAC/data/drivers"
	"strings"
)

func SplitHostPort(url string) (host, port string) {
	if !strings.Contains(url, ":") {
		return url, "5432"
	}
	split := strings.Split(url, ":")
	return split[0], split[1]
}

func InterfaceSliceToStringSlice(i []interface{}) (s []string) {
	for _, v := range i {
		s = append(s, v.(string))
	}
	return s
}

func Init() {
	pgssl := "disable"
	if config.Get("postgres.ssl").(bool) {
		pgssl = "enable"
	}
	pghost, pgport := SplitHostPort(config.Get("postgres.host").(string))
	pgdsn := drivers.PostgresDSN{
		Host:     pghost,
		Port:     pgport,
		User:     config.Get("postgres.username").(string),
		Password: config.Get("postgres.password").(string),
		Name:     config.Get("postgres.name").(string),
		SSL:      pgssl,
		TimeZone: config.SysTZ(),
	}
	rdsn := drivers.RedisDSN{
		Hosts:    InterfaceSliceToStringSlice(config.Get("redis.hosts").([]interface{})),
		Master:   config.Get("redis.master").(string),
		Username: config.Get("redis.username").(string),
		Password: config.Get("redis.password").(string),
		DB:       0,
	}
	actions.New(drivers.Postgres(pgdsn), drivers.Redis(rdsn))
}
