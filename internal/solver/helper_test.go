package solver

import (
	"os"
	"testing"

	"github.com/cchaiyatad/mss/internal/board"
	"github.com/cchaiyatad/mss/internal/layout"
)

func TestMain(m *testing.M) {
	os.Setenv("LAYOUT_PATH", "./layout_test/")
	layout.InitLayout()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func getBoard(layout string) *board.Board {
	board, _ := board.CreateBoard(layout)
	return board
}

func getOneTileBoard() *board.Board {
	return getBoard("one-tile")
}
func getOneRowBoard() *board.Board {
	return getBoard("one-row")
}
func getOneTwoOneTwoOneBoard() *board.Board {
	return getBoard("one-two-one-two-one")
}
func getTowerBoard() *board.Board {
	return getBoard("tower")
}
func getOneTwoThreeVerBoard() *board.Board {
	return getBoard("one-two-three-ver")
}
func getHBoard() *board.Board {
	return getBoard("h-tile")
}
func getSixteenBoard() *board.Board {
	return getBoard("sixteen-tile")
}

func getIdx(b *board.Board, vector board.Vector3) int {
	sizeX, sizeY, _ := b.Size.ToInt()
	x, y, z := vector.ToInt()
	return (z * sizeX * sizeY) + (y * sizeX) + x
}

type idHeuristic struct{}

func (idHeuristic) GetH(_ *board.Board, pair board.Pair) int {
	return pair[0].Id + pair[1].Id
}
