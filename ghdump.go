package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/knakayama/ghdump/credential"
	"github.com/knakayama/ghdump/dump"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghdump"
	app.Usage = "Dump your repository or starred repository in your github"
	app.HideHelp = true
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "Setup github credential",
			Action: func(c *cli.Context) {
				credential.SetCredential()
			},
		},
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
