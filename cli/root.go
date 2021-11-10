package cli

import (
	"context"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"sync"
)

func Init() {
	username := "UltiRequiem"

	client := github.NewClient(nil)

	user, _, error := client.Users.Get(context.Background(), username)

	if error != nil {
		color.Red(error.Error())
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

	var wg sync.WaitGroup

	wg.Add(*user.PublicRepos)

	for _, repo := range repos {
		go func(r *github.Repository, wg *sync.WaitGroup) {
			defer wg.Done()
			cloneRepository(*r.GitURL, *r.Name)
		}(repo, &wg)
	}

	wg.Wait()
}
