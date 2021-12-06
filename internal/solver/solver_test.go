package solver

import (
	"testing"

	"github.com/cchaiyatad/mss/internal/board"

	"github.com/go-playground/assert"
)

func TestCreateSolver(t *testing.T) {
	t.Run("CreateSolver with Random Strategy and Random Heuristic", func(t *testing.T) {
		givenStrategy := "random"
		givenHeuristic := "random"

		gotSolver, gotError := CreateSolver(givenStrategy, givenHeuristic)
		_, isStrategyOk := gotSolver.Strategy.(*RandomStrategy)
		_, isHeuristicOk := gotSolver.Heuristic.(*RandomHeuristic)

		assert.Equal(t, nil, gotError)
		assert.Equal(t, true, isStrategyOk)
		assert.Equal(t, true, isHeuristicOk)
	})
	t.Run("CreateSolver with multipleFirst Strategy and Random Heuristic", func(t *testing.T) {
		givenStrategy := "multipleFirst"
		givenHeuristic := "random"

		gotSolver, gotError := CreateSolver(givenStrategy, givenHeuristic)
		_, isStrategyOk := gotSolver.Strategy.(*MultipleFirst)
		_, isHeuristicOk := gotSolver.Heuristic.(*RandomHeuristic)

		assert.Equal(t, nil, gotError)
		assert.Equal(t, true, isStrategyOk)
		assert.Equal(t, true, isHeuristicOk)
	})
	t.Run("CreateSolver with Random Strategy and maxBlock Heuristic", func(t *testing.T) {
		givenStrategy := "random"
		givenHeuristic := "maxBlock"

		gotSolver, gotError := CreateSolver(givenStrategy, givenHeuristic)
		_, isStrategyOk := gotSolver.Strategy.(*RandomStrategy)
		_, isHeuristicOk := gotSolver.Heuristic.(*MaxBlock)

		assert.Equal(t, nil, gotError)
		assert.Equal(t, true, isStrategyOk)
		assert.Equal(t, true, isHeuristicOk)
	})
	t.Run("CreateSolver with multipleFirst Strategy and maxBlock Heuristic", func(t *testing.T) {
		givenStrategy := "multipleFirst"
		givenHeuristic := "maxBlock"

		gotSolver, gotError := CreateSolver(givenStrategy, givenHeuristic)
		_, isStrategyOk := gotSolver.Strategy.(*MultipleFirst)
		_, isHeuristicOk := gotSolver.Heuristic.(*MaxBlock)

		assert.Equal(t, nil, gotError)
		assert.Equal(t, true, isStrategyOk)
		assert.Equal(t, true, isHeuristicOk)
	})
	t.Run("CreateSolver with non-exist Strategy and maxBlock Heuristic", func(t *testing.T) {
		givenStrategy := "non-exist"
		givenHeuristic := "maxBlock"

		gotSolver, gotError := CreateSolver(givenStrategy, givenHeuristic)
		expectError := "error: can not create slover strategy: non-exist heuristic: maxBlock"
		assert.Equal(t, expectError, gotError.Error())
		assert.Equal(t, nil, gotSolver)
	})
	t.Run("CreateSolver with random Strategy and non-exist Heuristic", func(t *testing.T) {
		givenStrategy := "random"
		givenHeuristic := "non-exist"

		gotSolver, gotError := CreateSolver(givenStrategy, givenHeuristic)
		expectError := "error: can not create slover strategy: random heuristic: non-exist"
		assert.Equal(t, expectError, gotError.Error())
		assert.Equal(t, nil, gotSolver)
	})
}

