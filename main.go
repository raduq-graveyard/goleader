package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/raduq/goleader/worker"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "goleader"
	app.Usage = "Run your own Recipes"
	app.Action = defaultAction
	app.Version = "1.0.robot"
	app.Run(os.Args)
}

//FIXME - code smell, melhorar algum dia
func defaultAction(c *cli.Context) {
	var filename = "./recipes/sample.yml" //TODO - receive by cli

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var prop map[string]string
	err = yaml.Unmarshal(yamlFile, &prop)
	if err != nil {
		panic(err)
	}
	for key, value := range prop {
		if strings.EqualFold(key, "log") {
			var logger = worker.NewLog()
			logger.Log(value)
		} else {
			log.Fatalf("Command '%s' not implemented", key)
		}
	}
}
