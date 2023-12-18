// creates the directory structure
package dirs

import (
	"fmt"
	"os"
)

var (
	mini           = []string{"cmd", "pkg"}
	miniWithBuilds = []string{"cmd", "pkg", "builds/darwin", "builds/windows", "builds/linux"}
)

func createParentDir(projectName string) error {
	err := os.Mkdir(projectName, 0750)
	if err != nil {
		fmt.Println(fmt.Errorf("could not create the directories due to: %q", err))
		return err
	}

	return nil
}

func CreateMiniDirs(rootDir string) error {
	return createDirs(rootDir, mini)
}

func CreateMiniDirsWithBuilds(rootDir string) error {
	return createDirs(rootDir, miniWithBuilds)
}

func createDirs(projectName string, dirs []string) error {
	createParentDir(projectName)

	for _, name := range dirs {
		name = fmt.Sprintf("%s/%s", projectName, name)
		err := os.MkdirAll(name, 0750)
		if err != nil {
			fmt.Println(fmt.Errorf("could not create the directories due to: %q", err))
			return err
		}
	}
	return nil
}
