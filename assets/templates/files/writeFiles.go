package files

import (
	"fmt"
	"os"
	"runtime"
)

const main = `
package main

func main() {
	fmt.Println("Hello wand!")
}
`

func AddMainFile(projectName string) {
	name := fmt.Sprintf("%s/%s/%s", projectName, "cmd", "main.go")
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(fmt.Errorf("could not open main file due to: %q", err))
	}
	if _, err := f.Write([]byte(main)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		fmt.Println(fmt.Errorf("could not write main the file due to: %q", err))
	}
	if err := f.Close(); err != nil {
		fmt.Println(fmt.Errorf("could not close main file due to: %q", err))
	}
}

func AddModFile(projectName string, modTitle string) {
	version := runtime.Version()

	version = version[:2] + " " + version[2:]
	file := fmt.Sprintf("%s/%s", projectName, "go.mod")
	modText := fmt.Sprintf("%s %s\n\n%s\n", "module", modTitle, string(version))

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(fmt.Errorf("could not open mod file due to: %q", err))
	}

	if _, err := f.Write([]byte(modText)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		fmt.Println(fmt.Errorf("could not write mod file due to: %q", err))
	}

	if err := f.Close(); err != nil {
		fmt.Println(fmt.Errorf("could not close mod file due to: %q", err))
	}
}
