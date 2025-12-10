package models

import "time"

type Role string

const (
	Admin          Role = "admin"
	Developer      Role = "developer"
	ProjectManager Role = "project_manager"
)

type User struct {
	ID             string `json:"id" bson:"_id,omitempty"`
	Name           string `json:"name" bson:"name"`
	Email          string `json:"email" bson:"email"`
	PasswordHashed string `json:"-" bson:"password_hashed"`
	Role           Role   `json:"role" bson:"role"`
}

type TaskStatus string

const (
	ToDo                   TaskStatus = "to_do"
	InProgress             TaskStatus = "in_progress"
	ReadyForImplementation TaskStatus = "ready_for_implementation"
	Done                   TaskStatus = "done"
)

type Task struct {
	ID          string     `json:"id" bson:"_id,omitempty"`
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	Status      TaskStatus `json:"status" bson:"status"`
	AssigneeID  string     `json:"assignee_id" bson:"assignee_id"`
	ProjectID   string     `json:"project_id" bson:"project_id"`
	SprintID    *string    `json:"sprint_id,omitempty" bson:"sprint_id,omitempty"`
	CreatedAt   time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" bson:"updated_at"`
}

type Project struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	OwnerID     string    `json:"owner_id" bson:"owner_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type SprintStatus string

const (
	SprintCreated SprintStatus = "created"
	SprintActive  SprintStatus = "active"
	SprintClosed  SprintStatus = "closed"
)

type Sprint struct {
	ID        string       `json:"id" bson:"_id,omitempty"`
	ProjectID string       `json:"project_id" bson:"project_id"`
	Name      string       `json:"name" bson:"name"`
	StartDate time.Time    `json:"start_date" bson:"start_date"`
	EndDate   time.Time    `json:"end_date" bson:"end_date"`
	Status    SprintStatus `json:"status" bson:"status"`
	CreatedAt time.Time    `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" bson:"updated_at"`
}