func TestMaxBlockHor(t *testing.T) {
	t.Run("getH 4-left 7-right (follow example in paper)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneRowBoard()
		givetTilePosition := board.Vector3{X: 8, Y: 0, Z: 0}

		expectedH := 1
		assert.Equal(t, expectedH, givenHeuristic.hor(givenBoard, givetTilePosition))
	})

	t.Run("getH 0-left 11-right", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneRowBoard()
		givetTilePosition := board.Vector3{X: 0, Y: 0, Z: 0}

		expectedH := 5
		assert.Equal(t, expectedH, givenHeuristic.hor(givenBoard, givetTilePosition))
	})

	t.Run("getH 11-left 0-right", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneRowBoard()
		givetTilePosition := board.Vector3{X: 22, Y: 0, Z: 0}

		expectedH := 5
		assert.Equal(t, expectedH, givenHeuristic.hor(givenBoard, givetTilePosition))
	})

	t.Run("getH one tile board", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneTileBoard()
		givetTilePosition := board.Vector3{X: 2, Y: 2, Z: 0}

		expectedH := 0
		assert.Equal(t, expectedH, givenHeuristic.hor(givenBoard, givetTilePosition))
	})
	t.Run("getH getOneTwoOneTwoOneBoard (middle tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneTwoOneTwoOneBoard()
		givetTilePosition := board.Vector3{X: 4, Y: 1, Z: 0}

		expectedH := 0
		assert.Equal(t, expectedH, givenHeuristic.hor(givenBoard, givetTilePosition))
	})

	t.Run("getH getOneTwoOneTwoOneBoard (most left tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneTwoOneTwoOneBoard()
		givetTilePosition := board.Vector3{X: 0, Y: 1, Z: 0}

		expectedH := 3
		assert.Equal(t, expectedH, givenHeuristic.hor(givenBoard, givetTilePosition))
	})
	t.Run("getH getOneTwoOneTwoOneBoard (most right tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneTwoOneTwoOneBoard()
		givetTilePosition := board.Vector3{X: 8, Y: 1, Z: 0}

		expectedH := 3
		assert.Equal(t, expectedH, givenHeuristic.hor(givenBoard, givetTilePosition))
	})
	t.Run("getH getOneTwoOneTwoOneBoard ({2,0,0} tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneTwoOneTwoOneBoard()
		givetTilePosition := board.Vector3{X: 2, Y: 0, Z: 0}

		expectedH := 1
		assert.Equal(t, expectedH, givenHeuristic.hor(givenBoard, givetTilePosition))
	})
	t.Run("getH getOneTwoOneTwoOneBoard ({2,2,0} tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneTwoOneTwoOneBoard()
		givetTilePosition := board.Vector3{X: 2, Y: 2, Z: 0}

		expectedH := 1
		assert.Equal(t, expectedH, givenHeuristic.hor(givenBoard, givetTilePosition))
	})
}

func TestMaxBlockVer(t *testing.T) {
	t.Run("getVer tower layout board (most top tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getTowerBoard()
		givetTilePosition := board.Vector3{X: 0, Y: 0, Z: 2}

		expectedH := 2
		assert.Equal(t, expectedH, givenHeuristic.ver(givenBoard, givetTilePosition))
	})
	t.Run("getVer tower layout board (middle tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getTowerBoard()
		givetTilePosition := board.Vector3{X: 0, Y: 0, Z: 1}

		expectedH := 1
		assert.Equal(t, expectedH, givenHeuristic.ver(givenBoard, givetTilePosition))
	})
	t.Run("getVer tower layout board (bottom tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getTowerBoard()
		givetTilePosition := board.Vector3{X: 0, Y: 0, Z: 0}

		expectedH := 0
		assert.Equal(t, expectedH, givenHeuristic.ver(givenBoard, givetTilePosition))
	})
	t.Run("getVer getOneTwoThreeVerBoard (most top tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneTwoThreeVerBoard()
		givetTilePosition := board.Vector3{X: 2, Y: 0, Z: 2}

		expectedH := 7
		assert.Equal(t, expectedH, givenHeuristic.ver(givenBoard, givetTilePosition))
	})
	t.Run("getVer getOneTwoThreeVerBoard ({1,0,1} tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneTwoThreeVerBoard()
		givetTilePosition := board.Vector3{X: 1, Y: 0, Z: 1}

		expectedH := 3
		assert.Equal(t, expectedH, givenHeuristic.ver(givenBoard, givetTilePosition))
	})

	t.Run("getVer getOneTwoThreeVerBoard ({3,0,1} tile)", func(t *testing.T) {
		givenHeuristic := &MaxBlock{}
		givenBoard := getOneTwoThreeVerBoard()
		givetTilePosition := board.Vector3{X: 3, Y: 0, Z: 1}

		expectedH := 3
		assert.Equal(t, expectedH, givenHeuristic.ver(givenBoard, givetTilePosition))
	})
}

