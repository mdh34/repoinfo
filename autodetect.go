package main

import (
	"log"
	"os"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

//GetRemoteDetails interprets the remote url and parses to guess a user and repository name
func GetRemoteDetails() (string, string) {
	cwd, err := os.Getwd()
	repo, err := git.PlainOpen(cwd)
	if err != nil {
		log.Fatal(err)
	}

	remote, err := repo.Remote("origin")
	if err != nil {
		log.Fatal(err)
	} else if remote == nil {
		return "", ""
	}

	url := remote.Config().URLs[0]
	url = strings.Replace(url, ".git", "", -1)
	url = strings.Replace(url, "https://", "", -1)

	repoName := url[strings.LastIndex(url, "/"):]
	repoName = strings.Trim(repoName, "/")
	var userName string
	if strings.Contains(url, "@") {
		userName = url[strings.Index(url, ":")+1 : strings.LastIndex(url, "/")]
	} else {
		userName = url[strings.Index(url, "/")+1 : strings.LastIndex(url, "/")]
	}

	return userName, repoName

}
