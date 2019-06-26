package main

import (
	"encoding/json"
	"flag"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatal("Please supply a path to search and an output file")
		os.Exit(1)
	}

	importMap := make(map[string][]string)
	err := filepath.Walk(flag.Arg(0),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasSuffix(path, ".go") {
				imports, err := getImports(path)
				if err != nil {
					return err
				}
				importMap[path] = imports
			}

			return nil
		})

	if err != nil {
		log.Fatal("Problem searching for .go files", err)
		os.Exit(2)
	}

	outputPath := flag.Arg(1)
	f, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0755)
	enc := json.NewEncoder(f)
	err = enc.Encode(importMap)
	if err != nil {
		log.Fatal("Problem encoding", err)
		os.Exit(3)
	}

}

func getImports(path string) ([]string, error) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
	if err != nil {
		return []string{}, err
	}

	imports := []string{}
	for _, i := range f.Imports {
		imports = append(imports, strings.Trim(i.Path.Value, `"`))
	}

	return imports, nil
}
