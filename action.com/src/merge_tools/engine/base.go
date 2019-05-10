package engine

type Engine interface {
	Request(*Item)
	Response() *Item
	Run()
}
