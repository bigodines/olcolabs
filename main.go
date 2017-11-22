package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
)

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)
	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List(ctx, "bigodines", nil)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%+v", orgs)
}
