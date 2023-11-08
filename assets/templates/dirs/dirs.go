// creates the directory structure
package dirs

import (
	"fmt"
	"os"
)

var (
	mini = []string{"cmd", "pkg"}
)

func CreateHomeDir(projectName string) error {
	err := os.Mkdir(projectName, 0750)
	if err != nil {
		fmt.Println(fmt.Errorf("could not create the directories due to: %q", err))
		return err
	}

	return nil
}

func CreateMiniDirs(projectName string) error {

	for _, name := range mini {
		name = fmt.Sprintf("%s/%s", projectName, name)
		err := os.Mkdir(name, 0750)
		if err != nil {
			fmt.Println(fmt.Errorf("could not create the directories due to: %q", err))
			return err
		}
	}
	return nil
}
