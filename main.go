package main

import (
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
				&cli.StringFlag{Name: "id"},
		  	},
		  	Action: func(c *cli.Context) error {
			  	resp := commands.Bug(c.String("id"))
			  	fmt.Println(resp)
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