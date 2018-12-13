package mock

import (
	"fmt"
)

//Retriver mooc
type Retriver struct {
	Contents string
}

func (r *Retriver) String() string {
	return fmt.Sprintf("Retiver:{Contents=%s}", r.Contents)
}

//Post 方法
func (r *Retriver) Post(url string, from map[string]string) string {
	r.Contents = from["contents"]
	return "ok"
}

//Get 方法
func (r *Retriver) Get(url string) string {
	return r.Contents
}
