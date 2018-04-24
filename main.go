package main

import (
	"log"

	"github.com/kevingates/exchange/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
