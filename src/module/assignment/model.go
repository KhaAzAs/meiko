package assignment

import (
	"database/sql"
	"time"
)

// Assignment is a model of assignments table
type Assignment struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Status     int8      `db:"status"`
	UploadDate time.Time `db:"upload_date"`
	DueDate    time.Time `db:"due_date"`
}

// GradeParameter is a model of grade_parameters table
type GradeParameter struct {
	ID         int64  `db:"id"`
	Type       string `db:"type"`
	Percentage int8   `db:"percentage"`
}

// CompleteAssignment is a model of p_users_assignments table
type CompleteAssignment struct {
	AssignmentID int64          `db:"assignments_id"`
	UserID       int64          `db:"users_id"`
	Description  sql.NullString `db:"description"`
	Score        float32        `db:"score"`
}
