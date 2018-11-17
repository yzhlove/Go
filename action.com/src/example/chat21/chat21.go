package main

//go操作postgres数据库

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "jump_development"
)


func main() {

	plsqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	db , err := sql.Open("postgres", plsqlInfo)
	if err != nil {
		fmt.Println("Link ERR!")
		return
	}

	defer db.Close()

	//Ping
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("Link SUCCESSFUL!")

	sql := `SELECT * FROM game_users WHERE id = 1;`
	row, _ := db.Query(sql)

	fmt.Printf("%T ,%v \n",row,row)

}
