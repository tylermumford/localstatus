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
	initialized bool

	interval time.Duration // How long Block will block
	until    time.Time     // The next time Block will unblock
	unblock  chan bool     // The internal signal to unblock
}

// Blocks the caller until either
// a set amount of time elapses,
// or the user presses the Enter key.
func (w *Watch) Block() {
	if !w.initialized {
		w.initialize()
	}

	<-w.unblock
	w.until = time.Now().Add(w.interval)
}

func (w *Watch) initialize() {
	w.initialized = true
	w.unblock = make(chan bool)
	w.interval = 3 * time.Minute

	// Time-based unblock
	go func() {
		w.until = time.Now().Add(w.interval)

		for {
			time.Sleep(w.until.Sub(time.Now()))

			// (If the user pressed Enter,
			// w.until could be in the future again,
			// so check to make sure it's actually in the past.)
			isPast := w.until.Sub(time.Now()) <= 0
			if isPast {
				w.unblock <- true
			}
		}
	}()

	// Enter-based unblock
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for {
			s.Scan()
			w.unblock <- true
		}
	}()
}
