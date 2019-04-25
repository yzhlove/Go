package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "greet"
	app.Version = "1.0.0"
	app.Description = "This is app..."
	app.Authors = []cli.Author{
		{Name: "yzh", Email: "lcmm5201314@gmail.com"},
		{Name: "xjj", Email: "xjjlove@qq.com"},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang,l",
			Value: "english",
			Usage: "input name >>",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:        "description",
			Aliases:     []string{"commit"},
			Usage:       "use input d",
			Description: "This is description ...",
			Action: func(c *cli.Context) error {
				for i := 0; i < c.NArg(); i++ {
					fmt.Println(os.Args[i])
				}
				fmt.Println("lang = ", c.String("lang"))
				fmt.Println("l = ", c.String("l"))
				return nil
			},
		},
	}

	_ = app.Run(os.Args)

}
