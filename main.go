// brock is a tool for automatically monitoring your local dev environment.

package main

import (
	"fmt"
	"os"

	"github.com/tylermumford/friendly-broccoli/brock"
)

func main() {
	err := brock.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
