package checks

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

/*
check = "npm.install"

Passes if npm dependencies are installed.
This is implemented with the `npm ls` command,
so npm must be installed and on the PATH already.

  - package: A string with the path to a package.json file.
*/
type CheckNpmInstall struct{}

func (n CheckNpmInstall) Run(p Params) CheckResult {
	packageFile := p.GetString("package")

	_, err := os.Lstat(packageFile)
	if os.IsNotExist(err) {
		return newBasicResult(false, packageFile+" does not exist")
	}

	dir := filepath.Dir(packageFile)

	cmd := exec.Command("npm", "ls")
	cmd.Dir = dir
	err = cmd.Start()
	if err != nil {
		return newBasicResult(false, fmt.Sprintf("npm in %s: %s", dir, err))
	}

	err = cmd.Wait()
	if err != nil {
		return newBasicResult(false, packageFile+" dependencies not satisfied")
	}

	return newBasicResult(true, packageFile+" dependencies")
}
