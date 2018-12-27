package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	Routes()
}

func TestSendJson(t *testing.T) {
	t.Log("Start Test Server Endpoint...")
	{

		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tCreate Request Err:", ballotX, err)
		}
		t.Log("\tCreate Request Successful!", checkMark)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tReceive Code Err:", ballotX, rw.Code)
		}
		t.Log("\tReceive Code is 200", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("JsonDecode Err.  ", ballotX)
		}
		t.Log("JsonDecode Successful. ", checkMark)
		t.Logf("UserInfo :%v %v\n", u, checkMark)
	}
}
