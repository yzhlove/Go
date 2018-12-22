package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"log"
	"sync"

	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
	postgresHost = "127.0.0.1"
	postgresPort = "5432"
	postgresUser = "postgres"
	postgresPassword = ""
	postgresDBName ="jump_development"
)

type PostgresOpt struct {
	Client *sql.DB
	Mutex sync.Mutex
}

type QueryModel struct {
	TableName string
	TableFiled string
	Conditions string
	Bind map[string]interface{}
	Group string
	Having string
	Order string
	Page int
	Total int
}

type QueryResult struct {
	Page int
	TotalPage int
	Total int
	TotalEntries int
	Result map[int]map[string]string
}

func NewPostgresClient() (*PostgresOpt , bool) {
	sqlUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", postgresHost, postgresPort, postgresUser, postgresDBName, postgresPassword)
	driver , err := sql.Open(driverName,sqlUrl)
	if err != nil {
		log.Fatalln(err)
		return nil , false
	}
	if err = driver.Ping();err != nil {
		return nil , false
	}
	return &PostgresOpt{Client:driver} ,true
}

func main() {

	driver , ok  := NewPostgresClient()
	if !ok {
		return
	}

	//查询数量
	count := driver.GetCount(QueryModel{
		TableName:"questions",
		TableFiled:"*",
	})
	fmt.Printf("count = %v \n" , count)

	//带条件查询
	count = driver.GetCount(QueryModel{
		TableName:"questions",
		TableFiled:"*",
		Conditions:"rank = :rank: and status = :status:",
		Bind: map[string]interface{}{"rank":1,"status":1,},
	})
	fmt.Printf("count = %v \n",count)

	//分页查询
	 if dataTable ,ok := driver.FindPage(&QueryModel{
	 	TableName:"questions",
	 	TableFiled:"*",
	 	Conditions:"rank = :rank:",
	 	Bind: map[string]interface{}{"rank":2},
		Order:"id desc",
		Page:13,
	 }); ok {
	 	fmt.Printf("page:%v total:%v allPage:%v allTotal:%v \n",
	 		dataTable.Page,dataTable.Total,dataTable.TotalPage,dataTable.TotalEntries)
		for k,v := range dataTable.Result {
			fmt.Printf("%v %v \n",k,v)
		}
	 }

}

func handleResults(rows *sql.Rows) map[int]map[string]string {
	//关闭rows
	defer func() {
		if err := rows.Close();err != nil {
			log.Fatalln(err)
		}
	}()

	sqlTypes , err := rows.Columns()
	if err != nil {
		return nil
	}
	length := len(sqlTypes)
	bytes := make([][]byte,length)
	scans := make([]interface{},length)
	for i:= 0;i < length;i++ {
		scans[i] = &bytes[i]
	}
	dataMap := make(map[int]map[string]string)
	index := 0
	for rows.Next() {
		if err := rows.Scan(scans...) ; err != nil {
			log.Fatalln(err)
			continue
		}
		tempMap := make(map[string]string)
		for i := 0;i < length;i++ {
			tempMap[sqlTypes[i]] = string(bytes[i])
		}
		dataMap[index] = tempMap
		index++
	}
	return dataMap
}

func getQueryString(model *QueryModel) (string,[]interface{},bool) {
	if model.TableName == "" || model.TableFiled == ""{
		return "",nil,false
	}
	querySQL := "SELECT " + model.TableFiled + " FROM " + model.TableName
	if model.Conditions != "" {
		querySQL += " WHERE " + model.Conditions
	}
	if model.Group != "" {
		querySQL += " GROUP BY " + model.Group
	}
	if model.Having != "" {
		querySQL += " HAVING " + model.Having
	}
	if model.Order != "" {
		querySQL += " ORDER BY " + model.Order
	}
	length := len(model.Bind)
	typeFields := make([]string,0,length)
	typeValues := make([]interface{},0,length)
	for k,v := range model.Bind {
		typeFields = append(typeFields,k)
		typeValues = append(typeValues,v)
	}
	fmt.Printf("BeginSQL = %v \n",querySQL)
	//replace
	for i := 0;i < length;i++ {
		querySQL = strings.Replace(querySQL,":"+typeFields[i]+":","$"+strconv.Itoa(i+1),-1)
	}
	fmt.Printf("EndSQL = %v \n",querySQL)
	return querySQL , typeValues , true
}


func(opt *PostgresOpt) GetCount(model QueryModel) int {
	model.TableFiled = "count(*)"
	model.Order = ""
	model.Group = ""
	model.Having = ""
	if sqlStr ,values , ok:= getQueryString(&model);ok {
		count := 0
		if err := opt.Client.QueryRow(sqlStr,values...).Scan(&count);err == nil {
			return count
		}
	}
	return 0
}

func (opt *PostgresOpt)FindPage(model *QueryModel) (*QueryResult , bool) {
	if model.Total == 0 {
		model.Total = 20
	}
	if model.Page == 0 {
		model.Page = 1
	}
	totalEntries := opt.GetCount(*model)
	fmt.Printf("totalEntries = %v \n",totalEntries)
	totalPage := int(totalEntries / model.Total)
	if totalEntries % model.Total != 0 {
		totalPage ++
	}
	sqlStr , values ,ok := getQueryString(model)
	//添加分页
	sqlStr += " LIMIT " + strconv.Itoa(model.Total) + " OFFSET " + strconv.Itoa((model.Page - 1) * model.Total)
	if !ok {
		return nil , false
	}
	if rows , err := opt.Client.Query(sqlStr,values...);err == nil {
		dataMap := handleResults(rows)
		if dataMap != nil {
			return &QueryResult{Page:model.Page, Total:len(dataMap),
				TotalEntries:totalEntries, TotalPage:totalPage, Result:dataMap,} ,true
		}
	}
	return nil , false
}

func (opt *PostgresOpt) InsertSQL(tableName string,data map[string]interface{}) (int , bool) {

	if data == nil || len(data) == 0{
		return 0 ,false
	}
	length := len(data)
	typeFields := make([]string,0,length)
	typeValues := make([]interface{},0,length)
	for  k ,v := range data {
		typeFields = append(typeFields,k)
		typeValues  =append(typeValues,v)
	}
	insertSQL := "INSERT INTO " + tableName + " (" + strings.Join(typeFields,",") + ") "
	for i := 0 ; i < length;i++ {
		typeFields[i] = "$"+strconv.Itoa(i+1)
	}
	insertSQL += " VALUES(" + strings.Join(typeFields,",") + ") RETURNING id"

	opt.Mutex.Lock()
	defer opt.Mutex.Unlock()
	var id int
	if err := opt.Client.QueryRow(insertSQL,typeValues...).Scan(&id);err == nil {
		return id, true
	}
	return 0 ,false
}




