package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

/*
delete add --start 500 --end 100
delete add -uid 123 --start 100 --end 100
delete detail
delete commit
delete remove 3
*/

func main() {

	app := cli.NewApp()
	app.Name = "delete"
	app.Usage = "delete module"
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"add"},
			Usage:   "add user",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start",
					Value: 0,
					Usage: "start Index",
				},
				cli.IntFlag{
					Name:  "end",
					Value: 0,
					Usage: "end Index",
				},
				cli.IntFlag{
					Name:  "uid",
					Usage: "user id",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println(c.Int("start"))
				fmt.Println(c.Int("end"))
				fmt.Println(c.Int("uid"))
				return nil
			},
		},
		{
			Name:    "detail",
			Aliases: []string{"detail"},
			Usage:   "detail module",
			Action: func(c *cli.Context) error {
				fmt.Println("detail ....")
				return nil
			},
		},
		{
			Name:    "commit",
			Aliases: []string{"commit"},
			Usage:   "commit module",
			Action: func(c *cli.Context) error {
				fmt.Println("commit ....")
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"remove"},
			Usage:   "remove module",
			Action: func(c *cli.Context) error {
				fmt.Println("remove ....")
				return nil
			},
		},
	}
	_ = app.Run(os.Args)
}
