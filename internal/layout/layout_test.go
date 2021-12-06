package layout

import (
	"os"
	"testing"

	"github.com/go-playground/assert"
)

func TestMain(m *testing.M) {
	os.Setenv(key, "./layout_test/")
	InitLayout()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestInitLayouts(t *testing.T) {
	given := "simple"
	expected := "simple.xml"
	assert.Equal(t, expected, layouts[given])
}

func TestLayoutIsValid(t *testing.T) {
	givenExistLayout := "simple"
	givenNotExistLayout := "not-exist"

	assert.Equal(t, true, IsLayoutValid(givenExistLayout))
	assert.Equal(t, false, IsLayoutValid(givenNotExistLayout))
}
func TestGetVector3UniqueKey(t *testing.T) {
	var given_x, given_y, given_z uint8
	given_x, given_y, given_z = 0, 10, 3
	expected := "0_10_3"

	assert.Equal(t, expected, getVector3UniqueKey(given_x, given_y, given_z))
}
func TestGetLayoutOption(t *testing.T) {
	expected := []string{"simple"}
	assert.Equal(t, expected, GetLayoutOption())
}
func TestGetLayout(t *testing.T) {
	t.Run("get layout from simple (exist) layout", func(t *testing.T) {
		given := "simple"
		gotLayout, err := getLayout(given)

		assert.Equal(t, "simple", gotLayout.Name)
		assert.Equal(t, Size{8, 8, 2}, gotLayout.Size)
		assert.Equal(t, nil, err)
	})
	t.Run("get layout from not-exist layout", func(t *testing.T) {
		given := "not-exist"
		gotLayout, err := getLayout(given)
		expectedErrorText := `error: layout "not-exist" not exist`

		assert.Equal(t, nil, gotLayout)
		assert.Equal(t, expectedErrorText, err.Error())

	})
}
func TestGetSize(t *testing.T) {
	t.Run("get size from simple (exist) layout", func(t *testing.T) {
		given := "simple"
		gotSize, err := GetSize(given)

		assert.Equal(t, [3]uint8{8, 8, 2}, gotSize)
		assert.Equal(t, nil, err)
	})

	t.Run("get Size from not-exist layout", func(t *testing.T) {
		given := "not-exist"
		gotSize, err := GetSize(given)
		expectedErrorText := `error: layout "not-exist" not exist`

		assert.Equal(t, [3]uint8{}, gotSize)
		assert.Equal(t, expectedErrorText, err.Error())

	})
}

func TestGetHasTilesFunc(t *testing.T) {
	t.Run("get HasTilesFunc from simple (exist) layout", func(t *testing.T) {
		given := "simple"
		gotFunc, err := GetHasTilesFunc(given)

		assert.Equal(t, true, gotFunc(4, 2, 0))
		assert.Equal(t, false, gotFunc(0, 10, 2))
		assert.Equal(t, false, gotFunc(14, 12, 20))
		assert.Equal(t, nil, err)
	})

	t.Run("get HasTilesFunc from not-exist layout", func(t *testing.T) {
		given := "not-exist"
		gotFunc, err := GetHasTilesFunc(given)
		expectedErrorText := `error: layout "not-exist" not exist`

		assert.Equal(t, nil, gotFunc)
		assert.Equal(t, expectedErrorText, err.Error())
	})
}
