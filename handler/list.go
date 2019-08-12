package handler

import (
	"github.com/Huweicai/goto/alfred"
	"github.com/Huweicai/goto/config"
	"log"
	"strings"
)

func List(args []string) *alfred.Output {
	nest, err := config.NewNest(config.GetConfigPath())
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}
	outKV := nest.ListWithPre(args)
	output := alfred.NewOutput()
	for k, v := range outKV {
		arg := strings.Join(append(args[:len(args)-1], k), " ")
		//add a space for auto complete convenient
		output.AddSimpleTip(k, v, arg, arg+" ")
	}
	return output
}

/**
list for add
*/
func Aist(args []string) (out *alfred.Output) {
	defer func() {
		if r := recover(); r != nil || out == nil || out.Items == nil {
			out = &alfred.Output{Items: []*alfred.Item{
				alfred.NewSimpleItem("GADD", "input url to add it to goto", strings.Join(args, " "), ""),
			}}
		}
	}()
	return List(args)
}
