package main

import (
	"os"

	"github.com/bryo13/wand/cmd/wand"
)

func main() {
	cmdArgs := os.Args[1]
	if cmdArgs == "bin" {
		wand.Cram()
	}

}
