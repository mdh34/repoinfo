package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		User string `short:"u" long:"user" description:"the repo user" required:"false"`
		Repo string `short:"r" long:"repo" description:"the repo name" required:"false"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	if opts.User == "" || opts.Repo == "" {
		opts.User, opts.Repo = GetRemoteDetails()
	}

	issues, pr := GetIssues(opts.User, opts.Repo)
	last := LastBuild(opts.User, opts.Repo)

	if last == "passed" {
		color.Green("Last build: %v\n", last)
	} else {
		color.Red("Last build: %v\n", last)
	}
	fmt.Printf("%v issues open\n", issues)
	fmt.Printf("%v pull requests open\n", pr)
}
