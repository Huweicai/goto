package handler

const (
	add = "add"
	get = "get"
)

var handlers = map[string]func(args []string){
	add: Add,
	get: Get,
}

func GetHandler(name string) func(args []string) {
	return handlers[name]
}
