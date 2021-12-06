package board

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestCreateBoard(t *testing.T) {
	t.Run("create board with exist layout", func(t *testing.T) {
		givenBoard, err := CreateBoard("simple")

		assert.Equal(t, 8*8*2, len(givenBoard.Tiles))
		assert.Equal(t, Size{8, 8, 2}, givenBoard.Size)
		assert.Equal(t, "simple", givenBoard.Layout)

		assert.Equal(t, nil, err)
	})
	t.Run("create board with not-exist layout", func(t *testing.T) {
		givenBoard, err := CreateBoard("not-exist")

		assert.Equal(t, nil, givenBoard)
		assert.Equal(t, "error: can not create board: error: layout \"not-exist\" not exist", err.Error())
	})
}

func TestIsSolve(t *testing.T) {
	t.Run("IsSolve no-tile board", func(t *testing.T) {
		given := getNoTileBoard()
		expected := true

		assert.Equal(t, expected, given.IsSolved())
	})

	t.Run("IsSolve not gen simple board", func(t *testing.T) {
		given := getSimpleBoard()
		expected := false

		assert.Equal(t, expected, given.IsSolved())
	})

	t.Run("IsSolve already gen simple board", func(t *testing.T) {
		given := getSimpleBoard()
		_ = given.RandomSolvableFace()
		expected := false

		assert.Equal(t, expected, given.IsSolved())
	})

	t.Run("IsSolve not exist board", func(t *testing.T) {
		given := getNotExistBoard()
		expected := false

		assert.Equal(t, expected, given.IsSolved())
	})
}
func TestToJSON(t *testing.T) {
	t.Run("to json simple board", func(t *testing.T) {
		given := getSimpleBoard()
		expected := []byte(`{"tiles":[{"id":0,"face":0,"position":{"x":0,"y":0,"z":0}},{"id":2,"face":0,"position":{"x":2,"y":0,"z":0}},{"id":4,"face":0,"position":{"x":4,"y":0,"z":0}},{"id":6,"face":0,"position":{"x":6,"y":0,"z":0}},{"id":16,"face":0,"position":{"x":0,"y":2,"z":0}},{"id":18,"face":0,"position":{"x":2,"y":2,"z":0}},{"id":20,"face":0,"position":{"x":4,"y":2,"z":0}},{"id":22,"face":0,"position":{"x":6,"y":2,"z":0}},{"id":32,"face":0,"position":{"x":0,"y":4,"z":0}},{"id":34,"face":0,"position":{"x":2,"y":4,"z":0}},{"id":36,"face":0,"position":{"x":4,"y":4,"z":0}},{"id":38,"face":0,"position":{"x":6,"y":4,"z":0}},{"id":48,"face":0,"position":{"x":0,"y":6,"z":0}},{"id":50,"face":0,"position":{"x":2,"y":6,"z":0}},{"id":52,"face":0,"position":{"x":4,"y":6,"z":0}},{"id":54,"face":0,"position":{"x":6,"y":6,"z":0}},{"id":82,"face":0,"position":{"x":2,"y":2,"z":1}},{"id":84,"face":0,"position":{"x":4,"y":2,"z":1}},{"id":98,"face":0,"position":{"x":2,"y":4,"z":1}},{"id":100,"face":0,"position":{"x":4,"y":4,"z":1}}],"size":{"x_size":8,"y_size":8,"z_size":2},"layout":"simple"}`)

		assert.Equal(t, expected, given.ToJSON())
	})

	t.Run("to json not exist board", func(t *testing.T) {
		given := getNotExistBoard()
		expected := []byte("null")

		assert.Equal(t, expected, given.ToJSON())
	})
}
func TestPickTile(t *testing.T) {
	t.Run("pick tile in a simple board", func(t *testing.T) {
		givenBoard := getSimpleBoard()
		givenTile := givenBoard.Tiles[0]
		expected := 0

		assert.Equal(t, expected, givenBoard.pickTile(givenTile))
	})
	t.Run("pick a nil tile", func(t *testing.T) {
		givenBoard := getSimpleBoard()
		var givenTile *Tile
		expected := -1

		assert.Equal(t, expected, givenBoard.pickTile(givenTile))
	})

	t.Run("PickTile that does not in a board (with not exist id in board)", func(t *testing.T) {
		givenBoard := getSimpleBoard()
		givenTile := &Tile{Id: 35}
		expected := -1

		assert.Equal(t, expected, givenBoard.pickTile(givenTile))
	})

	t.Run("PickTile that does not in a board (with same id)", func(t *testing.T) {
		givenBoard := getSimpleBoard()
		givenTile := &Tile{Id: 0}
		expected := -1

		assert.Equal(t, expected, givenBoard.pickTile(givenTile))
	})
}

