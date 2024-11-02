package main

import (
	"database/sql"
	"log"
)

// Migration SQL statements
const (
	createEmployeeJobsTable = `
	CREATE TABLE IF NOT EXISTS employee_jobs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			employee_id INTEGER NOT NULL,
			department TEXT NOT NULL,
			job_title TEXT NOT NULL,
			FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE
	);`

	dropEmployeeJobsTable = `DROP TABLE IF EXISTS employee_jobs;`
)

func Migrate(db *sql.DB) error {
	// Create migrations table if it doesn't exist
	_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS migrations (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					name TEXT NOT NULL,
					created_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);
	`)
	if err != nil {
		return err
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if migration has already been applied
	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM migrations WHERE name = ?)",
		"create_employee_jobs_table").Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		// Create the employee_jobs table
		_, err = tx.Exec(createEmployeeJobsTable)
		if err != nil {
			return err
		}

		// Record the migration
		_, err = tx.Exec("INSERT INTO migrations (name) VALUES (?)",
			"create_employee_jobs_table")
		if err != nil {
			return err
		}

		log.Println("Migration 'create_employee_jobs_table' applied successfully")
	} else {
		log.Println("Migration 'create_employee_jobs_table' already applied")
	}

	return tx.Commit()
}

func Rollback(db *sql.DB) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Drop the employee_jobs table
	_, err = tx.Exec(dropEmployeeJobsTable)
	if err != nil {
		return err
	}

	// Remove migration record
	_, err = tx.Exec("DELETE FROM migrations WHERE name = ?",
		"create_employee_jobs_table")
	if err != nil {
		return err
	}

	log.Println("Rollback 'create_employee_jobs_table' completed successfully")

	return tx.Commit()
}
