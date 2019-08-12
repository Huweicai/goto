package handler

import (
	"github.com/Huweicai/goto/alfred"
	"github.com/Huweicai/goto/config"
	"log"
)

func Add(args []string) *alfred.Output {
	nest, err := config.NewNest(config.GetConfigPath())
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}
	nest.AddScalar(args[:len(args)-1], args[len(args)-1])
	nest.Flush()
	return nil
}
