package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
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

	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}
}