func TestTileByTopLeftAt(t *testing.T) {
	t.Run("getTileByTopLeftAt in one-tile board at {2, 2, 0} (at tile top left postion)", func(t *testing.T) {
		givenBoard := getOneTileBoard()
		givenPosition := Vector3{2, 2, 0}
		expected := givenBoard.Tiles[14]

		assert.Equal(t, expected, givenBoard.getTileByTopLeftAt(givenPosition))
	})
	t.Run("getTileByTopLeftAt in one-tile-odd board at {1, 1, 0} (at tile top left postion)", func(t *testing.T) {
		givenBoard := getOneTileOddBoard()
		givenPosition := Vector3{1, 1, 0}
		expected := givenBoard.Tiles[7]

		assert.Equal(t, expected, givenBoard.getTileByTopLeftAt(givenPosition))
	})
	t.Run("getTileByTopLeftAt in one-tile board at {3, 2, 0} (at tile top right postion)", func(t *testing.T) {
		givenBoard := getOneTileBoard()
		givenPosition := Vector3{3, 2, 0}
		var expected *Tile

		assert.Equal(t, expected, givenBoard.getTileByTopLeftAt(givenPosition))
	})
	t.Run("getTileByTopLeftAt in one-tile board at {2, 3, 0} (at tile bottom left postion)", func(t *testing.T) {
		givenBoard := getOneTileBoard()
		givenPosition := Vector3{2, 3, 0}
		var expected *Tile

		assert.Equal(t, expected, givenBoard.getTileByTopLeftAt(givenPosition))
	})
	t.Run("getTileByTopLeftAt in one-tile board at {3, 3, 0} (at tile bottom right left postion)", func(t *testing.T) {
		givenBoard := getOneTileBoard()
		givenPosition := Vector3{3, 3, 0}
		var expected *Tile

		assert.Equal(t, expected, givenBoard.getTileByTopLeftAt(givenPosition))
	})
}

func TestGetTileAt(t *testing.T) {
	t.Run("getTileAt in one-tile board at {2, 2, 0} (at tile top left postion)", func(t *testing.T) {
		givenBoard := getOneTileBoard()
		givenPosition := Vector3{2, 2, 0}
		expected := givenBoard.Tiles[14]

		assert.Equal(t, expected, givenBoard.getTileAt(givenPosition))
	})
	t.Run("getTileAt in one-tile board at {3, 2, 0} (at tile top right postion)", func(t *testing.T) {
		givenBoard := getOneTileBoard()
		givenPosition := Vector3{3, 2, 0}
		expected := givenBoard.Tiles[14]

		assert.Equal(t, expected, givenBoard.getTileAt(givenPosition))
	})
	t.Run("getTileAt in one-tile board at {2, 3, 0} (at tile bottom left postion)", func(t *testing.T) {
		givenBoard := getOneTileBoard()
		givenPosition := Vector3{2, 3, 0}
		expected := givenBoard.Tiles[14]

		assert.Equal(t, expected, givenBoard.getTileAt(givenPosition))
	})
	t.Run("getTileAt in one-tile board at {3, 3, 0} (at tile bottom right postion)", func(t *testing.T) {
		givenBoard := getOneTileBoard()
		givenPosition := Vector3{3, 3, 0}
		expected := givenBoard.Tiles[14]

		assert.Equal(t, expected, givenBoard.getTileAt(givenPosition))
	})

	t.Run("getTileAt in one-tile-odd board at {1, 1, 0} (at tile top left postion)", func(t *testing.T) {
		givenBoard := getOneTileOddBoard()
		givenPosition := Vector3{1, 1, 0}
		expected := givenBoard.Tiles[7]

		assert.Equal(t, expected, givenBoard.getTileAt(givenPosition))
	})

	t.Run("getTileAt in two-tile board at {3, 2, 0} (at tile top left postion) should get second tile", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenPosition := Vector3{3, 2, 0}
		expected := givenBoard.Tiles[15]

		assert.Equal(t, expected, givenBoard.getTileAt(givenPosition))
	})
	t.Run("getTileAt in two-tile board at {4, 2, 0} (at tile top right postion) should get second tile", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenPosition := Vector3{4, 2, 0}
		expected := givenBoard.Tiles[15]

		assert.Equal(t, expected, givenBoard.getTileAt(givenPosition))
	})
	t.Run("getTileAt in two-tile board at {3, 3, 0} (at tile bottom left postion) should get second tile", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenPosition := Vector3{3, 3, 0}
		expected := givenBoard.Tiles[15]

		assert.Equal(t, expected, givenBoard.getTileAt(givenPosition))
	})
	t.Run("getTileAt in two-tile board at {4, 3, 0} (at tile bottom right postion) should get second tile", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenPosition := Vector3{4, 3, 0}
		expected := givenBoard.Tiles[15]

		assert.Equal(t, expected, givenBoard.getTileAt(givenPosition))
	})
}

