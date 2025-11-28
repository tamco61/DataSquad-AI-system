package postgres

import (
	"context"
	"workflow/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *User {
	return &User{pool: pool}
}

func (r *User) CreateUser(ctx context.Context, u *model.UserDTO) error {
	query := `
		INSERT INTO users (id, email, name, departament)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.pool.Exec(ctx, query, u.ID, u.Email, u.Name, u.Departament)
	return err
}

func (r *User) GetUser(ctx context.Context, id string) (*model.UserDTO, error) {
	query := `
		SELECT id, email, name, departament, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var dto model.UserDTO

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&dto.ID,
		&dto.Email,
		&dto.Name,
		&dto.Departament,
		&dto.CreatedAt,
		&dto.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *User) GetUsers(ctx context.Context) (*[]model.UserDTO, error) {
	query := `
		SELECT id, email, name, departament, created_at, updated_at
		FROM users
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.UserDTO{}

	for rows.Next() {
		var dto model.UserDTO

		err := rows.Scan(
			&dto.ID,
			&dto.Email,
			&dto.Name,
			&dto.Departament,
			&dto.CreatedAt,
			&dto.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, dto)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &users, nil
}
