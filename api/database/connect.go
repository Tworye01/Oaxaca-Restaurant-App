package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Returns a connection to a database, please ensure that environmental variables have been loaded before running this function.
func PgConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf(`
        host=%s 
        user=%s 
        password=%s 
        dbname=%s 
        port=%s 
        sslmode=disable 
        TimeZone=Europe/London`,
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Failed to connect!")
		return nil, err
	}

	log.Println("Connected!")

	return db, nil
}

// Returns an SQLite database connection, please ensure environmental variables have been loaded before running this function.
// Uses the SQLITE_DB environmental variable in order to get the file for creating the connection.
func SQLiteConn() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}

	return db, nil
}
