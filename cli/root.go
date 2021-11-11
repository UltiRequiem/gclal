package cli

import (
	"context"
	"fmt"
	"sync"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
)

func Init() {
	username := getParams()

	if username == "default" {
		color.Red("Please enter a username.")
		return
	}

	client := github.NewClient(nil)

	user, _, error := client.Users.Get(context.Background(), username)

	if error != nil {
		color.Red(error.Error())
		return
	}

	if *user.PublicRepos == 0 {
		color.Green(fmt.Sprintf("No public repos found for user %s.", *user.Name))
		return
	}

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: *user.PublicRepos},
	}

	repos, _, err := client.Repositories.List(context.Background(), username, opt)

	if err != nil {
		color.Red(err.Error())
		return
	}

	color.Green(fmt.Sprintf("About to clone %d repos.", len(repos)))

	var wg sync.WaitGroup
	wg.Add(len(repos))

	for _, repo := range repos {
		go func(r *github.Repository, wg *sync.WaitGroup) {
			defer wg.Done()
			cloneRepository(*r.GitURL, *r.Name)
		}(repo, &wg)
	}

	wg.Wait()
}
