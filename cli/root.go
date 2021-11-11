package cli

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"strings"
	"sync"
)

func Init() {
	username, ssh, apiKey := getParams()

	ctx := context.Background()

	var client *github.Client

	if apiKey != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: apiKey},
		)

		client = github.NewClient(oauth2.NewClient(ctx, ts))
	} else {
		client = github.NewClient(nil)
	}

	if username == "default" {
		color.Red("Please enter a username.")
		return
	}

	user, _, error := client.Users.Get(ctx, username)

	if error != nil {
		color.Red(error.Error())
		return
	}

	if *user.PublicRepos == 0 {
		color.Green(fmt.Sprintf("No public repos found for user %s.", *user.Name))
		return
	}

	err, repos := getAllRepositories(user, client, username)

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

			name, err := cloneRepository(url, *r.Name)

			if err != nil {
				failedRepos = append(failedRepos, name)
				fmt.Println(err)
			}

		}(repo, &wg)
	}

	wg.Wait()

	color.Blue(fmt.Sprintf("There was a problem while cloning:\n [%s]. \n", strings.Join(failedRepos, " ")))
}
