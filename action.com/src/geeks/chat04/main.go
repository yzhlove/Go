package main

import (
	"errors"
	"fmt"
	"time"
)

//object pool

func main() {
	index := 10
	pool := NewPool(index)

	for i := 0; i < index; i++ {
		if user, err := pool.GetUser(time.Second); err != nil {
			fmt.Printf("Get User Err:%v \n", err)
		} else {
			fmt.Println(user)
			if err = pool.ReleaseUser(user); err != nil {
				fmt.Printf("Release User Err:%v \n", err)
			}
		}
	}

}

type User struct {
	Name string
	Age  int
}

type UserPool struct {
	bufChan chan *User
}

func NewPool(cache int) *UserPool {
	userPool := &UserPool{}
	userPool.bufChan = make(chan *User, cache)
	for i := 0; i < cache; i++ {
		userPool.bufChan <- &User{Name: "none"}
	}
	return userPool
}

func (pool *UserPool) GetUser(timeout time.Duration) (*User, error) {
	select {
	case user := <-pool.bufChan:
		return user, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (pool *UserPool) ReleaseUser(user *User) error {
	select {
	case pool.bufChan <- user:
		return nil
	default:
		return errors.New("overflow")
	}
}
