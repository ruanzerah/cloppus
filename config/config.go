package config

import (
	"fmt"
	"os"

	"github.com/ruanzerah/cloppus/config/api"
	"github.com/ruanzerah/cloppus/config/database"
)

var (
	addr       = os.Getenv("API_ADDR")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")

	dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
)

func InitConfigs() error {
	err := database.InitDatabase(dsn)
	if err != nil {
		return err
	}
	err = api.InitAPI(":" + addr)
	if err != nil {
		return err
	}

	return nil
}
