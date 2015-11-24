package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/knakayama/ghdump/dump"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghdump"
	app.Usage = "Dump your repository or starred repository in your github"
	app.HideHelp = true
	app.Commands = []cli.Command{
		{
			Name:  "repo",
			Usage: "Dump your repository",
			Action: func(c *cli.Context) {
				dump.DumpRepository()
			},
		},
		{
			Name:  "star",
			Usage: "Dump your starred repository",
			Action: func(c *cli.Context) {
				dump.DumpStarredRepository()
			},
		},
	}
	app.Run(os.Args)
}
