package checks

import (
	"bytes"
	"fmt"
	"os"
)

/*
check = "file.contains"

Passes if the given file contains
a literal string.

  - path: A string containing the path to check.
  - string: A string containing the character sequence to look for.
*/
type CheckFileContains struct{}

func (f CheckFileContains) Run(p Params) CheckResult {
	path := p.GetString("path")
	str  := p.GetString("string")

	info, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return newBasicResult(false, path+" does not exist")
	}

	if info.IsDir() {
		return newBasicResult(false, path+" is a directory (file.contains only checks files)")
	}

	good := contains(path, str)
	if !good {
		return newBasicResult(false, path+" does not contain \""+str+"\"")
	}

	msg := fmt.Sprintf("%s contains \"%s\"", path, str)
	return newBasicResult(true, msg)
}

func contains(path string, str string) bool {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	
	return bytes.Contains(file, []byte(str))
}

