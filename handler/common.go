package handler

const (
	add  = "add"
	get  = "get"
	list = "list"
	show = "show"
)

var handlers = map[string]func(args []string){
	add:  Add,
	get:  Get,
	list: List,
	show: Show,
}

func GetHandler(name string) func(args []string) {
	return handlers[name]
}
