package handler

import (
	"strings"

	"github.com/Huweicai/goto/alfred"
)

const (
	add  = "add"
	get  = "get"
	list = "list"
	show = "show"
	aist = "aist"

	cmdPrefix = "$"
)

type handler func(args []string) *alfred.Output

var handlers = map[string]handler{
	add:  Add,
	get:  Get,
	list: List,
	show: Show,
	aist: Aist,
}

// builtinCmds maps $-prefixed commands to their handlers
var builtinCmds = map[string]handler{
	"add":              Add,
	"import-textfiles": ImportTextfiles,
}

func GetHandler(name string) handler {
	return handlers[name]
}

// resolveBuiltinCmd checks if the first arg is a $command and returns the handler and remaining args
func resolveBuiltinCmd(args []string) (handler, []string, bool) {
	if len(args) == 0 {
		return nil, nil, false
	}
	if !strings.HasPrefix(args[0], cmdPrefix) {
		return nil, nil, false
	}
	cmd := strings.TrimPrefix(args[0], cmdPrefix)
	h, ok := builtinCmds[cmd]
	if !ok {
		return nil, nil, false
	}
	return h, args[1:], true
}