func TestGetTouchTileLeft(t *testing.T) {
	t.Run("getTouchTileLeft on two-tile board (whole tile is touched)", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenPosition := Vector3{3, 2, 0}
		gotTouch := givenBoard.getTouchTileLeft(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[13]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})

	t.Run("getTouchTileLeft on h-tile board (part of two tiles is touched) ({3, 2, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{3, 2, 0}
		gotTouch := givenBoard.getTouchTileLeft(givenPosition)

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[9]
		expectedTile2 := givenBoard.Tiles[25]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile1, gotTouch[0])
		assert.Equal(t, expectedTile2, gotTouch[1])
	})
	t.Run("getTouchTileLeft on h-tile board ({5, 1, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{5, 1, 0}
		gotTouch := givenBoard.getTouchTileLeft(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})
	t.Run("getTouchTileLeft on h-tile board ({5, 3, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{5, 3, 0}
		gotTouch := givenBoard.getTouchTileLeft(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})

	t.Run("getTouchTileLeft on two-tile board (no left touched tile)", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenPosition := Vector3{1, 2, 0}
		gotTouch := givenBoard.getTouchTileLeft(givenPosition)

		expectedLen := 0

		assert.Equal(t, expectedLen, len(gotTouch))
	})
}

func TestGetTouchTileRight(t *testing.T) {
	t.Run("getTouchTileRight on two-tile board (whole tile is touched)", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenPosition := Vector3{1, 2, 0}
		gotTouch := givenBoard.getTouchTileRight(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[15]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})

	t.Run("getTouchTileRight on h-tile board (part of two tiles is touched)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{3, 2, 0}
		gotTouch := givenBoard.getTouchTileRight(givenPosition)

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[13]
		expectedTile2 := givenBoard.Tiles[29]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile1, gotTouch[0])
		assert.Equal(t, expectedTile2, gotTouch[1])
	})

	t.Run("getTouchTileRight on h-tile board ({1, 1, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{1, 1, 0}
		gotTouch := givenBoard.getTouchTileRight(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})
	t.Run("getTouchTileRight on h-tile board ({1, 3, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{1, 3, 0}
		gotTouch := givenBoard.getTouchTileRight(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})

	t.Run("getTouchTileRight on two-tile board (no right touched tile)", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenPosition := Vector3{3, 2, 0}
		gotTouch := givenBoard.getTouchTileRight(givenPosition)

		expectedLen := 0
		assert.Equal(t, expectedLen, len(gotTouch))
	})
}

