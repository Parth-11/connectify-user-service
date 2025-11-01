package repository

import (
	"context"

	"github.com/Parth-11/connectify-user-service/database"
	"github.com/Parth-11/connectify-user-service/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (name,email,password,created_at,updated_at)
			VALUES($1,$2,$,$4,$5)
			RETURNING id`

	return database.DB.QueryRow(ctx, query,
		user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt,
	).Scan(&user.UserID)
}

func GetUserById(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT * FROM users
			WHERE id=$1`

	row := database.DB.QueryRow(ctx, query, id)

	var user models.User
	err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
