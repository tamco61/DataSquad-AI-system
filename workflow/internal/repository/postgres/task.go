package postgres

import (
	"context"
	"fmt"
	"workflow/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Task struct {
	pool *pgxpool.Pool
}

func NewTaskRepo(pool *pgxpool.Pool) *Task {
	return &Task{pool: pool}
}

func (r *Task) CreateTask(ctx context.Context, t *model.TaskDTO) error {
	query := `
		INSERT INTO tasks (
		 	id, title, raw_text,
			classification, summary, risk_level,
			assignee_id, addressed_id, status,
			sla_deadline, resolved_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.pool.Exec(ctx, query,
		t.ID, t.Title, t.RawText,
		t.Classification, t.Summary, t.RiskLevel,
		t.AssigneeID, t.AddressedID, t.Status,
		t.SLADeadline, t.ResolvedAt,
	)

	return err
}

func (r *Task) GetTask(ctx context.Context, id string) (*model.TaskDTO, error) {
	query := `
		SELECT
			id, title, raw_text,
			classification, summary, risk_level,
			assignee_id, addressed_id, status,
			sla_deadline, resolved_at,
			created_at, updated_at
		FROM tasks
		WHERE id = $1
	`

	var dto model.TaskDTO

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&dto.ID,
		&dto.Title,
		&dto.RawText,
		&dto.Classification,
		&dto.Summary,
		&dto.RiskLevel,
		&dto.AssigneeID,
		&dto.AddressedID,
		&dto.Status,
		&dto.SLADeadline,
		&dto.ResolvedAt,
		&dto.CreatedAt,
		&dto.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *Task) GetTasksFiltering(ctx context.Context, f model.TaskFilter) (*[]model.TaskDTO, error) {
	query := `
		SELECT id, title, raw_text, classification, summary, risk_level,
		       assignee_id, status, sla_deadline, resolved_at,
		       created_at, updated_at
		FROM tasks
		WHERE 1=1
	`

	args := []interface{}{}
	argIdx := 1

	if f.Status != "" {
		query += fmt.Sprintf(" AND status = $%d", argIdx)
		args = append(args, f.Status)
		argIdx++
	}

	if f.AssigneeID != "" {
		query += fmt.Sprintf(" AND assignee_id = $%d", argIdx)
		args = append(args, f.AssigneeID)
		argIdx++
	}

	query += " ORDER BY created_at DESC"

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.TaskDTO

	for rows.Next() {
		var dto model.TaskDTO

		err := rows.Scan(
			&dto.ID,
			&dto.Title,
			&dto.RawText,
			&dto.Classification,
			&dto.Summary,
			&dto.RiskLevel,
			&dto.AssigneeID,
			&dto.Status,
			&dto.SLADeadline,
			&dto.ResolvedAt,
			&dto.CreatedAt,
			&dto.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, dto)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return &tasks, nil
}
