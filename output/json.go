package output

import (
	"fmt"
	"os"
	"path/filepath"
)

// toJSON export to json file
func toJSON(dest string, bytes []byte) error {
	fullname := dest
	dir, filename := filepath.Split(fullname)
	if dir == "" {
		tmpDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return err
		}
		dir = tmpDir
		fullname = fmt.Sprintf("%s/%s", dir, filename)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("%s not exists", dir)
	}

	file, err := os.Create(fullname)
	if err == nil {
		defer file.Close()
		file.Write(bytes)
		fmt.Println("file saved: ", fullname)
	}
	return err
}
