package main

import (
	"driver/examp04/opt"
)

func main() {

	db := opt.NewClient()
	defer db.Close()

	// opt.ExampleString(db)
	// opt.ExampleHash(db)
	opt.ExampleZset(db)

}
