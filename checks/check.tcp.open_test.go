package checks

import (
	"net/http"
	"testing"
)

func TestLabelOptional(t *testing.T) {
	// Create a server to ensure TCP can be opened
	temp := http.Server{Addr: "localhost:9871", Handler: http.DefaultServeMux}
	go temp.ListenAndServe()
	defer temp.Close()

	chk := CheckTcpOpen{}
	options := Params{
		"address": "localhost:9871",
		// no label
	}

	result := chk.Run(options)

	if !result.IsOkay() {
		t.Fail()
	}
}