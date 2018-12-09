package proto

import (
	"github.com/davyxu/cellnet/util"

	"reflect"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
)

//ACK 返回消息头
type ACK struct {
	ErrorCode   int         `json:"error_code"`
	Number      string      `json:"room_number"`
	ErrorReason string      `json:"error_reason"`
	Result      interface{} `json:"result"`
}

func init() {
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*ACK)(nil)).Elem(),
		ID:    int(util.StringHash("proto.ACK")),
	})
}
