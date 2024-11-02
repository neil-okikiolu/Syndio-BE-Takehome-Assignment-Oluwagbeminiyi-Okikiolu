package main

import (
	"database/sql"
	"strings"
)

// Employee model
type Employee struct {
	EmployeeId int    `json:"employee_id"`
	Department string `json:"department"`
	JobTitle   string `json:"job_title"`
}

// EmployeeUpdate payload
type EmployeeUpdate struct {
	Department *string `json:"department,omitempty"`
	JobTitle   *string `json:"job_title,omitempty"`
}

// Repository handles database operations
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new repository instance
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Repository methods
func (r *Repository) GetEmployee(id string) (*Employee, error) {
	var emp Employee
	err := r.db.QueryRow("SELECT employee_id, department, job_title FROM employees WHERE id = ?", id).Scan(&emp.EmployeeId, &emp.Department, &emp.JobTitle)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *Repository) GetAllEmployees() ([]Employee, error) {
	rows, err := r.db.Query("SELECT employee_id, department, job_title FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.EmployeeId, &emp.Department, &emp.JobTitle); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

func (r *Repository) CreateEmployee(emp *Employee) error {
	result, err := r.db.Exec("INSERT INTO employees (department, job_title) VALUES (?)", emp.Department, emp.JobTitle)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	emp.EmployeeId = int(id)
	return nil
}

func (r *Repository) UpdateEmployee(id string, update *EmployeeUpdate) error {
	// Build dynamic query based on provided fields
	query := "UPDATE employees SET"
	var args []interface{}
	var setClauses []string

	if update.Department != nil {
		setClauses = append(setClauses, " department = ?")
		args = append(args, *update.Department)
	}

	if update.JobTitle != nil {
		setClauses = append(setClauses, " job_title = ?")
		args = append(args, *update.JobTitle)
	}

	// If no fields to update, return early
	if len(setClauses) == 0 {
		return nil
	}

	// Combine set clauses with commas
	query += strings.Join(setClauses, ",")
	query += " WHERE employee_id = ?"
	args = append(args, id)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
