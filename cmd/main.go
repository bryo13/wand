package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bryo13/wand/assets/templates/dirs"
	"github.com/bryo13/wand/assets/templates/files"
)

var (
	mini        int
	projectName string
	modTitle    string
)

func init() {
	fmt.Println("Welcome to wand!")
	fmt.Println("Checking if go is installed")
	if !isGoInstalled() {
		fmt.Println("please install Go from 'https://go.dev/doc/install'")
		os.Exit(1)
	}
}

func isGoInstalled() bool {
	// 1. we can look for go in /usr/local directory but that limits us to linux
	// 2. we can run go commands using exec to check if go is installed
	// 3. the safer option is to use the LookPath function in the exec package

	path, err := exec.LookPath("go")
	if err != nil {
		return false
	}
	fmt.Println("Go lives in ", path)
	return true
}

func main() {

	fmt.Print("Project name: ")
	fmt.Scanf("%s", &projectName)

	fmt.Print("mod init: ")
	fmt.Scanf("%s", &modTitle)

	fmt.Print("Choose dir type 1 for mini: ")
	fmt.Scanf("%d", &mini)
	if mini == 1 {
		dirs.CreateHomeDir(projectName)
		dirs.CreateMiniDirs(projectName)
		files.AddMainFile(projectName)
		files.AddModFile(projectName, modTitle)
	}
}
