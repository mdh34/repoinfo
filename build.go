package main

import (
	"log"

	travis "github.com/Ableton/go-travis"
)

//LastBuild gets the status of the given repo and user's last travis build
func LastBuild(user string, repo string) string {
	client := travis.NewDefaultClient("")
	name := user + "/" + repo
	builds, _, _, _, err := client.Builds.ListFromRepository(name, nil)
	if err != nil {
		log.Fatal(err)
	}

	return builds[len(builds)-1].State
}
