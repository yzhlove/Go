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
	count  = 50
	number = 5
)

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(count)
	p, err := pool.New(conn.New, number)
	if err != nil {
		log.Fatal("[ERROR] pool is exit .")
	}
	defer p.Close()
	for i := 0; i < count; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
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
	time.Sleep(5 * time.Second)
	log.Printf("[INFO] %v Done .", c.(*conn.Conn).ID)
}
