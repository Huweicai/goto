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
	//try to open it
	if err = cmd.Run(); err != nil {
		//not url
		//print for copy to clipboard
		fmt.Print(value)
		return
		log.Println(err.Error())
	}
}
