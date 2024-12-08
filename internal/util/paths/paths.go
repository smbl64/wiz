// paths package provides helpers functions to interact with files and directories.
package paths

import (
	"os"
)

// Exists checks to see if the given path exists.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil

}
