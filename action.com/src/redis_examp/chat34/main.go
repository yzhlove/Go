package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

//创建一个新的AOF文件

func main() {

	request := GetRequestPool()
	if request == nil {
		panic("request err")
	}
	response := GetResponsePool()
	if response == nil {
		panic("response err")
	}

	respConn := request.Get()
	defer respConn.Close()

	keys, err := redis.Strings(respConn.Do("KEYS", "*"))
	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		if _, err = respConn.Do("MIGRATE", "saved", "6379", key, "0", "1000", "COPY"); err != nil {
			fmt.Printf("[ERROR] key %v err %v \n", key, err)
		}
	}

	fmt.Println("Done .")

}

func GetRequestPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     2,
		IdleTimeout: time.Duration(5 * time.Second),
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(0),
				redis.DialConnectTimeout(time.Duration(5*time.Second)))
		},
	}
}

func GetResponsePool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     2,
		IdleTimeout: time.Duration(5 * time.Second),
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6380", redis.DialDatabase(0),
				redis.DialConnectTimeout(time.Duration(5*time.Second)))
		},
	}
}

func ping(c redis.Conn) (err error) {
	if err = c.Err(); err != nil {
		return fmt.Errorf("[ERROR] connection err %v \n", err)
	}
	for {
		if _, err = c.Do("PING"); err == nil {
			return nil
		}
		time.Sleep(2 * time.Second)
	}
}
