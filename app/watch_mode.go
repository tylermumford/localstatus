package app

import (
	"bufio"
	"os"
	"time"
)

// A Watch is used for its Block method,
// which blocks the caller until either
// a set amount of time elapses,
// or the user presses the Enter key.
type Watch struct {
	
}

// Blocks the caller until either
// a set amount of time elapses,
// or the user presses the Enter key.
func (w  Watch) Block() {
	unblock := make(chan bool)

	go func() {
		time.Sleep(3 * time.Minute)
		unblock <- true
	}()

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		unblock <- true
	}()

	<-unblock
}