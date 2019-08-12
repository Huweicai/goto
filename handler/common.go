package handler

import "github.com/Huweicai/goto/alfred"

const (
	add  = "add"
	get  = "get"
	list = "list"
	show = "show"
	aist = "aist"
)

type handler func(args []string) *alfred.Output

var handlers = map[string]handler{
	add:  Add,
	get:  Get,
	list: List,
	show: Show,
	aist: Aist,
}

func GetHandler(name string) handler {
	return handlers[name]
}
