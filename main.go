package main

import (
	"context"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/github"
	"sync"
)

func main() {
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
			fmt.Printf("Cloning %s... \n", *r.Name)

			_, err := git.PlainClone(*r.Name, false, &git.CloneOptions{URL: *r.GitURL})

			if err != nil {
				fmt.Printf("Error while cloning %s, killing the process", *r.GitURL)
			}

			fmt.Printf("%s cloned successfully\n", *r.Name)

		}(repo, &wg)
	}

        wg.Wait()
}
