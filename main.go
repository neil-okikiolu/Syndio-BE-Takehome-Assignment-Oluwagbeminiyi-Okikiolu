package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine, h *Handler) {
	r.GET("/employees", h.GetAllEmployees)
	r.GET("/employees/:id", h.GetEmployee)
	r.POST("/employees", h.CreateEmployee)
	r.PATCH("/employees/:id", h.UpdateEmployee)
}

func main() {

	db, err := sql.Open("sqlite3", "./employees.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create repository and handler
	repo := NewRepository(db)
	handler := NewHandler(repo)

	// Setup Gin router
	r := gin.Default()

	// Setup routes
	setupRoutes(r, handler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port) // listen and serve on
}
