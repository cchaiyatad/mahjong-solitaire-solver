package solver

import (
	"sort"

	"github.com/cchaiyatad/mss/internal/board"
)

type Strategy interface {
	Solve(board *board.Board, heuristic Heuristic, callBack func(pickedTileIDs []int))
}

type MultipleFirst struct{}

func (*MultipleFirst) Solve(board *board.Board, heuristic Heuristic, callBack func(pickedTileIDs []int)) {
	pairs := board.FindPossiblePair()
	for len(pairs) > 0 {
		sortDecsByHeuristic(heuristic, board, pairs)

		playableTiles := board.GetPlayableTiles()

		face, count := getMostFaceCountTileFromPlayableTiles(playableTiles)
		candidatePairs := pairs.FindPairWithFace(face)

		if count == 4 {
			for _, pair := range candidatePairs {
				callBack(board.PickPair(pair))
			}
		} else {
			callBack(board.PickPair(candidatePairs[0]))
		}

		pairs = board.FindPossiblePair()
	}
}

func getMostFaceCountTileFromPlayableTiles(playableTiles []*board.Tile) (board.Face, int) {
	counts := make(map[board.Face]int)
	faces := make([]board.Face, 0)

	for _, tile := range playableTiles {
		counts[tile.Face] += 1
		faces = append(faces, tile.Face)
	}

	more := func(i, j int) bool {
		return counts[faces[i]] > counts[faces[j]]
	}

	sort.Slice(faces, more)
	return faces[0], counts[faces[0]]
}

type RandomStrategy struct{}

func (*RandomStrategy) Solve(board *board.Board, heuristic Heuristic, callBack func(pickedTileIDs []int)) {
	pairs := board.FindPossiblePair()
	for len(pairs) > 0 {
		sortDecsByHeuristic(heuristic, board, pairs)

		callBack(board.PickPair(pairs[0]))
		pairs = board.FindPossiblePair()
	}
}
