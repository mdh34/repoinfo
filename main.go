package main

import (
	"os"
	"sync"

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

	var wg sync.WaitGroup
	wg.Add(2)

	go LastBuild(opts.User, opts.Repo, &wg)

	if opts.Service == "gitlab" {
		go GetGitlabIssues(opts.User, opts.Repo, &wg)
	} else if opts.Service == "github" {
		go GetGithubIssues(opts.User, opts.Repo, &wg)
	}

	wg.Wait()
}
