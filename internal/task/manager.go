package task

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Manager struct { DB *sql.DB }

func (m Manager) Create(ctx context.Context, prompt string) (string, error) {
	id := fmt.Sprintf("task-%d", time.Now().UnixNano())
	_, err := m.DB.ExecContext(ctx, `INSERT INTO tasks(id,status,prompt) VALUES(?,?,?)`, id, "queued", prompt)
	return id, err
}

func (m Manager) SetStatus(ctx context.Context, id, status string) error {
	_, err := m.DB.ExecContext(ctx, `UPDATE tasks SET status=? WHERE id=?`, status, id)
	return err
}
