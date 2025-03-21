package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	bb_base_url = "https://bitbucket.org"
	gh_base_url = "https://github.com"
)

func parseFlags() (string, string, string, string, string, string) {
	bb_token := flag.String("bb_token", "", "BitBucket Token/App password")
	gh_token := flag.String("gh_token", "", "GitHub Token/App password")
	bb_workspace := flag.String("bb_workspace", "", "BitBucket Workspace")
	gh_org := flag.String("gh_org", "", "Target Github Organization")
	bb_repo := flag.String("bb_repo", "", "BitBucket Source Repository")
	gh_repo := flag.String("gh_repo", "", "The target GitHub Repo")
	flag.Parse()

	if *bb_token == "" || *gh_token == "" || *bb_workspace == "" || *gh_org == "" || *bb_repo == "" || *gh_repo == "" {
		fmt.Println("Error: required parameters are not passed")
		flag.Usage()
		os.Exit(1)
	}

	return *bb_token, *gh_token, *bb_workspace, *gh_org, *bb_repo, *gh_repo
}

func main() {
	bb_token, gh_token, bb_workspace, gh_org, bb_repo, gh_repo := parseFlags()
	fmt.Printf("%s, %s, %s, %s, %s, %s", bb_token, gh_token, bb_workspace, gh_org, bb_repo, gh_repo)
}
