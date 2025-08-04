package cli

import (
	"errors"
	"os"
	"os/exec"
	"regexp"

	"github.com/fatih/color"
)

const (
	command   = "new"
	usageText = "Usage: echo-mvc new <directory>"
)

func Run() {
	err := validateArgs(os.Args)
	if err != nil {
		color.Red("Error: %s", err)
		color.Cyan(usageText)
		return
	}

	err = cloneTemplateRepository(os.Args[2])
	if err != nil {
		color.Red("Error: %s", err)
	}

	err = installGenerator()
	if err != nil {
		color.Red("Error: %s", err)
	}
}

func validateArgs(args []string) error {
	if len(args) < 3 {
		return errors.New("arguments not provided")
	}

	argCmd := args[1]
	argDir := args[2]

	if argCmd != command {
		return errors.New("command not provided")
	}

	if !isDirectory(argDir) {
		return errors.New("invalid path")
	}

	return nil
}

func isDirectory(path string) bool {
	r, err := regexp.Compile(`^(\.{0,2}[\\/]?|[a-zA-Z0-9_-]+[\\/]?)*[\\/]?$`)
	matched := r.MatchString(path)

	return err == nil && matched
}

func cloneTemplateRepository(dir string) error {
	cmd := exec.Command(
		"git",
		"clone",
		"--depth",
		"1",
		"https://github.com/sk1t0n/echo-mvc-template.git",
		dir,
	)

	_, err := cmd.Output()
	if err != nil {
		return errors.New("failed to clone repository github.com/sk1t0n/echo-mvc-template")
	}

	color.Green("Repository successfully cloned to %s", dir)

	return nil
}

func installGenerator() error {
	cmd := exec.Command("go", "install", "github.com/sk1t0n/echo-mvc-generator@latest")

	_, err := cmd.Output()
	if err != nil {
		return errors.New("failed to install github.com/sk1t0n/echo-mvc-generator")
	}

	color.Green("CLI echo-mvc-generator installed successfully")

	return nil
}
