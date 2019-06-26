package main

import (
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

var testData = filepath.Join(".", "testdata")
var testFiles = []struct {
	path     string
	expected []string
}{
	{
		path:     filepath.Join(testData, "simple_imports", "empty.go"),
		expected: []string{},
	},
	{
		path:     filepath.Join(testData, "simple_imports", "os_fmt.go"),
		expected: []string{"os", "fmt"},
	},
	{
		path: filepath.Join(testData, "go-loud", "loudbot.go"),
		expected: []string{
			"fmt",
			"log",
			"os",
			"regexp",
			"strings",
			"unicode",
			"github.com/go-redis/redis",
			"github.com/grokify/html-strip-tags-go",
			"github.com/joho/godotenv",
			"github.com/nlopes/slack",
		},
	},
}

func TestGetImports(t *testing.T) {
	for _, tf := range testFiles {
		t.Run(tf.path, func(t *testing.T) {
			imports, err := GetImports(tf.path)
			if err != nil {
				t.Fatal("Problem:", err)
			}

			sort.Strings(tf.expected)
			sort.Strings(imports)

			if eql := reflect.DeepEqual(tf.expected, imports); !eql {
				t.Fatalf("wanted %+v, got %+v\n", tf.expected, imports)
			}
		})
	}
}

var testPaths = []struct {
	path     string
	expected []string
}{
	{
		path:     filepath.Join(testData, "simple_imports"),
		expected: []string{"empty.go", "os_fmt.go"},
	},
	{
		path:     filepath.Join(testData, "go-loud"),
		expected: []string{"loudbot.go", "cmd/savelouds/savelouds.go", "cmd/seedlouds/seedlouds.go"},
	},
}

func TestGetGoFiles(t *testing.T) {
	for _, tf := range testPaths {
		t.Run(tf.path, func(t *testing.T) {
			imports, err := GetGoFiles(tf.path)
			if err != nil {
				t.Fatal("Problem:", err)
			}

			sort.Strings(tf.expected)
			sort.Strings(imports)

			if eql := reflect.DeepEqual(tf.expected, imports); !eql {
				t.Fatalf("wanted %+v, got %+v\n", tf.expected, imports)
			}
		})
	}
}
