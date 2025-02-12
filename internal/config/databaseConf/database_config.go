package databaseconf

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
}

func NewDBConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func ConnectDB() (*gorm.DB, error) {
	dbConfig, err := NewDBConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConfig.DBHost,
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.DBName,
		dbConfig.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
		return nil, err
	}
	return db, nil
}
