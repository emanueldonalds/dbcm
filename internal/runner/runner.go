package runner

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/emanueldonalds/dbcm/internal/config"
)

type Runner struct {
	Cfg *config.Config
}

func NewRunner(cfg *config.Config) *Runner {
	return &Runner{
		Cfg: cfg,
	}
}

func (r *Runner) Run() error {
	if err := r.pickConnection(); err != nil {
		return fmt.Errorf("pick connection: %w", err)
	}
	return nil
}

func (r *Runner) pickConnection() error {
	var connectionNames []string

	for _, connection := range r.Cfg.Connections {
		connectionNames = append(connectionNames, connection.Name)
	}

	out , err := cmd("fzf", connectionNames...)
	if err != nil {
		return err
	}

	fmt.Printf("You selected: %s", out)
	return nil
}

func cmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name)
	argStr := strings.Join(args, "\n")
	cmd.Stdin = bytes.NewBufferString(argStr)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("run '%s': %w", name, err)
	}

	return strings.TrimSpace(string(out)), nil
}
