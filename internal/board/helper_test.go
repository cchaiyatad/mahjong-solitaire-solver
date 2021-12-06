package board

import (
	"os"
	"testing"

	"github.com/cchaiyatad/mss/internal/layout"
)

func TestMain(m *testing.M) {
	os.Setenv("LAYOUT_PATH", "./layout_test/")
	layout.InitLayout()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func getBoard(layout string) *Board {
	board, _ := CreateBoard(layout)
	return board
}
func getNotExistBoard() *Board {
	return getBoard("not-exist")
}
func getSimpleBoard() *Board {
	return getBoard("simple")
}
func getNoTileBoard() *Board {
	return getBoard("no-tile")
}
func getOneTileBoard() *Board {
	return getBoard("one-tile")
}
func getOneTileOddBoard() *Board {
	return getBoard("one-tile-odd")
}
func getTwoTileBoard() *Board {
	return getBoard("two-tile")
}
func getHBoard() *Board {
	return getBoard("h-tile")
}
func getFiveTileBoard() *Board {
	return getBoard("five-tile")
}
func getSmallPyramidBoard() *Board {
	return getBoard("small-pyramid")
}
func getTowerBoard() *Board {
	return getBoard("tower")
}
func getTBoard() *Board {
	return getBoard("t-tile")
}
func getTowerFourBoard() *Board {
	return getBoard("tower-four")
}
func getTwoHundredTileBoard() *Board {
	return getBoard("two-hundred-tile")
}
func getOneTwoOneTwoOneBoard() *Board {
	return getBoard("one-two-one-two-one")
}
