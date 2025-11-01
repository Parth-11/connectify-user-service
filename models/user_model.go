package models

import "time"

type User struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Bio       string    `json:"bio"`
	AvatarUrl string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
}

//User DTO

type UserDTO struct {
}
