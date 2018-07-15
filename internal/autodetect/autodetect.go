package autodetect

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
	userName := url[strings.Index(url, "/"):strings.LastIndex(url, "/")]
	userName = strings.Trim(userName, "/")

	return userName, repoName

}
