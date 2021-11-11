package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

const (
  VERSION = "0.0.1"
)

func cloneRepository(url, name string) {
	color.Blue("Cloning %s... \n", name)

	_, err := git.PlainClone(name, false, &git.CloneOptions{URL: url})

	if err != nil {
		color.Red(`Error while cloning %s: "%s".`, url, err.Error())
		return
	}

	color.Green("%s cloned successfully\n", name)
}

func printHelp() {
	helpMessage := `gclal %s`

	color.Green(fmt.Sprintf(helpMessage, VERSION))
}
