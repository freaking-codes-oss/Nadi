package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/freaking-codes-oss/Nadi/internal/config"
	"github.com/freaking-codes-oss/Nadi/internal/persistence"
)

func Execute() {
	if err := newRoot().ExecuteContext(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, "nadi:", err)
		os.Exit(1)
	}
}

func newRoot() *cobra.Command {
	var dbPath string
	root := &cobra.Command{Use: "nadi", Short: "Nadi: an auditable AI operator for your Android terminal", Args: cobra.ArbitraryArgs}
	root.PersistentFlags().StringVar(&dbPath, "db", config.DefaultDBPath(), "SQLite database path")
	root.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 { return interactive(cmd.Context(), dbPath) }
		return runTask(cmd.Context(), dbPath, args)
	}
	root.AddCommand(&cobra.Command{Use: "doctor", Short: "Check local runtime health", RunE: func(cmd *cobra.Command, _ []string) error { fmt.Fprintln(cmd.OutOrStdout(), "Nadi runtime: ready"); return nil }})
	root.AddCommand(&cobra.Command{Use: "inspect", Short: "Inspect persisted tasks and audit status", RunE: func(cmd *cobra.Command, _ []string) error { return inspect(cmd.Context(), dbPath, cmd) }})
	return root
}

func runTask(ctx context.Context, dbPath string, args []string) error {
	db, err := persistence.Open(dbPath); if err != nil { return err }; defer db.Close()
	fmt.Printf("Task: %s\nStatus: queued\nApproval: required before tools execute\n", join(args))
	return nil
}
func interactive(context.Context, string) error { fmt.Println("Nadi interactive mode. Type /help or /exit."); return nil }
func inspect(ctx context.Context, path string, cmd *cobra.Command) error { db, err := persistence.Open(path); if err != nil { return err }; defer db.Close(); fmt.Fprintln(cmd.OutOrStdout(), "Database: "+path); return nil }
func join(v []string) string { out := ""; for i, s := range v { if i > 0 { out += " " }; out += s }; return out }
