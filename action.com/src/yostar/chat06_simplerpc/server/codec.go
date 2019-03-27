package server

import (
	"bytes"
	"encoding/gob"
)

//gob编解码

type RpcData struct {
	Name string
	Args []interface{}
}

func encode(data RpcData) ([]byte, error) {
	var (
		buf bytes.Buffer
		err error
	)
	bufEncode := gob.NewEncoder(&buf)
	if err = bufEncode.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(b []byte) (RpcData, error) {
	var (
		data RpcData
		err  error
	)
	buf := bytes.NewBuffer(b)
	bufDecode := gob.NewDecoder(buf)
	if err = bufDecode.Decode(&data); err != nil {
		return RpcData{}, err
	}
	return data, nil
}
