package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/555tv/databaze/user_service/internal/models"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUser(
	ctx context.Context,
	id int,
) (*models.User, error) {

	user := &models.User{}

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			first_name,
			last_name,
			email,
			birth_date,
			password_hash,
			role,
			created_at
		FROM users
		WHERE id = $1
		`,
		id,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.BirthDate,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
