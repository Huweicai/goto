package handler

import (
	"github.com/Huweicai/goto/config"
	"log"
)

func Add(args []string) {
	nest, err := config.NewNest(config.GetConfigPath())
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	nest.AddScalar(args[:len(args)-1], args[len(args)-1])
	nest.Flush()
}
