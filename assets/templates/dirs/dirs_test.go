package dirs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDirs(t *testing.T) {
	parent_name := "test"
	t.Run("create parent dir", func(t *testing.T) {
		err := createParentDir(parent_name)
		require.NoError(t, err, "Dir should be created without an error")
		require.DirExists(t, parent_name, "Directory should have been created")
	})

	t.Run("test creation of mini dirs", func(t *testing.T) {
		err := CreateMiniDirs(parent_name)
		require.DirExists(t, "test/cmd", "Dir cmd should be present")
		require.DirExists(t, "test/pkg", "Dir pkg should be present")
		require.NoError(t, err, "mini dirs should be created without an error")
	})

	defer func() {
		os.RemoveAll("test")
	}()
}