func TestSortDecsByHeuristic(t *testing.T) {
	t.Run("SortDecsByHeuristic", func(t *testing.T) {
		givenHeuristic := &idHeuristic{}
		givenBoard := getHBoard()
		givenPairs := givenBoard.FindPossiblePair()

		givenTile1 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 1, Y: 1, Z: 0})]
		givenTile2 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 5, Y: 1, Z: 0})]
		_ = givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 3, Y: 2, Z: 0})]
		givenTile4 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 1, Y: 3, Z: 0})]
		givenTile5 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 5, Y: 3, Z: 0})]

		// note that board.Tiles is store in [z][y][x] so the higher the y, the higher the id
		sortDecsByHeuristic(givenHeuristic, givenBoard, givenPairs)
		expectedLen := 6
		expectedPair1 := board.Pair{givenTile4, givenTile5}
		expectedPair2 := board.Pair{givenTile2, givenTile5}
		expectedPair3 := board.Pair{givenTile1, givenTile5}
		expectedPair4 := board.Pair{givenTile2, givenTile4}
		expectedPair5 := board.Pair{givenTile1, givenTile4}
		expectedPair6 := board.Pair{givenTile1, givenTile2}

		expectedPairs := board.PairIterator{expectedPair1, expectedPair2, expectedPair3, expectedPair4, expectedPair5, expectedPair6}

		assert.Equal(t, expectedLen, len(givenPairs))
		assert.Equal(t, expectedPairs, givenPairs)
	})
}

func TestRandomStrategy(t *testing.T) {
	t.Run("TestRandomStrategy using idHeuristic (all have same face)", func(t *testing.T) {
		givenStrategy := &RandomStrategy{}
		givenHeuristic := &idHeuristic{}
		givenBoard := getHBoard()

		givenTile1 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 1, Y: 1, Z: 0})]
		givenTile2 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 5, Y: 1, Z: 0})]
		_ = givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 3, Y: 2, Z: 0})]
		givenTile4 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 1, Y: 3, Z: 0})]
		givenTile5 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 5, Y: 3, Z: 0})]

		gotOrderList := make([][]int, 0)

		givenSaveOrderCallback := func(pickPairIDs []int) {
			if len(pickPairIDs) == 2 && pickPairIDs[0] != -1 && pickPairIDs[1] != -1 {
				gotOrderList = append(gotOrderList, pickPairIDs)
			}
		}

		givenStrategy.Solve(givenBoard, givenHeuristic, givenSaveOrderCallback)

		expectedLen := 2
		expectedOrder1 := []int{
			givenTile4.Id,
			givenTile5.Id,
		}
		expectedOrder2 := []int{
			givenTile1.Id,
			givenTile2.Id,
		}
		expectedPairs := [][]int{expectedOrder1, expectedOrder2}

		assert.Equal(t, expectedLen, len(gotOrderList))
		assert.Equal(t, expectedPairs, gotOrderList)
	})

	t.Run("TestRandomStrategy using idHeuristic", func(t *testing.T) {
		givenStrategy := &RandomStrategy{}
		givenHeuristic := &idHeuristic{}
		givenBoard := getHBoard()

		_ = givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 1, Y: 1, Z: 0})]
		givenTile2 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 5, Y: 1, Z: 0})]
		givenTile3 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 3, Y: 2, Z: 0})]
		givenTile4 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 1, Y: 3, Z: 0})]
		givenTile5 := givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 5, Y: 3, Z: 0})]

		givenTile2.Face = 1
		givenTile5.Face = 1

		gotOrderList := make([][]int, 0)

		givenSaveOrderCallback := func(pickPairIDs []int) {
			if len(pickPairIDs) == 2 && pickPairIDs[0] != -1 && pickPairIDs[1] != -1 {
				gotOrderList = append(gotOrderList, pickPairIDs)
			}
		}

		givenStrategy.Solve(givenBoard, givenHeuristic, givenSaveOrderCallback)

		expectedLen := 2
		expectedOrder1 := []int{
			givenTile2.Id,
			givenTile5.Id,
		}
		expectedOrder2 := []int{
			givenTile3.Id,
			givenTile4.Id,
		}

		expectedPairs := [][]int{expectedOrder1, expectedOrder2}

		assert.Equal(t, expectedLen, len(gotOrderList))
		assert.Equal(t, expectedPairs, gotOrderList)
	})
}

