package model

import "todo-api/types"

type User struct {
	ID        types.UserID `json:"id" gorm:"primaryKey"`
	Email     string       `json:"email" gorm:"unique"`
	Password  string       `json:"password"`
	CreatedAt int64        `json:"created_at"`
}

type UserResponse struct {
	ID    types.UserID `json:"id" gorm:"primaryKey"`
	Email string       `json:"email" gorm:"unique"`
}
