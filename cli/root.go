package cli

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"strings"
	"sync"
)

func Init() {
	username, ssh := getParams()

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

	color.Green(fmt.Sprintf("About to clone %d repos!", len(repos)))

	var wg sync.WaitGroup
	wg.Add(len(repos))

	failedRepos := []string{}

	for _, repo := range repos {
		go func(r *github.Repository, wg *sync.WaitGroup) {
			defer wg.Done()

			url := ""

			if ssh {
				url = *r.SSHURL
			} else {
				url = *r.GitURL
			}

			err := cloneRepository(url, *r.Name)

			if err != nil {
				failedRepos = append(failedRepos, err.Error())
			}

		}(repo, &wg)
	}

	wg.Wait()

	color.Blue(fmt.Sprintf("There was a problem while cloning: [%s]. \n", strings.Join(failedRepos, " ")))
}
