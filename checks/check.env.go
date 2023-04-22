package checks

import (
	"os"
	"strconv"
	"strings"
)

/*
check = "env"

Passes if all the given environment variables are set.
Variables set to an empty string DO count as being set.

variables_required: An array of strings,
each with an env var name.
*/
type CheckEnv struct{}

func (b CheckEnv) Run(p Params) CheckResult {
	required := p.GetStrings("variables_required")
	unsetVars := []string{}

	for i := range required {
		_, found := os.LookupEnv(required[i])
		if !found {
			unsetVars = append(unsetVars, required[i])
		}
	}

	if len(unsetVars) > 0 {
		return newBasicResult(false, "Not set: "+strings.Join(unsetVars, ", "))
	}

	return newBasicResult(true, strconv.Itoa(len(required))+" env vars")
}
