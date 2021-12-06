package board

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestHasPlayableTile(t *testing.T) {
	t.Run("HasPlayableTile for no-tile board", func(t *testing.T) {
		given := getNoTileBoard()
		expected := false

		assert.Equal(t, expected, given.HasPlayableTile())
	})

	t.Run("HasPlayableTile for h board", func(t *testing.T) {
		given := getHBoard()
		expected := true

		assert.Equal(t, expected, given.HasPlayableTile())
	})

	t.Run("HasPlayableTile for two-tile board", func(t *testing.T) {
		given := getTwoTileBoard()
		expected := true

		assert.Equal(t, expected, given.HasPlayableTile())
	})
	t.Run("HasPlayableTile for two-tile board (already both two tile already picked) ", func(t *testing.T) {
		given := getTwoTileBoard()
		for _, tile := range given.Tiles {
			if tile != nil {
				tile.isPicked = true
			}
		}
		expected := false

		assert.Equal(t, expected, given.HasPlayableTile())
	})
}
func TestGetPlayableTiles(t *testing.T) {
	t.Run("GetPlayableTiles for no-tile board", func(t *testing.T) {
		given := getNoTileBoard()
		expected := []*Tile{}

		assert.Equal(t, expected, given.GetPlayableTiles())
	})

	t.Run("GetPlayableTiles for h board", func(t *testing.T) {
		givenBoard := getHBoard()
		gotPlayableTiles := givenBoard.GetPlayableTiles()

		expectedLen := 4
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 1, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 3, 0})]
		expectedTile4 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 3, 0})]
		expectedTile := []*Tile{expectedTile1, expectedTile2, expectedTile3, expectedTile4}

		assert.Equal(t, expectedLen, len(gotPlayableTiles))
		assert.Equal(t, expectedTile, gotPlayableTiles)
	})

	t.Run("GetPlayableTiles for two-tile board", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		gotPlayableTiles := givenBoard.GetPlayableTiles()

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 2, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]
		expectedTile := []*Tile{expectedTile1, expectedTile2}

		assert.Equal(t, expectedLen, len(gotPlayableTiles))
		assert.Equal(t, expectedTile, gotPlayableTiles)
	})

	t.Run("GetPlayableTiles for two-tile board (already both two tile already picked) ", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		for _, tile := range givenBoard.Tiles {
			if tile != nil {
				tile.isPicked = true
			}
		}
		gotPlayableTiles := givenBoard.GetPlayableTiles()

		expectedLen := 0
		expectedTile := []*Tile{}

		assert.Equal(t, expectedLen, len(gotPlayableTiles))
		assert.Equal(t, expectedTile, gotPlayableTiles)
	})

	t.Run("GetPlayableTiles for small-pyramid board", func(t *testing.T) {
		givenBoard := getSmallPyramidBoard()
		gotPlayableTiles := givenBoard.GetPlayableTiles()

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 1, 1})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 1, 1})]
		expectedTile := []*Tile{expectedTile1, expectedTile2}

		assert.Equal(t, expectedLen, len(gotPlayableTiles))
		assert.Equal(t, expectedTile, gotPlayableTiles)
	})
}

func TestIsTilePlayable(t *testing.T) {
	t.Run("isTilePlayable for small-pyramid board (partly block from above)", func(t *testing.T) {
		givenBoard := getSmallPyramidBoard()

		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 0, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 2, 0})]
		expectedTile4 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 2, 0})]
		expectedTile5 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 1, 1})]
		expectedTile6 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 1, 1})]

		assert.Equal(t, false, givenBoard.isTilePlayable(expectedTile1))
		assert.Equal(t, false, givenBoard.isTilePlayable(expectedTile2))
		assert.Equal(t, false, givenBoard.isTilePlayable(expectedTile3))
		assert.Equal(t, false, givenBoard.isTilePlayable(expectedTile4))
		assert.Equal(t, true, givenBoard.isTilePlayable(expectedTile5))
		assert.Equal(t, true, givenBoard.isTilePlayable(expectedTile6))
	})

	t.Run("isTilePlayable for h-tile board (middle tile is blocked from side)", func(t *testing.T) {
		givenBoard := getHBoard()

		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 1, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]
		expectedTile4 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 3, 0})]
		expectedTile5 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 3, 0})]

		assert.Equal(t, true, givenBoard.isTilePlayable(expectedTile1))
		assert.Equal(t, true, givenBoard.isTilePlayable(expectedTile2))
		assert.Equal(t, false, givenBoard.isTilePlayable(expectedTile3))
		assert.Equal(t, true, givenBoard.isTilePlayable(expectedTile4))
		assert.Equal(t, true, givenBoard.isTilePlayable(expectedTile5))
	})

	t.Run("isTilePlayable for tower board (tile is blocked from above)", func(t *testing.T) {
		givenBoard := getTowerBoard()

		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 1})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 2})]

		assert.Equal(t, false, givenBoard.isTilePlayable(expectedTile1))
		assert.Equal(t, false, givenBoard.isTilePlayable(expectedTile2))
		assert.Equal(t, true, givenBoard.isTilePlayable(expectedTile3))
	})
}

