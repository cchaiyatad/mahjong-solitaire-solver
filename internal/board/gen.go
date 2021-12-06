package board

import (
	"errors"
	"math/rand"
)

func (board *Board) RandomSolvableFace() error {
	countTry := 0
	countTiles := countTiles(board)

	if countTiles%2 == 1 {
		return errors.New("number of tiles can not be odd")
	}
	if countTiles > 144 {
		return errors.New("number of tiles can not be more than 144")
	}

start:
	board.ResetTiles()
	playableTiles := board.GetPlayableTiles()
	orders := make(PairIterator, 0)

	for !board.IsSolved() && len(playableTiles) > 1 {
		pair := getRandomPairFromTiles(playableTiles)
		orders = append(orders, pair)
		board.PickPair(pair)
		playableTiles = board.GetPlayableTiles()
	}

	if !board.IsSolved() {
		if countTry < 5 {
			countTry++
			goto start
		}
		return errors.New("fail to random solvable face")
	}

	board.ResetTiles()

	tileIds := getShuffleFace()
	for idx, pair := range orders {
		face := tileIds[idx]

		board.Tiles[pair[0].Id].Face = face
		board.Tiles[pair[1].Id].Face = face
	}
	return nil
}
func countTiles(board *Board) int {
	count := 0
	for _, tile := range board.Tiles {
		if tile != nil {
			count++
		}
	}
	return count
}

func getShuffleFace() []Face {
	faces := getFacesInSuits()

	rand.Shuffle(len(faces), func(i, j int) {
		faces[i], faces[j] = faces[j], faces[i]
	})
	return faces
}

func getRandomPairFromTiles(tiles []*Tile) Pair {
	rand.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	return Pair{tiles[0], tiles[1]}
}
