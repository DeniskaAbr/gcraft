package main

import (
	"gcraft/app"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	instance := app.Create()

	if err := instance.Run(); err != nil {
		return
	}
}
