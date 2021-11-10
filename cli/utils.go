package cli

import (
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

func cloneRepository(url, name string) {
	color.Blue("Cloning %s... \n", name)

	_, err := git.PlainClone(name, false, &git.CloneOptions{URL: url})

	if err != nil {
		color.Red("Error while cloning %s, killing the process", url)
		color.Red(err.Error())
		return
	}

	color.Green("%s cloned successfully\n", name)
}
