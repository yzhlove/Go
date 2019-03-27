package server

//gob编解码

type RpcData struct {
	Name string
	Args []interface{}
}

func encode(data RpcData) ([]byte, error) {

}

func decode(b []byte) (RpcData, error) {

}
