// localstatus is a tool for automatically monitoring your local dev environment.

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
