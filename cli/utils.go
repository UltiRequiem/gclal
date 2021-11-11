package cli

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os/exec"
)

const (
	VERSION = "0.0.1"
)

func cloneRepository(url, name string) error {
	color.Blue("Cloning %s... \n", name)

	cmd := exec.Command("git", "clone", url)

	_, err := cmd.Output()

	if err != nil {
		return errors.New(name)
	}

	color.Green("%s cloned successfully\n", name)

	return nil
}

func printHelp() {
	helpMessage := `gclal %s`

	color.Green(fmt.Sprintf(helpMessage, VERSION))
}
