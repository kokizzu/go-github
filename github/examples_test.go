// Copyright 2016 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// These examples are inlined in godoc.

package github_test

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v74/github"
)

func ExampleMarkdownService_Render() {
	client := github.NewClient(nil)

	input := "# heading #\n\nLink to issue #1"
	opt := &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}

	ctx := context.Background()
	output, _, err := client.Markdown.Render(ctx, input, opt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
}

func ExampleRepositoriesService_GetReadme() {
	client := github.NewClient(nil)

	ctx := context.Background()
	readme, _, err := client.Repositories.GetReadme(ctx, "google", "go-github", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := readme.GetContent()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("google/go-github README:\n%v\n", content)
}

func ExampleRepositoriesService_ListByUser() {
	client := github.NewClient(nil)

	user := "willnorris"
	opt := &github.RepositoryListByUserOptions{Type: "owner", Sort: "updated", Direction: "desc"}

	ctx := context.Background()
	repos, _, err := client.Repositories.ListByUser(ctx, user, opt)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Recently updated repositories by %q: %v", user, github.Stringify(repos))
}

func ExampleRepositoriesService_CreateFile() {
	// In this example we're creating a new file in a repository using the
	// Contents API. Only 1 file per commit can be managed through that API.

	// Note that authentication is needed here as you are performing a modification
	// so you will need to modify the example to provide an oauth client to
	// github.NewClient() instead of nil. See the following documentation for more
	// information on how to authenticate with the client:
	// https://pkg.go.dev/github.com/google/go-github/github#hdr-Authentication
	client := github.NewClient(nil)

	ctx := context.Background()
	fileContent := []byte("This is the content of my file\nand the 2nd line of it")

	// Note: the file needs to be absent from the repository as you are not
	// specifying a SHA reference here.
	opts := &github.RepositoryContentFileOptions{
		Message:   github.Ptr("This is my commit message"),
		Content:   fileContent,
		Branch:    github.Ptr("master"),
		Committer: &github.CommitAuthor{Name: github.Ptr("FirstName LastName"), Email: github.Ptr("user@example.com")},
	}
	_, _, err := client.Repositories.CreateFile(ctx, "myOrganization", "myRepository", "myNewFile.md", opts)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleUsersService_ListAll() {
	client := github.NewClient(nil)
	opts := &github.UserListOptions{}
	for {
		ctx := context.Background()
		users, _, err := client.Users.ListAll(ctx, opts)
		if err != nil {
			log.Fatalf("error listing users: %v", err)
		}
		if len(users) == 0 {
			break
		}
		opts.Since = *users[len(users)-1].ID
		// Process users...
	}
}

func ExamplePullRequestsService_Create() {
	// In this example we're creating a PR and displaying the HTML url at the end.

	// Note that authentication is needed here as you are performing a modification
	// so you will need to modify the example to provide an oauth client to
	// github.NewClient() instead of nil. See the following documentation for more
	// information on how to authenticate with the client:
	// https://pkg.go.dev/github.com/google/go-github/github#hdr-Authentication
	client := github.NewClient(nil)

	newPR := &github.NewPullRequest{
		Title:               github.Ptr("My awesome pull request"),
		Head:                github.Ptr("branch_to_merge"),
		Base:                github.Ptr("master"),
		Body:                github.Ptr("This is the description of the PR created with the package `github.com/google/go-github/github`"),
		MaintainerCanModify: github.Ptr(true),
	}

	ctx := context.Background()
	pr, _, err := client.PullRequests.Create(ctx, "myOrganization", "myRepository", newPR)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("PR created: %s\n", pr.GetHTMLURL())
}

func ExampleTeamsService_ListTeams() {
	// This example shows how to get a team ID corresponding to a given team name.

	// Note that authentication is needed here as you are performing a lookup on
	// an organization's administrative configuration, so you will need to modify
	// the example to provide an oauth client to github.NewClient() instead of nil.
	// See the following documentation for more information on how to authenticate
	// with the client:
	// https://pkg.go.dev/github.com/google/go-github/github#hdr-Authentication
	client := github.NewClient(nil)

	teamName := "Developers team"
	ctx := context.Background()
	opts := &github.ListOptions{}

	for {
		teams, resp, err := client.Teams.ListTeams(ctx, "myOrganization", opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, t := range teams {
			if t.GetName() == teamName {
				fmt.Printf("Team %q has ID %d\n", teamName, t.GetID())
				return
			}
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	fmt.Printf("Team %q was not found\n", teamName)
}

func ExampleUsersService_ListUserSocialAccounts() {
	client := github.NewClient(nil)
	ctx := context.Background()
	opts := &github.ListOptions{}
	for {
		accounts, resp, err := client.Users.ListUserSocialAccounts(ctx, "shreyjain13", opts)
		if err != nil {
			log.Fatalf("Failed to list user social accounts: %v", err)
		}
		if resp.NextPage == 0 || len(accounts) == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
}
