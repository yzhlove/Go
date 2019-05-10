package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "English",
			Usage: "input language >>",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("--------------- [Args] ---------------")
		fmt.Println(c.NArg())
		for i := 0; i < c.NArg(); i++ {
			fmt.Println(c.Args().Get(i))
		}
		fmt.Println("--------------------------------------")

		if c.String("lang") == "chinese" {
			fmt.Println("China")
		} else {
			fmt.Println("Other")
		}
		return nil
	}
	_ = app.Run(os.Args)
}
