package model

import (
	"time"
	"todo-api/types"
)

type Task struct {
	ID        types.TaskID `json:"id" gorm:"primaryKey"`
	Title     string       `json:"title" gorm:"not null"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      User         `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	UserID    types.UserID `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID        types.TaskID `json:"id" gorm:"primaryKey"`
	Title     string       `json:"title" gorm:"not null"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
