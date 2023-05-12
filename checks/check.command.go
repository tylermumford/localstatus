package checks

import (
	"fmt"
	"os/exec"
	"strings"
)

/*
check = "command"

Experimental.
(I may decide to run this through shells in the future.)

Runs a command directly (not through a shell).
Passes only if the command's exit code is zero.
Note that the output will look like there are no quotes,
but arguments are kept intact (quoted) when passed to the program.

  - program: A string with the name of the program to run.
    If you want to run something in the current directory,
    use ./ before the file name.
    If the program is on your PATH, it will be found by its name alone.
  - args: An array of strings containing any arguments to pass.
    Optional.
  - dir: TODO.
*/
type CheckCommand struct{}

func (c CheckCommand) Run(p Params) CheckResult {
	program := p.GetString("program")
	args := make([]string, 0)
	if len(p.GetStrings("args")) > 0 {
		args = p.GetStrings("args")
	}

	line := program
	if len(args) > 0 {
		line += " "
		line += strings.Join(args, " ")
	}

	cmd := exec.Command(program, args...)
	err := cmd.Start()
	if err != nil {
		return newBasicResult(false, fmt.Sprintf("%s: %s", line, err.Error()))
	}

	err = cmd.Wait()
	if err != nil {
		return newBasicResult(false, fmt.Sprintf("%s: %s", line, err.Error()))
	}

	return newBasicResult(true, line)
}
