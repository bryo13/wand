// assembly point before we get to main

package wand

import (
	"fmt"
	"os/exec"
)

func init() {
	if !isGoInstalled() {
		fmt.Println("please install Go from 'https://go.dev/doc/install'")
	}
}

func isGoInstalled() bool {
	// we can look for go in /usr/local directory but that limits us to linux
	// we can run go commands using exec to check if go is installed
	// the safer option is to use the LookPath function in the exec package

	path, err := exec.LookPath("go")
	if err != nil {
		return false
	}
	fmt.Println("Go lives in ", path)
	return true
}

// optional, used to zip the directory
func Cram() {
	fmt.Println("generating zip file of your project")
}
