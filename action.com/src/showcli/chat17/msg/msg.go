package msg

//go:generate msgp -io=false -tests=false

type Message struct {
	ID  int
	Msg string
	Ts  int
}