func TestMultipleFirstStrategy(t *testing.T) {
	t.Run("TestMultipleFirst using idHeuristic", func(t *testing.T) {
		// 1 3 4  1
		// 5 6 7  2
		// 1 2 2  8
		// 1 9 10 2

		// MutltipleFirst strategy should pick face = 1 first (there are 4 playable tiles)
		// then 2 (1,2) and (3,3) (there are 3 playable tiles)
		// then 2 (2,2) and (3,1) (there are 2 playable tiles)

		givenStrategy := &MultipleFirst{}
		givenHeuristic := &idHeuristic{}
		givenBoard := getSixteenBoard()
		givenTiles := []*board.Tile{
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 0, Y: 0, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 2, Y: 0, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 4, Y: 0, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 6, Y: 0, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 0, Y: 2, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 2, Y: 2, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 4, Y: 2, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 6, Y: 2, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 0, Y: 4, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 2, Y: 4, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 4, Y: 4, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 6, Y: 4, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 0, Y: 6, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 2, Y: 6, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 4, Y: 6, Z: 0})],
			givenBoard.Tiles[getIdx(givenBoard, board.Vector3{X: 6, Y: 6, Z: 0})],
		}

		givenTiles[0].Face = 1
		givenTiles[1].Face = 3
		givenTiles[2].Face = 4
		givenTiles[3].Face = 1

		givenTiles[4].Face = 5
		givenTiles[5].Face = 6
		givenTiles[6].Face = 7
		givenTiles[7].Face = 2

		givenTiles[8].Face = 1
		givenTiles[9].Face = 2
		givenTiles[10].Face = 2
		givenTiles[11].Face = 8

		givenTiles[12].Face = 1
		givenTiles[13].Face = 9
		givenTiles[14].Face = 10
		givenTiles[15].Face = 2

		gotOrderList := make([][]int, 0)

		givenSaveOrderCallback := func(pickPairIDs []int) {
			if len(pickPairIDs) == 2 && pickPairIDs[0] != -1 && pickPairIDs[1] != -1 {
				gotOrderList = append(gotOrderList, pickPairIDs)
			}
		}

		givenStrategy.Solve(givenBoard, givenHeuristic, givenSaveOrderCallback)

		expectedLen := 4
		expectedOrder1 := []int{
			givenTiles[8].Id,
			givenTiles[12].Id,
		}
		expectedOrder2 := []int{
			givenTiles[0].Id,
			givenTiles[3].Id,
		}
		expectedOrder3 := []int{
			givenTiles[9].Id,
			givenTiles[15].Id,
		}
		expectedOrder4 := []int{
			givenTiles[7].Id,
			givenTiles[10].Id,
		}

		expectedPairs := [][]int{expectedOrder1, expectedOrder2, expectedOrder3, expectedOrder4}

		assert.Equal(t, expectedLen, len(gotOrderList))
		assert.Equal(t, expectedPairs, gotOrderList)

	})
}
