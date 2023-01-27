package main

import (
	"log"

	"gcraft/gc_app"
)

var (
	GitBranch = "local"
	GitCommit = "build"
)

func main() {
	log.SetFlags(log.Lshortfile)

	instance := gc_app.Create(GitBranch, GitCommit)

	if err := instance.Run(); err != nil {
		return
	}
}
