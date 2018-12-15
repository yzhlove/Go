package main

import (
	"log"
	"os"
	"project/runner/monitor"
	"time"
)

//通过通道来监视程序
//程序运行时间，以便在程序运行时间过长的时候终止程序

const timeout = 5 * time.Second

func main() {

	log.Println("Start Work.")

	//分配一个runner
	runner := monitor.New(timeout)

	runner.Add(createTask(), createTask(), createTask())

	if err := runner.Start(); err != nil {
		switch err {
		case monitor.ErrorTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case monitor.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Process - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
