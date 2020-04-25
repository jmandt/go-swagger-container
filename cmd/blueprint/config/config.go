package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	// the shared DB ORM object
	DB *gorm.DB

	ServerPort int `mapstructure:"server_port"`
	// the data source name (DSN) for connecting to the database. required.
}


func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("config")  // <- name of config file
	v.SetConfigType("yaml")
	v.SetEnvPrefix("blueprint")
	v.AutomaticEnv()
	port := getenv("PORT", "8080")
	v.SetDefault("server_port", port)
	fmt.Println("port set to ",)
	for _, path := range configPaths {
		v.AddConfigPath(path)  // <- // path to look for the config file in
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}