package server

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func queryUser(uid int) (User, error) {
	userDB := make(map[int]User)
	userDB[0] = User{Name: "yzh", Age: 22}
	userDB[1] = User{Name: "xjj", Age: 21}
	userDB[2] = User{Name: "xyj", Age: 20}
	userDB[3] = User{Name: "lcm", Age: 23}
	if u, ok := userDB[uid]; ok {
		return u, nil
	}
	return User{}, errors.New("not found")
}

func Test_SimpleRpc(t *testing.T) {
	gob.Register(User{})

	sev := NewRpcServer(address)
	sev.Register("queryUser", queryUser)
	go sev.Run()

	conn, err := net.Dial("tcp", address)
	if err != nil {
		t.Error(err)
	}

	cli := NewClient(conn)

	var query func(int) (User, error)
	cli.callRpc("queryUser", &query)
	u, err := query(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
}
