package handler

import (
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
	url, ok := nest.GetScalar(args)
	if !ok {
		log.Println("%+v not found", args)
		return
	}
	cmd := exec.Command("open", url)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
