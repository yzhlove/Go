package main

//go 操作postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

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
	// queryQuestion(db)
	// insertQuestion(db)
	// queryOneQuestion(db)
	// updateQuestion(db)
	// deleteQuestion(db)
	insertReturnQuestion(db)
	db.Close()

}

func queryQuestion(db *sql.DB) {

	rows, err := db.Query(`SELECT id,title,options,answer FROM questions`)
	if err != nil {
		panic(err)
	}
	// clos, _ := rows.Columns()
	// fmt.Printf("%v \n", clos)
	question := Question{}
	// rows.Next()
	// rows.Scan(&question.ID, &question.Title, &question.Options, &question.Answer)
	// fmt.Printf("question:%v \n", question)
	for rows.Next() {
		rows.Scan(&question.ID, &question.Title, &question.Options, &question.Answer)
		fmt.Printf("question:%v \n", question)
	}
}

func queryOneQuestion(db *sql.DB) {

	querySQL := `select id ,title,options,answer from questions where id = $1`
	row := db.QueryRow(querySQL, 18613)
	qs := Question{}
	err := row.Scan(&qs.ID, &qs.Title, &qs.Options, &qs.Answer)
	if err != nil {
		fmt.Println("query Err!")
	}
	fmt.Printf("rowQuery = %v \n", qs)

}

func insertQuestion(db *sql.DB) {

	insertSQL := `INSERT INTO questions(title,options,answer,status,created_at,updated_at,rank) VALUES($1,$2,$3,$4,$5,$6,$7)`
	stmt, _ := db.Prepare(insertSQL)
	// _, err := stmt.Exec("我是谁?", "[\"成龙\",\"李连杰\",\"甄子丹\",\"吴京\"]", "A", 1, time.Now().Unix(), time.Now().Unix(), 1)
	_, err := stmt.Exec("木兰辞出自", `["东汉","西汉","南宋","北宋"]`, "A", "1", time.Now().Unix(), time.Now().Unix(), "1")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Insert Successful!")
	}
}

func insertReturnQuestion(db *sql.DB) {

	var lastID int
	insertSQL := `INSERT INTO questions(title,options,answer,status,created_at,updated_at,rank) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	err := db.QueryRow(insertSQL, "中国四大发明是", `["西游记","水浒传","三国演义","红楼梦"]`, "A", "1", time.Now().Unix(), time.Now().Unix(), "1").Scan(&lastID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Last Insert Id = %v \n", lastID)
}

func updateQuestion(db *sql.DB) {

	updateSQL := `UPDATE questions SET title = $1 WHERE id = $2`
	stmt, err := db.Prepare(updateSQL)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec("汉乐府木兰辞出自?", 18613)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update Successful!")
}

func deleteQuestion(db *sql.DB) {
	deleteSQL := `DELETE FROM questions WHERE id = $1`
	stmt, err := db.Prepare(deleteSQL)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(18612)
	if err != nil {
		panic(err)
	}
	fmt.Println("Delete Successful!")
}
