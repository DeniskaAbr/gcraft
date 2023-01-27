package main

import (
	"log"

	"gcraft/app"
)

var (
	GitBranch = "local"
	GitCommit = "build"
)

func main() {
	log.SetFlags(log.Lshortfile)

	instance := app.Create(GitBranch, GitCommit)

	if err := instance.Run(); err != nil {
		return
	}
}
