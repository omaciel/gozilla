package main

import (
	"strings"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/omaciel/gozilla/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoZilla"
	app.Usage = "An opinionated command line for Bugzilla written in Go."

	app.Commands = []*cli.Command{
		{
			Name:  "bug",
			Usage: "Fetch a Bugzilla issue by its ID.",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     "id",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				ids := strings.Split(c.StringSlice("id")[0], ",")
				for _, id := range ids{
					resp := commands.GetBug(id)
					for _, bug := range resp.Bugs {
						data, _ := json.MarshalIndent(bug, "", "\t")
						fmt.Println(string(data))
					}
				}
				return nil
			},
		},
		{
			Name:  "history",
			Usage: "Fetch an issue's history.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				resp := commands.GetBugHistory(c.String("id"))
				for _, bug := range resp.Bugs {
					data, _ := json.MarshalIndent(bug, "", "\t")
					fmt.Println(string(data))
				}

				return nil
			},
		},
		{
			Name:  "version",
			Usage: "Shows the version for the Bugzilla server.",
			Action: func(c *cli.Context) error {
				resp := commands.Version()
				fmt.Println(resp)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
