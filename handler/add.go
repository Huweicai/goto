package handler

import (
	"log"

	"github.com/Huweicai/goto/alfred"
	"github.com/Huweicai/goto/config"
)

func Add(args []string) *alfred.Output {
	nest, err := config.NewNest(config.GetConfigPath())
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	nest.AddScalar(args[:len(args)-1], args[len(args)-1])
	if err := nest.Flush(); err != nil {
		log.Printf("flush failed: %v", err)
		return nil
	}

	return nil
}
