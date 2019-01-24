package main

import "fmt"

//多个类实现相同的接口

var (
	Index      = 0
	ServerList = map[int]string{}
)

type Docker interface {
	Create(remote string) error
	Run(id int) error
	Remove(id int) bool
}

type Service struct{}

func (s *Service) Create(remote string) error {
	Index++
	ServerList[Index] = remote
	fmt.Println("Docker[INFO]:create successful")
	return nil
}

func (s *Service) Run(id int) error {
	if ret, ok := ServerList[id]; !ok {
		fmt.Println("Docker[INFO]:run error")
	} else {
		fmt.Println("Docker[INFO]:run successful ", ret)
	}
	return nil
}

type Admin struct {
	Service
}

func (admin *Admin) Remove(id int) bool {
	if _, ok := ServerList[id]; !ok {
		fmt.Println("Admin[INFO]:not found ", id)
	} else {
		delete(ServerList, id)
		fmt.Println("Admin[INFO]:remove successful")
	}
	return true
}

func main() {

	admin := new(Admin)

	_ = admin.Create("http://git.365yf.com")
	_ = admin.Run(1)
	admin.Remove(1)

}
