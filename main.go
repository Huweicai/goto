package main

import (
	"flag"
	"log"

	"github.com/Huweicai/goto/alfred"
	"github.com/Huweicai/goto/handler"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 1 {
		panic("too few arguments")
	}
	alfred.SHOW(handler.GetHandler(args[0])(args[1:]))
}
