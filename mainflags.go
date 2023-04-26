package main

import "flag"

var isWatch = flag.Bool("watch", false, "run the checks every 3 minutes until canceled")

func mainFlags() {
	flag.Parse()
}
