package main

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

var (
	keyValueStore map[string]string
	keyValueMutex sync.RWMutex
)

func main() {
	keyValueStore = make(map[string]string)
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	http.HandleFunc("/remove", remove)
	http.HandleFunc("/list", list)
	_ = http.ListenAndServe(":1234", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, "Error:", err)
			return
		}
		if len(values.Get("key")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, "Error:", err)
			return
		}
		keyValueMutex.RLock()
		value := keyValueStore[string(values.Get("key"))]
		keyValueMutex.RUnlock()
		_, _ = fmt.Fprint(w, value)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, "Error:Only GET accept.")
	}
}

func set(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, "Error:", err)
			return
		}
		if len(values.Get("key")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, "Error:", "Wrong input key.")
			return
		}
		if len(values.Get("value")) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, "Error:", "Wrong input value.")
			return
		}
		keyValueMutex.Lock()
		keyValueStore[string(values.Get("key"))] = string(values.Get("value"))
		keyValueMutex.Unlock()
		_, _ = fmt.Fprint(w, "success")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, "Error:Only POST accept.")
	}
}

func remove(w http.ResponseWriter, r *http.Request) {

}

func list(w http.ResponseWriter, r *http.Request) {

}
