package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	initDB()
	router := gin.Default()
	//router.GET("/tasks", getTasks)
	//router.POST("/tasks", createTask)
	//router.PUT("/tasks/:id", updateTask)
	//router.DELETE("/tasks/:id", deleteTask)
	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

func initDB() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to the database")
		fmt.Printf("db.Stats(): %v\n", db.Stats())
	}
}
