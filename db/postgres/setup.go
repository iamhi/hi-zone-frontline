package postgres

import (
	"fmt"

	"github.com/iamhi/frontline/config"
	"github.com/iamhi/frontline/db/postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

var databaseInitialized = false

func Setup() {
	if !config.IsInitialized() {
		fmt.Printf("Configuration is not loaded, cant start postgres connection")

		return
	}

	config := config.GetPostgresConfig()

	connection_string := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode)

	db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})

	if err != nil {
		fmt.Printf("Unable to establish connection: %s", err)

		return
	}

	if err := db.AutoMigrate(&models.User{}, &models.Message{}, &models.Box{}); err != nil {
		fmt.Printf("Unable to migrate user: %s", err)

		return
	}

	fmt.Println("Migrated database")

	Db = db

	databaseInitialized = true
}
