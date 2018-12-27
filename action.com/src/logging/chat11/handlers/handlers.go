package handlers

import (
	"encoding/json"
	"net/http"
)

//设置路由服务

func SendJson(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "lcm",
		Email: "lcmm5201314@gmail.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}

func Routes() {
	http.HandleFunc("/sendjson", SendJson)
}
