package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

//ExampleSendJson 提供了基础示例
func ExampleSendJson() {

	r, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)
	var u struct {
		Name  string
		Email string
	}
	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("ERROR:", err)
	}
	fmt.Println(u)
	//Output:
	//{lcm lcmm5201314@gmail.com}
}
