package main

import (
	"log"

	"simpleblog/initializers"
	"simpleblog/loggerService"
	"simpleblog/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
		loggerService.Log.Printf("🚀 Could not load environment variables :%v", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.Blog{})
	log.Print("👍 Migration complete")
	loggerService.Log.Print("Migration Complete")
}
