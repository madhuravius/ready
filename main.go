package main

import (
	"log"
	"os"
	"ready/internal"
)

func main() {
	if err := internal.NewApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
