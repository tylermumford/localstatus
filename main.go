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
There are no command line options at this time.
*/
package main

import (
	"fmt"
	"os"

	"github.com/tylermumford/localstatus/app"
)

func main() {
	err := app.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
