package opt

import (
	"driver/examp04/ssdb"
	"fmt"
)

//ExampleString string相关操作
func ExampleString(db *ssdb.Client) {

	_, err := db.Set("ssdb_name", "go_opt_ssdb")
	if err != nil {
		panic(err)
	}
	nameString, err := db.Get("ssdb_name")
	if err != nil {
		panic(err)
	}
	fmt.Println("nameString = ", nameString)

	//getset
	result, err := db.Do("getset", "ssdb_name", "what_are_you_doing?")
	if err != nil {
		panic(err)
	}
	nameString, err = db.Get("ssdb_name")
	fmt.Printf(" result : = %v , nameString = %v \n", result, nameString)

	result, err = db.Do("getset", "ssdb_new_name", "How_Are_You")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result = %v \n", result)

	result, _ = db.Do("exists", "ssdb_name")
	fmt.Printf("exists result = %v \n", result)

	result, _ = db.Do("exists", "ssdn_name")
	fmt.Printf("exists result = %v \n", result)

	result, _ = db.Do("multi_set", "ssdb_name", "Hello_World", "ssdb_birthday", "1996-12-24")
	fmt.Printf("result = %v \n", result)

	result, _ = db.Do("multi_get", "ssdb_name", "ssdb_birthday")
	fmt.Printf("result = %v \n", result)

}
