package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		User    string `short:"u" long:"user" description:"the repo user" required:"false"`
		Repo    string `short:"r" long:"repo" description:"the repo name" required:"false"`
		Service string `short:"s" long:"service" description:"the service to check (github/gitlab)" required:"false"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	autoUser, autoRepo, autoService := GetRemoteDetails()
	if opts.User == "" {
		opts.User = autoUser
	} 
	if opts.Repo == "" {
		opts.Repo = autoRepo
	}
	if opts.Service == "" {
		opts.Service = autoService
	}

	var issues, pr int
	if opts.Service == "gitlab" {
		issues, pr = GetGitlabIssues(opts.User, opts.Repo)
	} else if opts.Service == "github" {
		issues, pr = GetGithubIssues(opts.User, opts.Repo)
	}

	last := LastBuild(opts.User, opts.Repo)
	if last == "passed" {
		color.Green("Last build: %v\n", last)
	} else {
		color.Red("Last build: %v\n", last)
	}
	fmt.Printf("%v issues open\n", issues)
	fmt.Printf("%v pull requests open\n", pr)
}
