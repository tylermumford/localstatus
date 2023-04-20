package checks

import (
	"net/http"
	"testing"
)

func TestBrockHttpOk(t *testing.T) {
	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Okay"))
	})
	temp := http.Server{Addr: "localhost:9871", Handler: http.DefaultServeMux}
	go temp.ListenAndServe()
	defer temp.Close()

	ok := BrockHttpOk{}
	result, err := ok.Run(map[string]any{"url": "http://localhost:9871/ping"})

	if err != nil {
		t.Fatal(err)
	}

	if result.IsOkay() == false {
		t.Fatalf("should have been true, but: %s", result.Label())
	}
}
