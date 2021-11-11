package cli

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
)

const (
	VERSION = "0.0.1"
)

func getAllRepositories(user *github.User, client *github.Client, username string) (error, []*github.Repository) {

	if *user.PublicRepos < 100 {
		opt := &github.RepositoryListOptions{
			ListOptions: github.ListOptions{PerPage: *user.PublicRepos},
		}

		repos, _, err := client.Repositories.List(context.Background(), username, opt)

		return err, repos
	}

	var repositories []*github.Repository

	reposCount := 0
	page := 0

	for reposCount < *user.PublicRepos {
		reposToGet := 0

		if *user.PublicRepos-reposCount > 100 {
			reposToGet = 100
			page += 1
		} else {
			reposToGet = *user.PublicRepos - reposCount
		}

		opt := &github.RepositoryListOptions{
			ListOptions: github.ListOptions{PerPage: reposToGet, Page: page},
		}

		repos, _, err := client.Repositories.List(context.Background(), username, opt)

		if err != nil {
			return err, nil
		}

		repositories = append(repositories, repos...)

	}

	return nil, repositories

}

func cloneRepository(url, name string) (string, error) {
	color.Blue("Cloning %s... \n", name)

	cmd := exec.Command("git", "clone", url)

	_, err := cmd.Output()

	if err != nil {
		return name, err
	}

	color.Green("%s cloned successfully\n", name)

	return "", nil
}

func printHelp() {
	helpMessage := `gclal %s`

	color.Green(fmt.Sprintf(helpMessage, VERSION))
}
