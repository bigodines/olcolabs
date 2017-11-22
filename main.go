package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	// list all organizations for user "willnorris"
	opt := &github.RepositoryListByOrgOptions{Type: "private"}
	repos, _, err := client.Repositories.ListByOrg(ctx, "gametimesf", opt)
	if err != nil {
		println("fuein")
		println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%+v", repos[0])
}
