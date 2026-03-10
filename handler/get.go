package handler

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/Huweicai/goto/alfred"
	"github.com/Huweicai/goto/config"
)

const fileScheme = "file://"

func Get(args []string) *alfred.Output {
	// check for $command (e.g. "$add key1 key2 url")
	if h, remaining, ok := resolveBuiltinCmd(args); ok {
		return h(remaining)
	}

	nest, err := config.NewNest(config.GetConfigPath())
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

	// file:// scheme: read file content and print for clipboard
	if strings.HasPrefix(scalar.Val, fileScheme) {
		filePath := expandHome(strings.TrimPrefix(scalar.Val, fileScheme))
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("read file %s failed: %v\n", filePath, err)
			fmt.Print(scalar.Val)
			return nil
		}
		fmt.Print(string(content))
		return nil
	}

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

func expandHome(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return home + path[1:]
	}

	return path
}
