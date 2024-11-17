package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ruanzerah/cloppus/config"
)

func main() {
	err := config.InitConfigs()
	if err != nil {
		log.Panic(err)
	}
}
