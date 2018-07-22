package main

import (
	"log"

	travis "github.com/Ableton/go-travis"
)

//LastBuild gets the status of the given repo and user's last travis build
func LastBuild(user string, repo string) string {
	client := travis.NewDefaultClient("")
	builds, _, _, _, err := client.Builds.ListFromRepository(user+"/"+repo, nil)
	if err != nil {
		log.Fatal(err)
	} else if len(builds) == 0 {
		return "No builds"
	}

	return builds[len(builds)-1].State
}
