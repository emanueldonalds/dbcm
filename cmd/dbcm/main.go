package main

import (
	"fmt"
	"os"

	"github.com/emanueldonalds/dbcm/internal/config"
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
	loader := config.NewLoader()
	cfg, err := loader.Load()
	if err != nil {
		return fmt.Errorf("could not load config: %w", err)
	}

	fmt.Printf("Loaded %d connections\n", len(cfg.Connections))
	return nil
}
