package handler

import (
	"fmt"
	"log"
	"strings"

	"github.com/Huweicai/goto/alfred"
	"github.com/Huweicai/goto/config"
)

func List(args []string) *alfred.Output {
	// if first arg starts with $, delegate to aist-style listing for that command
	if len(args) > 0 && strings.HasPrefix(args[0], cmdPrefix) {
		cmd := strings.TrimPrefix(args[0], cmdPrefix)
		if _, ok := builtinCmds[cmd]; ok {
			return Aist(args[1:])
		}
	}

	nest, err := config.NewNest(config.GetConfigPath())
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	keys := nest.ListWithPre(args)
	output := alfred.NewOutput()

	// show matching $commands when at root level
	if len(args) <= 1 {
		pre := ""
		if len(args) == 1 {
			pre = args[0]
		}
		for cmd := range builtinCmds {
			full := cmdPrefix + cmd
			if strings.HasPrefix(full, pre) {
				output.AddSimpleTip(full, "builtin: "+cmd, full, full+" ")
			}
		}
	}

	for _, key := range keys {
		arg := strings.Join(append(args[:len(args)-1], key.Key), " ")
		//add a space for auto complete convenient
		item := output.AddSimpleTip(key.Key, key.Val, arg, arg+" ")
		item.Rank = key.Frequency
		if key.Frequency != 0 {
			item.Subtitle = fmt.Sprintf("[%d] %s", key.Frequency, key.Val)
		}
	}

	return output
}

/*
*
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
