package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func New() (db *sql.DB, err error) {
	config, err := loadConfig()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error loading config.env file: %v", err))
	}

	db, err = initDatabase(config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error initializing DB: %v", err))
	}

	if err = db.Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("Error pinging DB: %v", err))
	}

	return db, nil
}

// Config represents structure of the config.env
type Config struct {
	dbUser string
	dbPass string
	dbName string
	dbHost string
	dbPort string
}

func loadConfig() (config *Config, err error) {
	err = godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading config.env file")
	}

	config = &Config{
		dbUser: os.Getenv("db_user"),
		dbPass: os.Getenv("db_pass"),
		dbName: os.Getenv("db_name"),
		dbHost: os.Getenv("db_host"),
		dbPort: os.Getenv("db_port"),
	}
	return config, err
}

func initDatabase(c *Config) (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.dbHost, c.dbPort, c.dbUser, c.dbPass, c.dbName)

	db, err = sql.Open("postgres", psqlInfo)
	return db, err
}