func TestFindPossiblePair(t *testing.T) {
	t.Run("FindPossiblePair for no-tile board", func(t *testing.T) {
		given := getNoTileBoard()
		expected := PairIterator{}

		assert.Equal(t, expected, given.FindPossiblePair())
	})

	t.Run("FindPossiblePair for h board", func(t *testing.T) {
		givenBoard := getHBoard()
		gotPossiblePairs := givenBoard.FindPossiblePair()

		givenTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 0})]
		givenTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 1, 0})]
		_ = givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]
		givenTile4 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 3, 0})]
		givenTile5 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 3, 0})]

		expectedLen := 6
		expectedPair1 := Pair{givenTile1, givenTile2}
		expectedPair2 := Pair{givenTile1, givenTile4}
		expectedPair3 := Pair{givenTile1, givenTile5}
		expectedPair4 := Pair{givenTile2, givenTile4}
		expectedPair5 := Pair{givenTile2, givenTile5}
		expectedPair6 := Pair{givenTile4, givenTile5}

		expectedPairs := PairIterator{expectedPair1, expectedPair2, expectedPair3, expectedPair4, expectedPair5, expectedPair6}
		assert.Equal(t, expectedLen, len(gotPossiblePairs))
		assert.Equal(t, expectedPairs, gotPossiblePairs)
	})

	t.Run("FindPossiblePair for h board when (1,1) and (1,3) are picked", func(t *testing.T) {
		givenBoard := getHBoard()

		givenTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 0})]
		givenTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 1, 0})]
		givenTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]
		givenTile4 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 3, 0})]
		givenTile5 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 3, 0})]

		givenTile1.isPicked = true
		givenTile4.isPicked = true

		gotPossiblePairs := givenBoard.FindPossiblePair()

		expectedLen := 3
		expectedPair1 := Pair{givenTile2, givenTile3}
		expectedPair2 := Pair{givenTile2, givenTile5}
		expectedPair3 := Pair{givenTile3, givenTile5}

		expectedPairs := PairIterator{expectedPair1, expectedPair2, expectedPair3}

		assert.Equal(t, expectedLen, len(gotPossiblePairs))
		assert.Equal(t, expectedPairs, gotPossiblePairs)
	})

	t.Run("FindPossiblePair for h board when (1,1) and (1,3) are picked and one has different face", func(t *testing.T) {
		givenBoard := getHBoard()
		givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 0})].isPicked = true
		givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 3, 0})].isPicked = true
		givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})].Face = 6

		gotPossiblePairs := givenBoard.FindPossiblePair()
		expectedLen := 1
		expectedPair := Pair{
			givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 1, 0})],
			givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 3, 0})],
		}

		expectedPairs := PairIterator{expectedPair}

		assert.Equal(t, expectedLen, len(gotPossiblePairs))
		assert.Equal(t, expectedPairs, gotPossiblePairs)
	})

	t.Run("FindPossiblePair for two-tile board", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		gotPossiblePairs := givenBoard.FindPossiblePair()

		expectedLen := 1
		expectedPair := Pair{
			givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 2, 0})],
			givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})],
		}
		expectedPairs := PairIterator{expectedPair}

		assert.Equal(t, expectedLen, len(gotPossiblePairs))
		assert.Equal(t, expectedPairs, gotPossiblePairs)
	})

	t.Run("FindPossiblePair for two-tile board when they have different face", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 2, 0})].Face = 1
		givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})].Face = 2

		gotPossiblePairs := givenBoard.FindPossiblePair()

		expectedLen := 0
		expectedPairs := PairIterator{}

		assert.Equal(t, expectedLen, len(gotPossiblePairs))
		assert.Equal(t, expectedPairs, gotPossiblePairs)
	})
}

