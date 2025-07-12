package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	initDB()
	runMigrations()
	router := gin.Default()
	//router.GET("/tasks", getTasks)
	//router.POST("/tasks", createTask)
	//router.PUT("/tasks/:id", updateTask)
	//router.DELETE("/tasks/:id", deleteTask)
	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

func initDB() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to the database")
		//fmt.Printf("db.Stats(): %v\n", db.Stats())
	}
}
func runMigrations() {
	m, err := migrate.New(
		"file://db/migrations",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
