package file

import (
	"io/ioutil"
	"os"
)

func GetAllFileNameinPath(path string) []string {
	filenames := make([]string, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return filenames
	}

	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames
}

func GetBytesFromFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}
