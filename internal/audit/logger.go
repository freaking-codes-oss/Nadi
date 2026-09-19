package audit

import (
	"context"
	"database/sql"
	"encoding/json"
)

type Logger struct { DB *sql.DB }
func (l Logger) Record(ctx context.Context, eventType, taskID string, payload any) error {
	b, err := json.Marshal(payload); if err != nil { return err }
	_, err = l.DB.ExecContext(ctx, `INSERT INTO audit_events(event_type, task_id, payload) VALUES (?, ?, ?)`, eventType, taskID, string(b))
	return err
}
