package main

//go:generate msgp -io=false -tests=false

type Foo struct {
	Bar string `msg:"bar"`
	Baz int32  `msg:"baz"`
}

func main() {

	//fo := Foo{
	//	Bar: "BigBigBar",
	//	Baz: 12345,
	//}

}
