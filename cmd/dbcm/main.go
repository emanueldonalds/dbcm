package main

import (
	"fmt"
	"os"

	"github.com/emanueldonalds/dbcm/internal/config"
	"github.com/emanueldonalds/dbcm/internal/runner"
	"github.com/emanueldonalds/dbcm/internal/deps"
)

func main() {
	fmt.Printf("  _//_ _  _ _\n")
	fmt.Printf("/_//_//_ / / /\n")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

}

func run() error {
	if err := deps.Check(); err != nil {
		return fmt.Errorf("%w", err)
	}

	l := config.NewLoader()
	cfg, err := l.Load()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	r := runner.NewRunner(cfg)
	if err := r.Run(); err != nil {
		return fmt.Errorf("runner: %w", err)
	}

	return nil
}
