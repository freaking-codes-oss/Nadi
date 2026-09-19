package persistence

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func Open(path string) (*sql.DB, error) {
	if dir := filepath.Dir(path); dir != "." { if err := os.MkdirAll(dir, 0700); err != nil { return nil, err } }
	db, err := sql.Open("sqlite3", path+"?_foreign_keys=on")
	if err != nil { return nil, err }
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS audit_events (id INTEGER PRIMARY KEY, created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP, event_type TEXT NOT NULL, task_id TEXT, payload TEXT NOT NULL); CREATE TABLE IF NOT EXISTS tasks (id TEXT PRIMARY KEY, created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP, status TEXT NOT NULL, prompt TEXT NOT NULL);`); err != nil { db.Close(); return nil, err }
	return db, nil
}
