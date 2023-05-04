package checks

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

/*
check = "git.branch"

Passes if the current branch
has all of the commits of the remote base branch.
(For example, if a feature branch
has all of the commits of origin/main.)

Runs `git fetch` and `git log`.
Git must already be on PATH.

  - dir: A string containing the path to the Git directory.
  - base: A string containing the name of the base branch.
    (TODO: Defaults to "origin/main".)
*/
type CheckGitBranch struct{}

func (g CheckGitBranch) Run(p Params) CheckResult {
	dir := p.GetString("dir")
	dirEnd := filepath.Base(dir)
	base := p.GetString("base")

	_, err := os.Lstat(dir)
	if os.IsNotExist(err) {
		return newBasicResult(false, dir+" does not exist")
	}

	gitDir := filepath.Join(dir, ".git")
	_, err = os.Lstat(gitDir)
	if os.IsNotExist(err) {
		return newBasicResult(false, dir+" is not a Git repo")
	}

	fetch := exec.Command("git", "fetch")
	fetch.Dir = dir
	fetch.Start() // Intentonally ignoring any fetch errors
	fetch.Wait()

	compare := "HEAD.." + base
	cmd := exec.Command("git", "log", "--oneline", compare)
	cmd.Dir = dir

	comparison, err := cmd.Output()
	if err != nil {
		msg := fmt.Sprintf("running git in %s: %s", dir, err.Error())
		return newBasicResult(false, msg)
	}

	// TODO: It would be nice to count the commits.
	if len(comparison) > 0 {
		return newBasicResult(false, dirEnd+": New commits on "+base)
	}

	return newBasicResult(true, dirEnd+": Up to date with "+base)
}
