package handler

import (
	"github.com/Huweicai/goto/alfred"
	"github.com/Huweicai/goto/config"
	"log"
	"strings"
)

func List(args []string) {
	nest, err := config.NewNest(config.GetConfigPath())
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	outKV := nest.ListWithPre(args)
	output := alfred.NewOutput()
	for k, v := range outKV {
		arg := strings.Join(append(args[:len(args)-1], k), " ")
		//add a space for auto complete convenient
		output.AddSimpleTip(k, v, arg, arg+" ")
	}
	output.Show()
}
