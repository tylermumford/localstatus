package checks

import (
	"net/http"
	"strconv"
)

type BrockHttpOk struct {
}

var _ Check = BrockHttpOk{}

// Sends an HTTP GET request and expects 200 OK.
// Options: url, label
func (b BrockHttpOk) Run(options map[string]any) CheckResult {
	response, err := http.DefaultClient.Get(options["url"].(string))
	if err != nil {
		return newBasicResult(false, "error: "+err.Error())
	}

	if response.StatusCode != 200 {
		return newBasicResult(false, "not ok: "+strconv.Itoa(response.StatusCode))
	}

	return newBasicResult(true, options["url"].(string))
}
