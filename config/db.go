package config

import (
	"fmt"
	"os"
	"test-stone/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// ConnectDB returns initialized gorm.DB
func ConnectDB() (*gorm.DB, error) {
	user := getEnvWithDefault("DB_USER", "admin")
	name := getEnvWithDefault("DB_NAME", "bank")
	password := getEnvWithDefault("DB_PASSWORD", "123456")
	host := getEnvWithDefault("DB_HOST", "localhost")
	port := getEnvWithDefault("DB_PORT", "5432")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, name, password)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	RunMigrations(db)

	return db, nil
}

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&domain.Account{}, &domain.Transfer{})
}

func getEnvWithDefault(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}
