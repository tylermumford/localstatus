// brock is a tool for automatically monitoring your local dev environment.

package main

import (
	"fmt"

	"github.com/tylermumford/friendly-broccoli/brock"
)

func main() {
	fmt.Println("Hello!")

	brock.Run()
}
