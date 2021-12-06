package file

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestGetAllFileNameinPath(t *testing.T) {
	t.Run("get filenames from exist path", func(t *testing.T) {
		givenPath := "./test_data/"
		expected := []string{"file1.txt", "file2.txt"}
		assert.Equal(t, expected, GetAllFileNameinPath(givenPath))
	})
	t.Run("get filenames from not-exist path", func(t *testing.T) {
		givenPath := "./not-exist-folder/"
		expected := []string{}
		assert.Equal(t, expected, GetAllFileNameinPath(givenPath))
	})
}

func TestGetBytesFromFile(t *testing.T) {
	t.Run("get bytes from exist file", func(t *testing.T) {
		givenPath := "./test_data/file1.txt"
		expected := []byte("data")

		gotByte, err := GetBytesFromFile(givenPath)
		assert.Equal(t, expected, gotByte)
		assert.Equal(t, nil, err)
	})

	t.Run("get bytes from not-exist file", func(t *testing.T) {
		givenPath := "./test_data/not-exist.txt"
		expected := []byte("")

		gotByte, err := GetBytesFromFile(givenPath)
		assert.Equal(t, expected, gotByte)
		assert.NotEqual(t, nil, err)
	})
}
