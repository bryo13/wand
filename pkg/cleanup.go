package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

func changeDir(path string) error {
	if err := os.Chdir(path); err != nil {
		fmt.Println(fmt.Errorf("could not change dir to %s due to: %q", path, err))
		return err
	}
	return nil
}

func autoImport(path string) error {
	path = fmt.Sprintf("%s/%s/%s", path, "cmd", "main.go")

	_ = path
	return nil
}

func tidyMod(path string) error {
	cmd := exec.Command("go", "mod", "tidy")
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Errorf("could not tidy mod due to: %q", err))
		return err
	}
	return nil
}

func compile() {
	// compile the project and
	// place it in specific folders
}

func InvokeCleanup(rootPath string) {
	changeDir(rootPath)
	autoImport(rootPath)
	tidyMod(rootPath)
}
