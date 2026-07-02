package main

import (
	"fmt"
	"os"

	"github.com/emanueldonalds/dbcm/internal/config"
)


func main() {
	fmt.Printf("  _//_ _  _ _\n")
	fmt.Printf("/_//_//_ / / /\n")

	loader := config.NewLoader()
	cfg, err := loader.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not load config. %v\n", err)
	}

	fmt.Printf("Loaded %d connections\n", len(cfg.Connections))
}
