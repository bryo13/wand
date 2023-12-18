package wand

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bryo13/wand/assets/templates/dirs"
	"github.com/bryo13/wand/assets/templates/files"
	"github.com/bryo13/wand/pkg"
)

const (
	colorBlue string = "\033[0;36m"
	colorRed  string = "\033[0;31m"
)

var (
	input    int
	rootPath string
	modTitle string
)

func init() {
	fmt.Println(colorBlue + "* Welcome to wand!")
	fmt.Println(colorBlue + "* Checking if go is installed")

	if !isGoInstalled() {
		fmt.Println(colorRed + "please install Go from 'https://go.dev/doc/install'")
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
	fmt.Println(colorBlue+"* Go lives in ", path)
	return true
}

func Layout() error {
	fmt.Print(colorBlue + "Project name: ")
	fmt.Scanf("%s", &rootPath)

	fmt.Print(colorBlue + "mod init: ")
	fmt.Scanf("%s", &modTitle)

	fmt.Print(colorBlue + "Choose dir type 1 for mini, 2 for mini with builds: ")
	fmt.Scanf("%d", &input)

	switch input {
	case 1:
		dirs.CreateMiniDirs(rootPath)
		files.WriteFiles(rootPath, modTitle)
	case 2:
		dirs.CreateMiniDirsWithBuilds(rootPath)
		files.WriteFiles(rootPath, modTitle)
	default:
		fmt.Println("Choose a layout")
		os.Exit(2)
	}

	defer pkg.InvokeCleanup(rootPath)
	return nil
}
