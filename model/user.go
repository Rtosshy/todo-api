package model

import "todo-api/types"

type User struct {
	ID        types.UserID `json:"id" gorm:"primaryKey"`
	Email     string       `json:"email" gorm:"unique"`
	Password  string       `json:"password"`
	CreatedAt int64        `json:"created_at" gorm:"autoUpdateTime"`
}

type UserResponse struct {
	ID    types.UserID `json:"id"`
	Email string       `json:"email"`
}
