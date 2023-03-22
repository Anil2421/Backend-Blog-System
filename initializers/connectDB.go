package initializers

import (
	"fmt"
	"simpleblog/loggerService"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Makassar", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	loggerService.Log.Print(dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		loggerService.Log.Fatal("Failed to connect to the Database")
	}
	loggerService.Log.Print("🚀 Connected Successfully to the Database")
	// set up the table inside the db
	createTableIfNotExist()
}
func createTableIfNotExist() {
	dbString := "create table IF NOT EXISTS blogs ( id INT PRIMARY KEY, title VARCHAR(28), content VARCHAR(300), author VARCHAR(28));"
	DB.Exec(dbString)
	loggerService.Log.Print("blog table is created")
}
