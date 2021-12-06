package solver

import (
	"math/rand"
	"sort"

	b "github.com/cchaiyatad/mss/internal/board"
	"github.com/cchaiyatad/mss/internal/utils"
)

type Heuristic interface {
	GetH(board *b.Board, pair b.Pair) int
}

type MaxBlock struct{}

func (maxBlock *MaxBlock) GetH(board *b.Board, pair b.Pair) int {
	return maxBlock.getHAt(board, pair[0].Position) + maxBlock.getHAt(board, pair[1].Position)
}

func (maxBlock *MaxBlock) getHAt(board *b.Board, position b.Vector3) int {
	return maxBlock.hor(board, position) + maxBlock.ver(board, position)
}

func (*MaxBlock) hor(board *b.Board, position b.Vector3) int {
	diff := len(board.GetObstructorsLeft(position)) - len(board.GetObstructorsRight(position))
	return utils.Abs(diff / 2)
}

func (maxBlock *MaxBlock) ver(board *b.Board, position b.Vector3) int {
	touchingTiles := board.GetObstructorsBelow(position)
	verValue := 0
	for _, tile := range touchingTiles {
		verValue += maxBlock.hor(board, tile.Position) + 1
	}
	return verValue
}

type RandomHeuristic struct{}

func (*RandomHeuristic) GetH(board *b.Board, pair b.Pair) int {
	return rand.Intn(136)
}

func getPairMoreFuncByHeuristic(h Heuristic, board *b.Board) func(pairs b.PairIterator) func(i, j int) bool {
	return func(pairs b.PairIterator) func(i, j int) bool {
		return func(i, j int) bool {
			return h.GetH(board, pairs[i]) > h.GetH(board, pairs[j])
		}
	}
}

func sortDecsByHeuristic(heuristic Heuristic, board *b.Board, pairs b.PairIterator) {
	more := getPairMoreFuncByHeuristic(heuristic, board)(pairs)
	sort.Slice(pairs, more)
}
