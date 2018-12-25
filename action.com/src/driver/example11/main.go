package main

import (
	"driver/example11/drive"
	"fmt"
)

//操作数据库

func main() {

	postgres, ok := drive.NewPostgresClient()
	if !ok {
		fmt.Printf("Create Postgres Connection Err")
		return
	}

	//插入
	/*

		dataMap := make(map[string]interface{})
		dataMap["uid"] = 12345
		dataMap["game_code"] = "knife"
		dataMap["code"] = "yuewan"
		dataMap["action"] = "fishing"
		dataMap["type"] = 2
		dataMap["status"] = 1
		dataMap["output"] = 12
		dataMap["income"] = 24
		dataMap["created_at"] = time.Now().Unix()

		if _, ok := postgres.InsertSQL("fish_recording", dataMap); !ok {
			fmt.Printf("Insert Error")
		} else {
			fmt.Printf("Insert Successful")
		}

	*/

	//查询数量

	count := postgres.GetCount(drive.QueryModel{
		TableName: "fish_recording",
	})

	fmt.Printf("count = %v \n", count)

	//查询结果

	data, ok := postgres.FindPage(&drive.QueryModel{
		TableName:  "fish_recording",
		TableFiled: "*",
		Order:      "id desc",
	})
	if !ok {
		fmt.Printf("Query Err")
		return
	}
	fmt.Printf("TotalPage %v TotalEntries %v \n", data.TotalPage, data.TotalEntries)
	for k, v := range data.Result {
		fmt.Printf("Index %v Info: %v \n", k, v)
	}

}
