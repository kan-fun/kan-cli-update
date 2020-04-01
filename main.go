package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var version string

func main() {
	app := &cli.App{
		Name:     "kan-update",
		Usage:    "🎈 Update Tool for Kan 🎈",
		HelpName: "kan-update",
		Action:   index,
		Version:  version,
	}
	app.UseShortOptionHandling = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
