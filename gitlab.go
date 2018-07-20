package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

//GetGitlabIssues gets all open GitLab issues and MR's for a given user and repo and returns the total
func GetGitlabIssues(user string, repo string) (int, int) {
	path := user + "/" + repo
	git := gitlab.NewClient(nil, "")

	issuesOpt := &gitlab.ListProjectIssuesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
		State: gitlab.String("opened"),
	}

	var allIssues []*gitlab.Issue
	for {
		issues, resp, err := git.Issues.ListProjectIssues(path, issuesOpt)
		if err != nil {
			log.Fatal(err)
		}

		allIssues = append(allIssues, issues...)
		if resp.NextPage == 0 {
			break
		}

		issuesOpt.Page = resp.NextPage
	}

	mergeOpt := &gitlab.ListProjectMergeRequestsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
		State: gitlab.String("opened"),
	}

	var allMerge []*gitlab.MergeRequest
	for {
		merges, resp, err := git.MergeRequests.ListProjectMergeRequests(path, mergeOpt)
		if err != nil {
			log.Fatal(err)
		}

		allMerge = append(allMerge, merges...)
		if resp.NextPage == 0 {
			break
		}

		mergeOpt.Page = resp.NextPage
	}

	return len(allIssues), len(allMerge)
}
