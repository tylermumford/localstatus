package checks

import (
	"fmt"
	"os"
)

/*
check = "file.exists"

Passes if the given file exists.
Works for both files and directories.

  - path: A string containing the path to check.
*/
type CheckFileExists struct{}

func (f CheckFileExists) Run(p Params) CheckResult {
	path := p.GetString("path")

	info, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return newBasicResult(false, path+" does not exist")
	}

	var note string
	if info.IsDir() {
		note = ""
	} else {
		note = fmt.Sprintf(" (%d bytes)", info.Size())
	}

	msg := fmt.Sprintf("%s exists%s", path, note)
	return newBasicResult(true, msg)
}
