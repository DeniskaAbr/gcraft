package main

import (
	"log"

	"gcraft/app"
)

func main() {
	log.SetFlags(log.Lshortfile)

	instance := app.Create()

	if err := instance.Run(); err != nil {
		return
	}
}
