package db

import (
	"Subscription/pkg/Server"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	username := viper.GetString("db.username")
	password := os.Getenv("DB_PASSWORD")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	dbName := viper.GetString("db.name")
	sslMode := viper.GetString("db.sslmode")

	// Формируем строку подключения (DSN)
	dsn := "host=" + host + " user=" + username + " password=" + password +
		" dbname=" + dbName + " port=" + port + " sslmode=" + sslMode
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database : %v", err)
	}
	if err := db.AutoMigrate(&Server.Subscription{}); err != nil {
		log.Fatalf("Could not migrate : %v", err)
	}

	return db, nil
}

// createTableQuery := `
// CREATE TABLE IF NOT EXISTS subscriptions (
//     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//     service_name VARCHAR(255) NOT NULL,
//     price NUMERIC(10, 2) NOT NULL,
//     user_id UUID NOT NULL,
//     start_date VARCHAR(7) NOT NULL,
//     end_date VARCHAR(7) NULL
// );`

// result := db.Exec(createTableQuery)
// if result.Error != nil {
// 	log.Fatalf("Could not create table: %v", result.Error)
// }
