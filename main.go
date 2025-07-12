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
	//runMigrations()
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTask)
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

func getTasks(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, description, completed, created_at, updated_at FROM tasks ORDER BY created_at DESC")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer rows.Close()

	var tasks []gin.H
	for rows.Next() {
		var id int
		var title, description string
		var completed bool
		var createdAt, updatedAt string

		err := rows.Scan(&id, &title, &description, &completed, &createdAt, &updatedAt)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to scan task"})
			return
		}

		tasks = append(tasks, gin.H{
			"id":          id,
			"title":       title,
			"description": description,
			"completed":   completed,
			"created_at":  createdAt,
			"updated_at":  updatedAt,
		})
	}

	if err = rows.Err(); err != nil {
		c.JSON(500, gin.H{"error": "Error iterating tasks"})
		return
	}

	c.JSON(200, gin.H{"tasks": tasks})
}

func createTask(c *gin.Context) {
	var task struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	var id int
	err := db.QueryRow(
		"INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id",
		task.Title, task.Description,
	).Scan(&id)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Task created successfully",
		"id":      id,
	})
}

/*func updateTask(c *gin.Context) {

}

func deleteTask(c *gin.Context) {
	
}
*/