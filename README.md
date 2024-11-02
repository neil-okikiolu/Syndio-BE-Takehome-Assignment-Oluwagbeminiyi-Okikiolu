# Employee Jobs API

A RESTful API service built with Go and Gin framework for managing employee job information. The service uses SQLite as its database.

## Prerequisites

- Go 1.16 or higher
- SQLite3

## Project Setup

1. Clone the repository
2. Navigate to the project directory
3. Initialize Go modules:

```bash
go mod init employee-jobs
```

4. Install dependencies:

```bash
go mod tidy
```

This will install the required packages:

- github.com/gin-gonic/gin
- github.com/ncruces/go-sqlite3

## Building and Running

1. Build the project:

```bash
go build
```

2. Run the application:

```bash
./employee-jobs
```

The server will start on port 8080 by default. You can set a different port using the `PORT` environment variable:

```bash
PORT=3000 ./employee-jobs
```

## API Endpoints

### Get All Employee Jobs

```
GET /employee/jobs
```

Response (200 OK):

```json
[
  {
    "id": 1,
    "employee_id": 100,
    "department": "Engineering",
    "job_title": "Software Engineer"
  },
  {
    "id": 2,
    "employee_id": 101,
    "department": "Marketing",
    "job_title": "Marketing Manager"
  }
]
```

### Get Single Employee Job

```
GET /employee/jobs/:id
```

Response (200 OK):

```json
{
  "id": 1,
  "employee_id": 100,
  "department": "Engineering",
  "job_title": "Software Engineer"
}
```

### Create Employee Job

```
POST /employee/jobs
```

Request Body:

```json
{
  "employee_id": 102,
  "department": "Sales",
  "job_title": "Sales Representative"
}
```

Response (201 Created):

```json
{
  "id": 3,
  "employee_id": 102,
  "department": "Sales",
  "job_title": "Sales Representative"
}
```

### Update Employee Job

```
PATCH /employee/jobs/:id
```

Request Body (partial updates supported):

```json
{
  "department": "Business Development",
  "job_title": "Senior Sales Representative"
}
```

Response (204 No Content)

## API Testing Examples

Using cURL:

1. Create a new employee job:

```bash
curl -X POST http://localhost:8080/employee/jobs \
  -H "Content-Type: application/json" \
  -d '{
    "employee_id": 100,
    "department": "Engineering",
    "job_title": "Software Engineer"
  }'
```

2. Get all employee jobs:

```bash
curl http://localhost:8080/employee/jobs
```

3. Get a specific employee job:

```bash
curl http://localhost:8080/employee/jobs/1
```

4. Update an employee job:

```bash
curl -X PATCH http://localhost:8080/employee/jobs/1 \
  -H "Content-Type: application/json" \
  -d '{
    "job_title": "Senior Software Engineer"
  }'
```

## Error Responses

The API returns the following error codes:

- `400 Bad Request`: Invalid request payload
- `404 Not Found`: Employee job not found
- `500 Internal Server Error`: Server-side errors

Error Response Format:

```json
{
  "error": "Error message here"
}
```

## Database Schema

The application uses SQLite with the following schema:

```sql
CREATE TABLE employee_jobs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    employee_id INTEGER NOT NULL,
    department TEXT NOT NULL,
    job_title TEXT NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE
);
```

Indexes are created for `employee_id`, `department`, and `job_title` columns for better query performance.

# Syndio Backend App (Original Instructions)

Using the `employees.db` sqlite database in this repository with the following table/data:

```
sqlite> .open employees.db
sqlite> .schema employees
CREATE TABLE employees (id INTEGER PRIMARY KEY, gender TEXT not null);
sqlite> SELECT * FROM employees;
1|male
2|male
3|male
4|female
5|female
6|female
7|non-binary
```

Create an API endpoint that saves job data for the corresponding employees.

Example job data:

```json
[
  {
    "employee_id": 1,
    "department": "Engineering",
    "job_title": "Senior Enginer"
  },
  {
    "employee_id": 2,
    "department": "Engineering",
    "job_title": "Super Senior Enginer"
  },
  { "employee_id": 3, "department": "Sales", "job_title": "Head of Sales" },
  { "employee_id": 4, "department": "Support", "job_title": "Tech Support" },
  {
    "employee_id": 5,
    "department": "Engineering",
    "job_title": "Junior Enginer"
  },
  { "employee_id": 6, "department": "Sales", "job_title": "Sales Rep" },
  {
    "employee_id": 7,
    "department": "Marketing",
    "job_title": "Senior Marketer"
  }
]
```

## Requirements

- The API must take an environment variable `PORT` and respond to requests on that port.
- You provide:
  - Basic setup instructions required to run the API
  - Guide on how to ingest the data through the endpoint
  - A way to update the existing database given to you

## Success

- We can run the API and ingest database on your setup instructions
- The API is written in Python or Go

## Not Required

- Tests
- Logging, monitoring, or anything more than basic error handling

## Submission

- Respond to the email you received giving you this with:
  - a zip file, or link to a git repo
  - instructions on how to setup and run the code (could be included w/ zip/git)
- We'll follow the instructions to test it on a local machine, then we'll get back to you

## Notes

- Keep it simple
- If the API does what we requested, then it's a success
- Anything extra (tests, other endpoints, ...) is not worth bonus/etc
- We expect this to take less than two hours, please try and limit your effort to that window
- We truly value your time and just want a basic benchmark and common piece of code to use in future interviews
- If we bring you in for in-person interviews, your submission might be revisited and built on during the interview process
