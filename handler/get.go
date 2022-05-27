package handler

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"

	"github.com/Huweicai/goto/alfred"
	"github.com/Huweicai/goto/config"
)

func Get(args []string) *alfred.Output {
	nest, err := config.NewNest("./config.yaml")
	if err != nil {
		log.Fatalf("init nest failed err:%s", err.Error())
		return nil
	}
	scalar, ok := nest.GetScalar(args)
	if !ok {
		log.Printf("%+v not found\n", args)
		return nil
	}

	nest.IncScalar(args)
	_ = nest.Flush()

	u, _ := url.Parse(scalar.Val)
	if u.Scheme == "" {
		// not url
		// print for copy to clipboard
		fmt.Print(scalar.Val)
		return nil
	}

	cmd := exec.Command("open", scalar.Val)
	// try to open it
	if err = cmd.Run(); err != nil {
		log.Println(err.Error())
	}
	return nil
}
