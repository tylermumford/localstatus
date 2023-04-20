package checks

import (
	"net/http"
	"strconv"
)

type CheckHttpOk struct{}

var _ Check = CheckHttpOk{} // just confirming it implements the interface

// Sends an HTTP GET request and expects 200 OK.
// Options: url, label
func (b CheckHttpOk) Run(options Params) CheckResult {
	url := options.GetString("url")
	response, err := http.DefaultClient.Get(url)
	if err != nil {
		return newBasicResult(false, err.Error()) // [1]
	}

	if response.StatusCode != 200 {
		return newBasicResult(false, "not ok: "+strconv.Itoa(response.StatusCode))
	}

	return newBasicResult(true, url)
}

// [1] Normally, I'd add more information about the error, but
// Go errors already include a ton of information. If anything,
// it should be shortened.
