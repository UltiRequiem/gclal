package cli

import "flag"

func getParams() (string, bool, string) {
	username := flag.String("username", "default", "The GitHub username.")
	ssh := flag.Bool("ssh", false, "Use SSH?")
	apiKey := flag.String("apiKey", "", "The GitHub API Key, useful if you have more than 100 Repos.")

	flag.Usage = printHelp

	flag.Parse()

	return *username, *ssh, *apiKey
}
