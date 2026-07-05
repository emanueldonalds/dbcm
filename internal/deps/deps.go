package deps

import (
	"fmt"
	"os/exec"
)

var dependencies = []string{
	"fzf",
}

func Check() error {
	for _, dependency := range dependencies {
		_, err := exec.LookPath(dependency)
		if err != nil {
			return fmt.Errorf("Required dependency not found: %w", err)
		}
	}
	return nil
}
