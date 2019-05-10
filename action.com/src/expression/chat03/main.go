package main

import (
	"expression/chat03/db"
	"expression/chat03/expre"
	_ "expression/chat03/match"
)

func main() {

	keys := db.GetKeys()

	expre.Run(keys...)

}
