package main

import (
	"net/http"
	"sync"
)

type Task struct{}

var dataStore []Task
var dataStoreMutex sync.RWMutex
var oldestNotFinishedTask int
var oNFTMutex sync.RWMutex

func main() {
	dataStore = make([]Task, 0)
	oldestNotFinishedTask = 0

	http.HandleFunc("/getByTd", getById)
	http.HandleFunc("/newTask", newTask)
	http.HandleFunc("/getNewTask", getNewTask)
	http.HandleFunc("/finishTask", finishTask)
	http.HandleFunc("/setById", setById)
	http.HandleFunc("/list", listView)

	_ = http.ListenAndServe(":2345", nil)
}

func getById(w http.ResponseWriter, r *http.Request) {

}

func newTask(w http.ResponseWriter, r *http.Request) {

}

func getNewTask(w http.ResponseWriter, r *http.Request) {

}

func finishTask(w http.ResponseWriter, r *http.Request) {

}

func setById(w http.ResponseWriter, r *http.Request) {

}

func listView(w http.ResponseWriter, r *http.Request) {

}