func TestGetTouchTileAbove(t *testing.T) {
	t.Run("getTouchTileAbove on five-tile board (bottom right is touched)", func(t *testing.T) {
		givenBoard := getFiveTileBoard()
		givenPosition := Vector3{0, 0, 0}
		gotTouch := givenBoard.getTouchTileAbove(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 1})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})
	t.Run("getTouchTileAbove on five-tile board (bottom left is touched)", func(t *testing.T) {
		givenBoard := getFiveTileBoard()
		givenPosition := Vector3{2, 0, 0}
		gotTouch := givenBoard.getTouchTileAbove(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 1})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})
	t.Run("getTouchTileAbove on five-tile board (top right is touched)", func(t *testing.T) {
		givenBoard := getFiveTileBoard()
		givenPosition := Vector3{0, 2, 0}
		gotTouch := givenBoard.getTouchTileAbove(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 1})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})
	t.Run("getTouchTileAbove on five-tile board (top left is touched)", func(t *testing.T) {
		givenBoard := getFiveTileBoard()
		givenPosition := Vector3{2, 2, 0}
		gotTouch := givenBoard.getTouchTileAbove(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 1})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})
	t.Run("getTouchTileAbove on small-Pyramid Board (bottom is touched)", func(t *testing.T) {
		givenBoard := getSmallPyramidBoard()
		givenPosition := Vector3{0, 0, 0}
		gotTouch := givenBoard.getTouchTileAbove(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 1, 1})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})

	t.Run("getTouchTileAbove on small-Pyramid Board (top is touched)", func(t *testing.T) {
		givenBoard := getSmallPyramidBoard()
		givenPosition := Vector3{0, 2, 0}
		gotTouch := givenBoard.getTouchTileAbove(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 1, 1})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})
	t.Run("getTouchTileAbove on tower Board (whole tile is touched)", func(t *testing.T) {
		givenBoard := getTowerBoard()
		givenPosition := Vector3{0, 0, 0}
		gotTouch := givenBoard.getTouchTileAbove(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 1})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})
}
func TestGetTouchTileBelow(t *testing.T) {
	t.Run("getTouchTileBelow on five-tile board (partly touch)", func(t *testing.T) {
		givenBoard := getFiveTileBoard()
		givenPosition := Vector3{1, 1, 1}
		gotTouch := givenBoard.getTouchTileBelow(givenPosition)

		expectedLen := 4
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 0, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 2, 0})]
		expectedTile4 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 2, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2, expectedTile3, expectedTile4}

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch)
	})

	t.Run("getTouchTileBelow on tower Board (whole tile is touched)", func(t *testing.T) {
		givenBoard := getTowerBoard()
		givenPosition := Vector3{0, 0, 2}
		gotTouch := givenBoard.getTouchTileBelow(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 1})]

		assert.Equal(t, expectedLen, len(gotTouch))
		assert.Equal(t, expectedTile, gotTouch[0])
	})
}

func TestGetObstructorsLeft(t *testing.T) {
	t.Run("GetObstructorsLeft on two-tile board", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenPosition := Vector3{3, 2, 0}
		gotObstructors := givenBoard.GetObstructorsLeft(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 2, 0})]

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors[0])
	})

	t.Run("GetObstructorsLeft on H-tile board (of {3, 2, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{3, 2, 0}
		gotObstructors := givenBoard.GetObstructorsLeft(givenPosition)

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 3, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
	t.Run("GetObstructorsLeft on H-tile board (of {5, 1, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{5, 1, 0}
		gotObstructors := givenBoard.GetObstructorsLeft(givenPosition)

		expectedLen := 3
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 3, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2, expectedTile3}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
	t.Run("GetObstructorsLeft on H-tile board (of {5, 3, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{5, 3, 0}
		gotObstructors := givenBoard.GetObstructorsLeft(givenPosition)

		expectedLen := 3
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 3, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2, expectedTile3}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
	t.Run("GetObstructorsLeft on 1-2-1-2-1 board (of {8, 1, 0} tile)", func(t *testing.T) {
		givenBoard := getOneTwoOneTwoOneBoard()
		givenPosition := Vector3{8, 1, 0}
		gotObstructors := givenBoard.GetObstructorsLeft(givenPosition)

		expectedLen := 6
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{6, 0, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{6, 2, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{4, 1, 0})]
		expectedTile4 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 0, 0})]
		expectedTile5 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 2, 0})]
		expectedTile6 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 1, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2, expectedTile3, expectedTile4, expectedTile5, expectedTile6}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
}
func TestGetObstructorsRight(t *testing.T) {
	t.Run("GetObstructorsRight on two-tile board", func(t *testing.T) {
		givenBoard := getTwoTileBoard()
		givenPosition := Vector3{1, 2, 0}
		gotObstructors := givenBoard.GetObstructorsRight(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors[0])
	})

	t.Run("GetObstructorsRight on H-tile board (of {3, 2, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{3, 2, 0}
		gotObstructors := givenBoard.GetObstructorsRight(givenPosition)

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 1, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 3, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
	t.Run("GetObstructorsRight on H-tile board (of {1, 1, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{1, 1, 0}
		gotObstructors := givenBoard.GetObstructorsRight(givenPosition)

		expectedLen := 3
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 1, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 3, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2, expectedTile3}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
	t.Run("GetObstructorsRight on H-tile board (of {1, 3, 0} tile)", func(t *testing.T) {
		givenBoard := getHBoard()
		givenPosition := Vector3{1, 3, 0}
		gotObstructors := givenBoard.GetObstructorsRight(givenPosition)

		expectedLen := 3
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{3, 2, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 1, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{5, 3, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2, expectedTile3}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
	t.Run("GetObstructorsRight on 1-2-1-2-1 board (of {0, 1, 0} tile)", func(t *testing.T) {
		givenBoard := getOneTwoOneTwoOneBoard()
		givenPosition := Vector3{0, 1, 0}
		gotObstructors := givenBoard.GetObstructorsRight(givenPosition)

		expectedLen := 6
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 0, 0})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 2, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{4, 1, 0})]
		expectedTile4 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{6, 0, 0})]
		expectedTile5 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{6, 2, 0})]
		expectedTile6 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{8, 1, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2, expectedTile3, expectedTile4, expectedTile5, expectedTile6}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})

}
func TestGetObstructorsAbove(t *testing.T) {
	t.Run("GetObstructorsAbove on tower board (on {0, 0, 1})", func(t *testing.T) {
		givenBoard := getTowerBoard()
		givenPosition := Vector3{0, 0, 1}
		gotObstructors := givenBoard.GetObstructorsAbove(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 2})]

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors[0])
	})
	t.Run("GetObstructorsAbove on tower board (on {0, 0, 0})", func(t *testing.T) {
		givenBoard := getTowerBoard()
		givenPosition := Vector3{0, 0, 0}
		gotObstructors := givenBoard.GetObstructorsAbove(givenPosition)

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 1})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 2})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
	t.Run("GetObstructorsAbove on tower board (on {0, 0, 2)", func(t *testing.T) {
		givenBoard := getTowerBoard()
		givenPosition := Vector3{0, 0, 2}
		gotObstructors := givenBoard.GetObstructorsAbove(givenPosition)

		expectedLen := 0
		expectedTile := TilesIterator{}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})

	t.Run("GetObstructorsAbove on t-tile board (on {2, 0, 0})", func(t *testing.T) {
		givenBoard := getTBoard()
		givenPosition := Vector3{2, 0, 0}
		gotObstructors := givenBoard.GetObstructorsAbove(givenPosition)

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 1})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 2})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
	t.Run("GetObstructorsAbove on t-tile board (on {0, 1, 0})", func(t *testing.T) {
		givenBoard := getTBoard()
		givenPosition := Vector3{0, 1, 0}
		gotObstructors := givenBoard.GetObstructorsAbove(givenPosition)

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 1})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 2})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
	t.Run("GetObstructorsAbove on t-tile board (on {2, 2, 0})", func(t *testing.T) {
		givenBoard := getTBoard()
		givenPosition := Vector3{2, 2, 0}
		gotObstructors := givenBoard.GetObstructorsAbove(givenPosition)

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 1})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 2})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
}

