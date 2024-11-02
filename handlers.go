package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// HTTP Status Codes
const (
	// Success Codes
	StatusOK        = 200 // For GET and successful retrievals
	StatusCreated   = 201 // For POST when resource is created
	StatusNoContent = 204 // For successful updates

	// Client Error Codes
	StatusBadRequest = 400 // Invalid request format/syntax
	StatusNotFound   = 404 // Resource not found
	StatusConflict   = 409 // Resource conflict (e.g., duplicate entry)

	// Server Error Code
	StatusInternalServerError = 500 // Server-side errors
)

// Handler struct holds dependencies
type Handler struct {
	repo *Repository
}

// NewHandler creates a new handler instance
func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

// Handler methods
func (h *Handler) GetEmployee(c *gin.Context) {
	id := c.Param("id")

	emp, err := h.repo.GetEmployee(id)
	if err != nil {
		c.JSON(StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if emp == nil {
		c.JSON(StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(StatusOK, emp)
}

func (h *Handler) GetAllEmployees(c *gin.Context) {
	employees, err := h.repo.GetAllEmployees()
	if err != nil {
		c.JSON(StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(StatusOK, employees)
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var emp Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.repo.CreateEmployee(&emp); err != nil {
		c.JSON(StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(StatusCreated, emp)
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	// Parse the partial update request
	var update EmployeeUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate that at least one field is being updated
	if update.Department == nil && update.JobTitle == nil {
		c.JSON(StatusBadRequest, gin.H{"error": "No fields to update"})
		return
	}

	err := h.repo.UpdateEmployee(id, &update)
	if err == sql.ErrNoRows {
		c.JSON(StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	if err != nil {
		c.JSON(StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(StatusNoContent, nil)
}
