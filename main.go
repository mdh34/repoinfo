package main

import (
	"fmt"
	"os"

	build "repoinfo/internal/build"
	gh "repoinfo/internal/gh"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		User string `short:"u" long:"user" description:"the repo user" required:"true"`
		Repo string `short:"r" long:"repo" description:"the repo name" required:"true"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	issues, pr := gh.GetIssues(opts.User, opts.Repo)
	last := build.LastBuild(opts.User, opts.Repo)

	if last == "passed" {
		color.Green("Last build: %v\n", last)
	} else {
		color.Red("Last build: %v\n", last)
	}
	fmt.Printf("%v issues open\n", issues)
	fmt.Printf("%v pull requests open\n", pr)
}
