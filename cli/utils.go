package cli

import (
	"fmt"
	"github.com/go-git/go-git/v5"
)

func cloneRepository(url, name string) {
	fmt.Printf("Cloning %s... \n", name)

	_, err := git.PlainClone(name, false, &git.CloneOptions{URL: url})

	if err != nil {
		fmt.Printf("Error while cloning %s, killing the process", url)
		fmt.Println(err)
	}

	fmt.Printf("%s cloned successfully\n", name)
}
