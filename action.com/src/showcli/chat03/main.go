package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	os.Args = []string{"yzh", "love", "xjj"}
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q", c.Args().Get(0))
		return nil
	}
	_ = app.Run(os.Args)
}
