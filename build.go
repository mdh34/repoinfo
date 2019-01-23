package main

import (
	"sync"
	"log"
	"github.com/fatih/color"
	travis "github.com/Ableton/go-travis"
)

//LastBuild gets the status of the given repo and user's last travis build
func LastBuild(user string, repo string, wg *sync.WaitGroup) {
	client := travis.NewDefaultClient("")
	builds, _, _, _, err := client.Builds.ListFromRepository(user+"/"+repo, nil)
	if err != nil {
		log.Fatal(err)
	} else if len(builds) == 0 {
		color.Red("No builds")
		return;
	}

	lastbuild := builds[len(builds)-1].State

	if lastbuild == "passed" {
		color.Green("Last build: %v\n", lastbuild)
	} else {
		color.Red("Last build: %v\n", lastbuild)
	}
	wg.Done()
}
