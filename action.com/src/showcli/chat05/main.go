package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	var language string

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "input language >> ",
			Destination: &language,
		},
	}

	app.Action = func(c *cli.Context) error {
		for i := 0; i < c.NArg(); i++ {
			fmt.Println(c.Args()[i])
		}

		if language == "chinese" {
			fmt.Println("China ....")
		} else {
			fmt.Println("Other Fuck ...")
		}
		return nil
	}

	_ = app.Run(os.Args)
}
