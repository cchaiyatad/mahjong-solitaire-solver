package board

import "sort"

type Pair [2]*Tile
type PairIterator []Pair

func (board *Board) FindPossiblePair() PairIterator {
	playabletiles := board.GetPlayableTiles()
	pairs := make(PairIterator, 0)

	for idx1, tile1 := range playabletiles {
		for idx2 := idx1 + 1; idx2 < len(playabletiles); idx2++ {
			tile2 := playabletiles[idx2]
			candidatePair := Pair{tile1, tile2}
			if candidatePair.isPair() {
				pairs = append(pairs, candidatePair)
			}
		}
	}

	return pairs
}

func (pairs PairIterator) FindPairWithFace(face Face) PairIterator {
	pairsWithFace := make(PairIterator, 0)
	for _, pair := range pairs {
		if pair[0].Face == face {
			pairsWithFace = append(pairsWithFace, pair)
		}
	}
	return pairsWithFace
}

func (pairs PairIterator) Sort(less func(i int, j int) bool) {
	sort.Slice(pairs, less)
}

func (pair Pair) isPair() bool {
	if pair[0] == nil || pair[1] == nil {
		return false
	}

	return pair[0].Face == pair[1].Face
}

func (board *Board) PickPair(pair Pair) []int {
	if pair.isPair() && !pair[0].isPicked && !pair[1].isPicked {
		return []int{board.pickTile(pair[0]), board.pickTile(pair[1])}
	}
	return []int{-1, -1}
}
