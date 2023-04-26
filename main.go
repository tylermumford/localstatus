/*
Localstatus checks your local dev environment
to ensure it's set up properly.
It reads a config file
to determine what to check,
such as services and files.

It reads this file, if it exists:

	~/.config/localstatus.toml

As a basic example of that file's structure:

	checks = [
		{check = "http.ok", url = "localhost:8000"}
	]

To learn which checks are available,
and how to configure them,
see the documentation of the checks package.

If any checks fail,
the command's exit code will be nonzero.

To have the program run itself repeatedly
(every 3 minutes),
run the program with the `--watch` flag.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/inancgumus/screen"
	"github.com/tylermumford/localstatus/app"
)

func main() {
	mainFlags()

	if !*isWatch {
		err := app.Run()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else {
		for {
			screen.Clear()
			screen.MoveTopLeft()
			fmt.Println("Running every 3 minutes...")
			fmt.Println("(press Enter to re-run immmediately; press Ctrl-C to stop)")

			err := app.Run()
			if err != nil {
				fmt.Println(err.Error())
			}

			blockUntilNextLoop()
			// Yes, Ctrl-C is the only way to exit
		}
	}
}

func blockUntilNextLoop() {
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
