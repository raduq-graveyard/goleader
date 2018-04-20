package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "goleader"
	app.Usage = "Run your own Recipes"
	app.Action = defaultAction
	app.Version = "1.0.robot"
	app.Run(os.Args)
}

func defaultAction(c *cli.Context) {
	log.Println("Hallo")
}
