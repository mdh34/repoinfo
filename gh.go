package main

import (
	"context"
	"log"

	"github.com/google/go-github/github"
)

//GetIssues gets all open issues for a given user and repo and returns the total
func GetIssues(user string, repo string) (int, int) {
	client := github.NewClient(nil)
	opt := new(github.IssueListByRepoOptions)

	var allIssues []*github.Issue
	for {
		issues, resp, err := client.Issues.ListByRepo(context.Background(), user, repo, opt)
		if err != nil {
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
