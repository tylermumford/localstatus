package checks

import (
	"net"
	"time"
)

/*
check = "tcp.open"

Passes if a TCP connection can be opened.
Useful for checking many types of services,
such as databases and caches.
There is a 9 second timeout.

address: A string containing the host:port combo to connect to.
label: A string with a descriptive label. Optional.
*/
type CheckTcpOpen struct{}

func (b CheckTcpOpen) Run(options Params) CheckResult {
	addr := options.GetString("address")
	label := options.GetString("label")
	var subject string
	if label != "" {
		subject = label + ": " + addr
	} else {
		subject = addr
	}
	timeout := 9 * time.Second

	tcp, err := net.DialTimeout("tcp", addr, timeout)
	suffix := ""
	if err != nil {
		suffix += " " + err.Error()
	} else {
		defer tcp.Close()
	}

	return newBasicResult(err == nil, subject+suffix)
}
