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
	r.GET("/employee/jobs", h.GetAllEmployees)
	r.GET("/employee/jobs/:id", h.GetEmployee)
	r.POST("/employee/jobs", h.CreateEmployee)
	r.PATCH("/employee/jobs/:id", h.UpdateEmployee)
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

	// Run migrations (if needed)
	if err := Migrate(db); err != nil {
		log.Fatal("Failed to run migration:", err)
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

	r.Run(":" + port) // listen and serve on 8080 or specified port
}
