package proto

import (
	"reflect"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"github.com/davyxu/cellnet/util"
)

//LOGIN 登陆协议
type LOGIN struct {
	SID           string `json:"sid"`
	CODE          string `json:"code"`
	GameHistoryID string `json:"game_history_id"`
	EnterFrom     string `json:"enter_from"`
	ADDR          string `json:"request_root"`
}

func init() {
	//注册登陆协议
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("json"),
		Type:  reflect.TypeOf((*LOGIN)(nil)).Elem(),
		ID:    int(util.StringHash("proto.LOGIN")),
	})

}
