package main

import (
	"context"
	"log"
	"strings"

	"github.com/google/go-github/github"
)

//GetGithubIssues gets all open GitHub issues and PR's for a given user and repo and returns the total
func GetGithubIssues(user string, repo string) (int, int) {
	client := github.NewClient(nil)
	opt := new(github.IssueListByRepoOptions)

	var allIssues []*github.Issue
	for {
		issues, resp, err := client.Issues.ListByRepo(context.Background(), user, repo, opt)
		if err != nil && strings.Contains(err.Error(), "404") {
			log.Fatal("No GitHub repo found")
		} else if err != nil {
			log.Fatal(err)
		}

		allIssues = append(allIssues, issues...)
		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}

	//Filter out PR's because they're also included
	var list []github.Issue
	var pr int
	for _, issue := range allIssues {
		if !issue.IsPullRequest() {
			list = append(list, *issue)
		} else {
			pr++
		}
	}
	return len(list), pr
}