func TestGetObstructorsBelow(t *testing.T) {
	t.Run("GetObstructorsBelow on tower board (on {0, 0, 1})", func(t *testing.T) {
		givenBoard := getTowerBoard()
		givenPosition := Vector3{0, 0, 1}
		gotObstructors := givenBoard.GetObstructorsBelow(givenPosition)

		expectedLen := 1
		expectedTile := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 0})]

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors[0])
	})
	t.Run("GetObstructorsBelow on tower board (on {0, 0, 2})", func(t *testing.T) {
		givenBoard := getTowerBoard()
		givenPosition := Vector3{0, 0, 2}
		gotObstructors := givenBoard.GetObstructorsBelow(givenPosition)

		expectedLen := 2
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 1})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 0, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
	t.Run("GetObstructorsBelow on tower board (on {0, 0, 0})", func(t *testing.T) {
		givenBoard := getTowerBoard()
		givenPosition := Vector3{0, 0, 0}
		gotObstructors := givenBoard.GetObstructorsBelow(givenPosition)

		expectedLen := 0
		expectedTile := TilesIterator{}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})

	t.Run("GetObstructorsBelow on t-tile board (on {1, 1, 2})", func(t *testing.T) {
		givenBoard := getTBoard()
		givenPosition := Vector3{1, 1, 2}
		gotObstructors := givenBoard.GetObstructorsBelow(givenPosition)

		expectedLen := 4
		expectedTile1 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{1, 1, 1})]
		expectedTile2 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 0, 0})]
		expectedTile3 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{0, 1, 0})]
		expectedTile4 := givenBoard.Tiles[vector3TotilesIteratorIndex(givenBoard.Size, Vector3{2, 2, 0})]
		expectedTile := TilesIterator{expectedTile1, expectedTile2, expectedTile3, expectedTile4}

		assert.Equal(t, expectedLen, len(gotObstructors))
		assert.Equal(t, expectedTile, gotObstructors)
	})
}
