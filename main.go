package main

import (
	"encoding/json"
	"flag"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Please supply a path to search")
		os.Exit(1)
	}

	importMap := make(map[string][]string)
	goFiles, err := GetGoFiles(flag.Arg(0))
	if err != nil {
		os.Exit(15)
	}

	for _, f := range goFiles {
		imports, err := GetImports(f)
		if err != nil {
			os.Exit(17)
		}
		importMap[f] = imports
	}

	if err != nil {
		log.Fatal("Problem searching for .go files", err)
		os.Exit(2)
	}

	var f io.Writer
	if outputPath := flag.Arg(1); outputPath != "" {
		f, err = os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0755)
	} else {
		f = os.Stdout
	}

	enc := json.NewEncoder(f)
	err = enc.Encode(importMap)
	if err != nil {
		log.Fatal("Problem encoding", err)
		os.Exit(3)
	}
}

func GetImports(path string) ([]string, error) {
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

func GetGoFiles(path string) ([]string, error) {
	var goFiles []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasSuffix(path, ".go") {
				goFiles = append(goFiles, path)
			}

			return nil
		})

	if err != nil {
		return []string{}, err
	}

	return goFiles, nil
}
