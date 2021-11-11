package cli

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

const (
	VERSION = "0.0.1"
)

func cloneRepository(url, name string) error {
	color.Blue("Cloning %s... \n", name)

	_, err := git.PlainClone(name, false, &git.CloneOptions{URL: url})

	if err != nil {
		color.Red(`Error while cloning %s: "%s".`, url, err.Error())
		return errors.New(name)
	}

	color.Green("%s cloned successfully\n", name)

        return nil
}

func printHelp() {
	helpMessage := `gclal %s`

	color.Green(fmt.Sprintf(helpMessage, VERSION))
}
