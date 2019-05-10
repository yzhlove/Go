package event

var loop = make(map[string][]func(interface{}))

//RegisterEvent 注册事件
func RegisterEvent(name string, callback func(interface{})) {
	list := loop[name]
	list = append(list, callback)
	loop[name] = list
}

//CallEvent 调用事件
func CallEvent(name string, param interface{}) {
	for _, callback := range loop[name] {
		callback(param)
	}
}
