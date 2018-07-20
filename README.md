# repoinfo
A small (WIP) command-line utility to check up on your repositories

## Features
- Travis build status
- Number of open GitHub issues
- Number of open GitHub PR's
- Autodetection of user and repo based on git remotes

## Usage
repoinfo will try to autodetect the username and repository of the cwd's git repo

Using the `/u` and `/r` flags will let you specify which user and repository to check

The `/s` flag specifies which service to check, GitHub or GitLab

For example, to check this repository you would run
```
repoinfo /u mdh34 /r repoinfo /s gitlab
```

## Installation
To install, run
```
go get gitlab.com/mdh34/repoinfo
```