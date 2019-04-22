package main

import (
	"fmt"
	"log"
	"redis_examp/chat31_transform/db"
	"sync"

	"github.com/garyburd/redigo/redis"
)

//临时迁移

func main() {
	_ = readSource()
}

func readSource() error {
	var (
		pool    *redis.Pool
		err     error
		index   int
		values  []interface{}
		keyList []string
		wg      = new(sync.WaitGroup)
	)
	if pool, err = db.GetRequestPool(); err != nil {
		return fmt.Errorf("[ERROR] request pool err: %v \n", err)
	}
	c := pool.Get()
	defer c.Close()
	for {
		if values, err = redis.Values(c.Do("SCAN", index)); err != nil {
			return fmt.Errorf("[ERROR] ")
		}
		index, _ = redis.Int(values[0], nil)
		keyList, _ = redis.Strings(values[1], nil)
		wg.Add(1)
		go migrate(pool, keyList, wg)
		if index == 0 {
			break
		}
	}
	wg.Wait()
	fmt.Println("Done .")
	return nil
}

func migrate(pool *redis.Pool, keys []string, wg *sync.WaitGroup) {
	defer wg.Done()
	c := pool.Get()
	defer c.Close()
	for _, key := range keys {
		if _, err := redis.String(c.Do("MIGRATE", "saved", "6379", key, 0, 1000, "COPY")); err != nil {
			log.Printf("[ERROR] migrate err:%v %v \n", key, err)
		}
	}
}
