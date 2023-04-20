package checks

import (
	"os"
	"strconv"
	"strings"
)

type BrockEnv struct{}

func (b BrockEnv) Run(p Params) CheckResult {
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
