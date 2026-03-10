package handler

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Huweicai/goto/alfred"
	"github.com/Huweicai/goto/config"
)

// ImportTextfiles imports all files from a directory into config.yaml under the given path prefix.
// Usage: goto $import-textfiles <path...> <dir>
// Example: goto $import-textfiles prompt xxx /Users/huweicai/workspace/ai/prompts
// This will scan the directory recursively and add each file as:
//
//	prompt xxx <relative-path-segments...> <filename-without-ext> -> textfile://<absolute-path>
func ImportTextfiles(args []string) *alfred.Output {
	if len(args) < 2 {
		log.Println("usage: $import-textfiles <path...> <directory>")
		return nil
	}

	dir := args[len(args)-1]
	pathPrefix := args[:len(args)-1]

	dir = expandHome(dir)

	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		log.Printf("'%s' is not a valid directory\n", dir)
		return nil
	}

	nest, err := config.NewNest(config.GetConfigPath())
	if err != nil {
		log.Fatalf("init nest failed: %s", err.Error())
		return nil
	}

	count := 0
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		// compute relative path from the scanned dir
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return nil
		}

		// split into path segments: dir1/dir2/file.md -> [dir1, dir2, file]
		parts := strings.Split(rel, string(filepath.Separator))
		// remove extension from the last segment (filename)
		last := parts[len(parts)-1]
		parts[len(parts)-1] = strings.TrimSuffix(last, filepath.Ext(last))

		// lowercase all segments
		for i := range parts {
			parts[i] = strings.ToLower(parts[i])
		}

		// full path in config: pathPrefix + relative segments
		fullPath := append(append([]string{}, pathPrefix...), parts...)
		value := textfileScheme + path

		nest.AddScalar(fullPath, value)
		count++
		return nil
	})

	if err != nil {
		log.Printf("walk directory failed: %v\n", err)
		return nil
	}

	if err := nest.Flush(); err != nil {
		log.Printf("flush failed: %v\n", err)
		return nil
	}

	log.Printf("imported %d files from %s\n", count, dir)
	return nil
}
