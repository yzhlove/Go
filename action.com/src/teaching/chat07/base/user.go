package base

type User interface {
	Info()
}

var (
	peopleFactory = make(map[string]func() User)
)

func Register(name string, factory func() User) {
	peopleFactory[name] = factory
}

func NewUser(name string) User {
	if factory, ok := peopleFactory[name]; ok {
		return factory()
	} else {
		panic("name not found")
	}
}
