package config

import (
	"os"
	"path/filepath"
)

func DefaultDBPath() string {
	if v := os.Getenv("NADI_DB"); v != "" { return v }
	home, err := os.UserHomeDir(); if err != nil { return "nadi.db" }
	return filepath.Join(home, ".nadi", "nadi.db")
}
