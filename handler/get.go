package handler

import (
	"fmt"
	"github.com/Huweicai/goto/config"
	"log"
	"os/exec"
)

func Get(args []string) {
	nest, err := config.NewNest("./config.yaml")
	if err != nil {
		log.Fatalf("init nest failed err:%s", err.Error())
		return
	}
	value, ok := nest.GetScalar(args)
	if !ok {
		log.Println("%+v not found", args)
		return
	}
	cmd := exec.Command("open", value)
	//print for copy to clipboard
	fmt.Print(value)
	//try to open it
	err = cmd.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
