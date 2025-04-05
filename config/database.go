package config

import (
	"fmt"
	"log"
	"time"

	"api-contact-form/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB is a global variable that holds the database connection instance.
// It is accessible throughout the application for executing database operations.
var DB *gorm.DB

// InitDB initializes the database connection using environment variables.
// It sets up the connection pool and performs automatic migrations for the Contact models.
//
// The function performs the following steps:
// 1. Retrieves database configuration from environment variables.
// 2. Constructs the Data Source Name (DSN) for MySQL connection.
// 3. Opens the database connection using GORM with a singular table naming strategy.
// 4. Configures the connection pool with specified limits.
// 5. Automatically migrates the Contact models to create or update the corresponding table.
//
// If any step fails, the function will panic with an appropriate error message.
func InitDB() {
	// Retrieve database configuration from environment variables with default values.
	dbUser := GetEnv("DB_USER", "")
	dbPassword := GetEnv("DB_PASSWORD", "")
	dbHost := GetEnv("DB_HOST", "db")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "contactsdb")
	log.Println(dbUser, dbPassword, dbHost, dbPort, dbName)

	// Construct the Data Source Name (DSN) for MySQL connection.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error

	// Open the database connection using GORM with a singular table naming strategy.
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(fmt.Sprint("Failed to connect to database : %v", err))
	}
	// Configure the connection pool settings.
	sqlDB, err := DB.DB()

	if err != nil {
		panic("Failed to get database instance!")
	}

	sqlDB.SetMaxOpenConns(10)           // Maximum number of open connections to the database.
	sqlDB.SetMaxIdleConns(5)            // Maximum number of idle connections in the pool.
	sqlDB.SetConnMaxLifetime(time.Hour) // Maximum amount of time a connection may be reused.

	if err := DB.AutoMigrate(&models.Contact{}); err != nil {
		panic(fmt.Sprintf("Automigrate failed : %v", err))
	}
}
