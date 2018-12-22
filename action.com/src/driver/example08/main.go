package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

//数据库查询操作

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "jump_development"
)

//Question 题库
type Question struct {
	ID        int64  `db:"id"`
	Title     string `db:"title"`
	Options   string `db:"options"`
	Answer    string `db:"answer"`
	Status    int    `db:"status"`
	BeginTime int64  `db:"created_at"`
	EndTime   int64  `db:"updated_at"`
	Rank      int    `db:"rank"`
}


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
	//queryQuestion(db)

	//query := &Query{
	//	TableName:"questions",
	//	Conditions:"id > :id: and status = :status: and rank = :rank:",
	//	Bind: map[string]interface{}{"id":18600,"status":1,"rank":2},
	//	Order:"id desc",
	//}
	//
	//results :=  queryQuestionInfo(query,db)
	////fmt.Printf("results = %v \n",results)
	//
	//for key ,value := range results {
	//	fmt.Printf("index = %v value = %v \n",key,value)
	//}

	queryCount(db)

}


type Query struct {
	TableName string
	Conditions string
	Bind map[string]interface{}
	Group string
	Having string
	Order string
	Page int
	Count int
}


func queryQuestionInfo(query *Query, db *sql.DB) map[int]map[string]string {

	if  query.TableName == "" {
		return nil
	}
	querySQL := "SELECT * FROM " + query.TableName
	if query.Conditions != "" {
		querySQL += " WHERE " + query.Conditions
	}
	if query.Group != "" {
		querySQL += " GROUP BY " + query.Group
	}
	if query.Having != "" {
		querySQL += " HAVING " + query.Having
	}
	if query.Order != "" {
		querySQL += " ORDER BY " + query.Order
	}
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Count <= 0 {
		query.Count = 20
	}
	querySQL += " LIMIT " + strconv.Itoa(query.Count) + " OFFSET " + strconv.Itoa((query.Page - 1) * query.Count)
	fmt.Printf("querySQL = %v \n",querySQL)
	length := len(query.Bind)
	keys := make([]string,0,length)
	values := make([]interface{},0,length)

	for key ,value := range query.Bind {
		keys = append(keys,key)
		values = append(values,value)
	}
	fmt.Println()

	for i := 0; i < length;i++ {

		oldStr := ":" + keys[i] +":"
		fmt.Printf("key = %v old = %v \n",keys[i],oldStr)
		querySQL = strings.Replace(querySQL,oldStr,"$"+strconv.Itoa(i+1),-1)
	}
	fmt.Println()
	fmt.Printf("querySQL = %v \n",querySQL)

	fmt.Printf("queryCalue %v \n",values)

	rows , err := db.Query(querySQL,values...)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	return handlerResults(rows)
}


func handlerResults(rows *sql.Rows) map[int]map[string]string {
	colTypes , _ := rows.Columns()
	length := len(colTypes)
	bytes := make([][]byte,length)
	scans := make([]interface{},length)
	for i := 0;i< length ;i++ {
		scans[i] = &bytes[i]
	}
	results := make(map[int]map[string]string)
	index := 0
	for rows.Next() {
		if err := rows.Scan(scans...);err != nil {
			panic(err)
			return nil
		}
		row :=make(map[string]string)
		//遍历bytes数组，因为bytes->scans
		for k , v := range bytes {
			key := colTypes[k]
			row[key] = string(v)
		}
		results[index] = row
		index++
	}
	rows.Close()
	return results
}


func queryCount(db *sql.DB) {
	defer db.Close()
	countSQL := `SELECT COUNT(*) as nums from  questions `
	if rows , err := db.Query(countSQL);err == nil {
		colTypes , _ := rows.Columns()
		fmt.Printf("colTypes = %v \n",colTypes)
		data := handlerResults(rows)
		for _,v := range data {
			fmt.Printf("nums = %v \n",v)
		}
	}

}


func queryQuestion( db *sql.DB) {
	querySQL := `select * from questions where id = $1`
	rows , err:= db.Query(querySQL,18613)
	if err != nil {
		panic(err)
	}
	clos , _ :=  rows.Columns()
	fmt.Printf("clos %v \n",clos)

	//var id ,title,options,answer string
	//for rows.Next() {
	//	rows.Scan(&id,&title,&options,&answer)
	//	fmt.Println(id,title,options,answer)
	//}
}

func Find(db *sql.DB) {
	defer db.Close()
	querySQL := "select * from questions where rank = $1"
	rows , err := db.Query(querySQL,1)
	if err != nil {
		panic(err)
	}
	data := handlerResults(rows)
	for _ , v := range data {
		fmt.Printf("data = %v \n",v)
	}
}

func FindFirst(db *sql.DB) {

	queryOneSQL := "select * from questions where rank = $1 limit 1 offset 0"
	rows , err := db.Query(queryOneSQL,1)
	if err != nil {
		panic(err)
	}
	data := handlerResults(rows)
	for _ , v := range data {
		fmt.Printf(" data =  %v \n",v)
	}
}





