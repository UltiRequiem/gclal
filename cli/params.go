package cli

import "flag"

func getParams() (string, bool) {
	username := flag.String("username", "default", "The GitHub username.")
	ssh := flag.Bool("ssh", false, "Use SSH?")

	flag.Usage = printHelp

	flag.Parse()

	return *username, *ssh
}
