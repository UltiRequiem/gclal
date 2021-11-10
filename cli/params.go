package cli

import "flag"

func getParams() string {
	username := flag.String("username", "default", "The GitHub username.")

	flag.Parse()

	return *username
}
