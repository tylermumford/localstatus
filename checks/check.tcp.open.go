package checks

import (
	"net"
	"time"
)

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
