package models

import "time"

type User struct {
	UserID    string
	Name      string
	UserName  string
	Email     string
	Password  string
	Bio       string
	AvatarUrl string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLogin struct {
}

//User DTO

type UserDTO struct {
}
