package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// Read global config
	viper.SetConfigFile("config.yaml")
	viper.SetDefault("Server", serverDef)
	viper.SetDefault("Postgres", postgresDef)
	viper.SetDefault("Redis", redisDef)
	viper.SetDefault("Encrypt", encDef)
	conf := &config{}
	if err := viper.ReadInConfig(); err != nil {
		logrus.Info("[Config] Config file not found, creating...")
		if err := viper.WriteConfig(); err != nil {
			logrus.Errorf("[Config] Can't write config: %s", err)
		}
	}
	if err := viper.Unmarshal(conf); err != nil {
		logrus.Fatalf("[Config] Unable to decode into struct: %s", err)
	}
}

func Get(key string) interface{} {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("[Config] Error reading config: %s", err)
	}
	return viper.Get(key)
}

func Set(key string, value interface{}) {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("[Config] Error reading config: %s", err)
	}
	viper.Set(key, value)
}
