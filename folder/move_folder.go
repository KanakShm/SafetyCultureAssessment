package folder

import (
	"errors"
	"fmt"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := f.folders
	res := []Folder{}

	// Loop through all folders
	for _, f := range folders {
		// Check if the current path contains the name
		if strings.Contains(f.Paths, name) {
			var index = strings.Index(f.Paths, name)

			var first = f.Paths[:index]
			var second = dst
			if index == 0 {
				second += "."
			}
			var third = f.Paths[index:]

			f.Paths = first + second + third
		}
	}

	return updatedFolders, nil
}