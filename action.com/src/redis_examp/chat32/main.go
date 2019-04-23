package main

import (
	"fmt"
	"log"
	"math/rand"
	"redis_examp/chat32/conn"
	"redis_examp/chat32/pool"
	"sync"
	"time"
)

const (
	count  = 20
	number = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := new(sync.WaitGroup)
	wg.Add(count)
	p, err := pool.New(conn.New, number)
	if err != nil {
		log.Fatal("[ERROR] pool is exit .")
	}
	defer p.Close()
	for i := 0; i < count; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		go func(id int) {
			defer wg.Done()
			backup(id, p)
		}(i)
	}
	wg.Wait()
	fmt.Println("Done .")
}

func backup(id int, p *pool.Pool) {
	c, err := p.Get()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Back(c)
	log.Printf("[INFO] %v Start .", c.(*conn.Conn).ID)
	time.Sleep(time.Duration(rand.Intn(10)+5) * time.Second)
	log.Printf("[INFO] %v Done .", c.(*conn.Conn).ID)
}
