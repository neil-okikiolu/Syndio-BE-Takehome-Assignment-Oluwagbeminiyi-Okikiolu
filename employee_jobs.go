package main

import (
	"database/sql"
	"strings"
)

type EmployeeStuct struct {
	Id     int    `json:"id"`
	Gender string `json:"gender"`
}

// EmployeeJob model
type EmployeeJob struct {
	Id         int    `json:"id"`
	EmployeeId int    `json:"employee_id"`
	Department string `json:"department"`
	JobTitle   string `json:"job_title"`
}

// EmployeeJobUpdate payload
type EmployeeJobUpdate struct {
	EmployeeId *int    `json:"employee_id,omitempty"`
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
func (r *Repository) GetEmployeeJob(id string) (*EmployeeJob, error) {
	var emp EmployeeJob
	err := r.db.QueryRow("SELECT id, employee_id, department, job_title FROM employee_jobs WHERE id = ?", id).Scan(&emp.Id, &emp.EmployeeId, &emp.Department, &emp.JobTitle)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *Repository) GetAllEmployeeJobs() ([]EmployeeJob, error) {
	rows, err := r.db.Query("SELECT id, employee_id, department, job_title FROM employee_jobs")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employee_jobs := make([]EmployeeJob, 0) // Key part: initialize with make() and length 0
	for rows.Next() {
		var emp EmployeeJob
		if err := rows.Scan(&emp.Id, &emp.EmployeeId, &emp.Department, &emp.JobTitle); err != nil {
			return nil, err
		}
		employee_jobs = append(employee_jobs, emp)
	}
	return employee_jobs, nil
}

func (r *Repository) CreateEmployee(emp *EmployeeJob) error {
	result, err := r.db.Exec("INSERT INTO employee_jobs (employee_id, department, job_title) VALUES (?, ?, ?)", emp.EmployeeId, emp.Department, emp.JobTitle)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	emp.Id = int(id)
	return nil
}

func (r *Repository) UpdateEmployee(id string, update *EmployeeJobUpdate) error {
	// Build dynamic query based on provided fields
	query := "UPDATE employee_jobs SET"
	var args []interface{}
	var setClauses []string

	if update.EmployeeId != nil {
		setClauses = append(setClauses, " employee_id = ?")
		args = append(args, *update.EmployeeId)
	}

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
	query += " WHERE id = ?"
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