func TestFindPairWithFace(t *testing.T) {
	givenBoard := getHBoard()

	givenTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 0})]
	givenTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 1, 0})]
	givenTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 3, 0})]
	givenTile4 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 3, 0})]

	givenTile1.Face = 1
	givenTile2.Face = 1
	givenTile3.Face = 2
	givenTile4.Face = 2

	gotPossiblePairs := givenBoard.FindPossiblePair()

	t.Run("FindPairWithFace for h board (when face = 1)", func(t *testing.T) {
		gotPossiblePairWithFaceEqualToOne := gotPossiblePairs.FindPairWithFace(1)

		expectedLen := 1
		expectedPair := Pair{givenTile1, givenTile2}
		expectedPairs := PairIterator{expectedPair}

		assert.Equal(t, expectedLen, len(gotPossiblePairWithFaceEqualToOne))
		assert.Equal(t, expectedPairs, gotPossiblePairWithFaceEqualToOne)
	})

	t.Run("FindPairWithFace for h board (when face = 2)", func(t *testing.T) {
		gotPossiblePairWithFaceEqualToTwo := gotPossiblePairs.FindPairWithFace(2)

		expectedLen := 1
		expectedPair := Pair{givenTile3, givenTile4}
		expectedPairs := PairIterator{expectedPair}

		assert.Equal(t, expectedLen, len(gotPossiblePairWithFaceEqualToTwo))
		assert.Equal(t, expectedPairs, gotPossiblePairWithFaceEqualToTwo)
	})

	t.Run("FindPairWithFace for h board (when face = 3)", func(t *testing.T) {
		gotPossiblePairWithFaceEqualToThree := gotPossiblePairs.FindPairWithFace(3)

		expectedLen := 0
		expectedPairs := PairIterator{}

		assert.Equal(t, expectedLen, len(gotPossiblePairWithFaceEqualToThree))
		assert.Equal(t, expectedPairs, gotPossiblePairWithFaceEqualToThree)
	})
}
func TestPickPair(t *testing.T) {
	t.Run("picked pair that have same id", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 2, 0})]
		givenTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]
		givenPair := Pair{givenTile1, givenTile2}

		gotPickedPairID := givenBoard.PickPair(givenPair)

		assert.Equal(t, givenTile1.Id, gotPickedPairID[0])
		assert.Equal(t, givenTile2.Id, gotPickedPairID[1])
		assert.Equal(t, true, givenTile1.isPicked)
		assert.Equal(t, true, givenTile2.isPicked)
	})

	t.Run("picked pair that have different id", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 2, 0})]
		givenTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]

		givenTile1.Face = 1
		givenTile2.Face = 2
		givenPair := Pair{givenTile1, givenTile2}

		gotPickedPairID := givenBoard.PickPair(givenPair)

		assert.Equal(t, -1, gotPickedPairID[0])
		assert.Equal(t, -1, gotPickedPairID[1])
		assert.Equal(t, false, givenTile1.isPicked)
		assert.Equal(t, false, givenTile2.isPicked)
	})
}

func TestRandomSolvableFace(t *testing.T) {
	t.Run("randomSolvableFace when board has odd number of tiles", func(t *testing.T) {
		givenBoard := getTowerBoard()
		gotError := givenBoard.RandomSolvableFace()
		expected := "number of tiles can not be odd"

		assert.Equal(t, expected, gotError.Error())
	})
	t.Run("randomSolvableFace when board has more than 144 tiles", func(t *testing.T) {
		givenBoard := getTwoHundredTileBoard()
		gotError := givenBoard.RandomSolvableFace()
		expected := "number of tiles can not be more than 144"

		assert.Equal(t, expected, gotError.Error())
	})
	t.Run("randomSolvableFace when there is only one playable tile when create", func(t *testing.T) {
		givenBoard := getTowerFourBoard()
		gotError := givenBoard.RandomSolvableFace()
		expected := "fail to random solvable face"

		assert.Equal(t, expected, gotError.Error())
	})

}
