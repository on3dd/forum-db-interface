package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"

	"forum-db-interface/controller"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config.env file: %v", err)
	}

	db, err := initDatabase(config)
	if err != nil {
		log.Fatalf("Error initializing DB: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging DB: %v", err)
	}
	defer db.Close()

	c := controller.New(db)

	server := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler: c,
	}

	log.Printf("Server successfully started at port %v\n", server.Addr)
	log.Println(server.ListenAndServe())
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
