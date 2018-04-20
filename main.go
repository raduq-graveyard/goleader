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

type Step struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Recipe struct {
	Steps []Step `yaml:"steps"`
}

//FIXME - code smell, melhorar algum dia
func defaultAction(c *cli.Context) {
	var filename = "./recipes/sample.yml" //TODO - receive by cli

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	var r Recipe
	err = yaml.Unmarshal(yamlFile, &r)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	for _, step := range r.Steps {
		if strings.EqualFold(step.Name, "log") {
			var logger = worker.NewLog()
			logger.Log(step.Value)
		} else {
			log.Fatalf("Command '%s' not implemented", step.Name)
		}
	}
}
