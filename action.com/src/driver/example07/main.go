package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
	"time"
)

//测试数据库操作

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "jump_development"
)


//NewPostgresClient 新建一个postgres客户端
func NewPostgresClient() *sql.DB {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	fmt.Println("conn = ", conn)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Postgres Connection Successful!")
	return db
}


func main() {

	db := NewPostgresClient()

	data := make(map[string]interface{})

	data["title"] = "谁是卧底?"
	data["options"] = `["徐克","袁和平","王家卫","李安"]`
	data["answer"] = "A"
	data["status"] = 1
	data["created_at"] = time.Now().Unix()
	data["updated_at"] = time.Now().Unix()
	data["rank"] = 2

	InsertOpt(data,"questions",db)

}

func InsertOpt(data map[string]interface{},project string ,db *sql.DB) {

	if  data == nil ||  len(data) == 0 {
		fmt.Printf("------------")
		return
	}

	insertSQL := "INSERT INTO " + project + " ("
	length := len(data)
	values := make([]interface{},0 , length)
	intoArr := make([]string,0,length)
	for key ,value := range data {
		intoArr = append(intoArr,key)
		values = append(values,value)
	}
	insertSQL += strings.Join(intoArr,",")
	insertSQL += ") VALUES("
	for i := 0;i < length;i++ {
		intoArr[i] = "$" + strconv.Itoa(i+1)
	}
	insertSQL += strings.Join(intoArr,",")
	insertSQL += ")"

	fmt.Printf("insertSQL = %v \n",insertSQL)
	fmt.Printf("Values = %v \n",values)

	stmt , _ := db.Prepare(insertSQL)
	result,err := stmt.Exec(values...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result = %v \n",result)
	fmt.Printf("Insert Successful")

}