// Command docs extracts documentation
// from comments in package checks.
package main

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"
)

//go:generate go run generate.go

func main() {
	fmt.Println("I'll generate some docs")

	deleteMarkdownFiles()

	// Parse the Go files in the checks package

	fset := &token.FileSet{}
	mode := parser.ParseComments
	pkgs, err := parser.ParseDir(fset, "../checks", nil, mode)

	if err != nil {
		fmt.Printf("parsing: %s", err)
		return
	}

	c := pkgs["checks"]
	computed := doc.New(c, "github.com/tylermumford/localstatus/checks", 0)

	// Prepare to grab the doc comments
	// of the types which are checks

	regex := regexp.MustCompile(`^check = "(\S+)"`)
	count := 0

	for _, t := range computed.Types {
		if !strings.HasPrefix(t.Doc, "check =") {
			continue
		}

		submatches := regex.FindStringSubmatch(t.Doc)
		name := submatches[1]

		// Write the doc comments to their own files

		const mode = 0666
		err := os.WriteFile(name+".md", []byte(t.Doc), mode)
		if err != nil {
			fmt.Printf("writing docs for "+name+": %s\n", err)
			continue
		}
		count += 1
	}

	fmt.Printf("generated %d docs\n", count)
}

func deleteMarkdownFiles() {
	files, _ := os.ReadDir(".")
	count := 0
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".md") {
			continue
		}
		os.Remove(f.Name())
		count += 1
	}
	fmt.Printf("(deleted %d existing docs)\n", count)
}

type docBit struct {
	name string
	docs string
}
