package main

import (
	"log"
	"os"

	"github.com/geraldofada/uncut-timer/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                   "uncut-timer",
		Usage:                  "Create timers for any kinds of purpose",
		UseShortOptionHandling: true,
	}

	app.Commands = []*cli.Command{
		{
			Name:    "start",
			Usage:   "Starts a timer at this moment",
			Aliases: []string{"s"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "name",
					Usage:   "Give your timer a name",
					Aliases: []string{"n"},
					Value:   "",
				},
			},
			Action: func(c *cli.Context) error {
				err := cmd.CliStart(c.String("name"))

				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:    "list",
			Usage:   "Lists all the ongoing timers",
			Aliases: []string{"l"},
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "finished",
					Usage:   "Lists all the finished timers instead of the ongoing",
					Aliases: []string{"f"},
					Value:   false,
				},
				&cli.IntFlag{
					Name:    "id",
					Usage:   "List a single timer with its ID instead of all the timers",
					Aliases: []string{"i"},
					Value:   -1,
				},
			},
			Action: func(c *cli.Context) error {
				var err error

				if c.Bool("finished") {
					err = cmd.CliList(c.Int("id"), "finished")
				} else {
					err = cmd.CliList(c.Int("id"), "ongoing")
				}

				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:    "stop",
			Usage:   "Stops a timer and save its data",
			Aliases: []string{"x"},
		},
		{
			Name:    "remove",
			Usage:   "Removes an ongoing timer",
			Aliases: []string{"r"},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}
}
