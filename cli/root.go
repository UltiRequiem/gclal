package cli

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"sync"
)

func Init() {
	username := "UltiRequiem"

	client := github.NewClient(nil)

	user, _, error := client.Users.Get(context.Background(), username)

	if error != nil {
		fmt.Println(error)
	}

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: *user.PublicRepos},
	}

	repos, _, err := client.Repositories.List(context.Background(), username, opt)

	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup

	for _, repo := range repos {
		wg.Add(1)
		go func(r *github.Repository, wg *sync.WaitGroup) {
			defer wg.Done()
			cloneRepository(*r.Name, *r.GitURL)
		}(repo, &wg)
	}

	wg.Wait()
}
