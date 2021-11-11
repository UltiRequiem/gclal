package cli

import "flag"

func getParams() string {
	username := flag.String("username", "default", "The GitHub username.")

	flag.Usage = printHelp

	flag.Parse()

	return *username
}
