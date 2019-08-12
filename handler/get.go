package handler

import (
	"fmt"
	"github.com/Huweicai/goto/alfred"
	"github.com/Huweicai/goto/config"
	"log"
	"os/exec"
)

func Get(args []string) *alfred.Output {
	nest, err := config.NewNest("./config.yaml")
	if err != nil {
		log.Fatalf("init nest failed err:%s", err.Error())
		return nil
	}
	value, ok := nest.GetScalar(args)
	if !ok {
		log.Println("%+v not found", args)
		return nil
	}
	cmd := exec.Command("open", value)
	//try to open it
	if err = cmd.Run(); err != nil {
		//not url
		//print for copy to clipboard
		fmt.Print(value)
		log.Println(err.Error())
	}
	return nil
}
