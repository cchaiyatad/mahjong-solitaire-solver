package board

import (
	"fmt"

	l "github.com/cchaiyatad/mss/internal/layout"
	"github.com/cchaiyatad/mss/internal/utils"
)

type Board struct {
	Tiles  TilesIterator `json:"tiles"`
	Size   Size          `json:"size"`
	Layout string        `json:"layout"`
}

func CreateBoard(layout string) (*Board, error) {
	arraySize, err := l.GetSize(layout)
	if err != nil {
		return nil, fmt.Errorf("error: can not create board: %s", err)
	}
	size := uint8ArrayToSize(arraySize)

	board := Board{Size: size, Layout: layout}
	board.ResetTiles()

	return &board, nil
}

func (board *Board) ResetTiles() {
	board.createEmptyTiles()
	board.createTileWithLayout()
}

func (board *Board) createEmptyTiles() {
	board.Tiles = make(TilesIterator, board.Size.getFlattenSize())
}

func (board *Board) createTileWithLayout() {
	isHasTile, _ := l.GetHasTilesFunc(board.Layout)

	for idx := 0; idx < board.Size.getFlattenSize(); idx++ {
		position := tilesIteratorIndexToVector3(board.Size, idx)
		x, y, z := position.ToUint8()

		if isHasTile(x, y, z) {
			board.Tiles[idx] = &Tile{Id: idx, Position: position}
		}
	}
}

func (board *Board) IsSolved() bool {
	if board == nil {
		return false
	}

	for _, tile := range board.Tiles {
		if tile == nil {
			continue
		}
		if !tile.isPicked {
			return false
		}
	}
	return true
}

func (board *Board) ToJSON() []byte {
	if board == nil {
		return []byte("null")
	}

	tiles := make(TilesIterator, 0)

	for _, tile := range board.Tiles {
		if tile != nil {
			tiles = append(tiles, tile)
		}
	}

	return utils.ToJSON(Board{
		Tiles:  tiles,
		Size:   board.Size,
		Layout: board.Layout,
	})
}

func (board *Board) pickTile(tile *Tile) int {
	if tile == nil {
		return -1
	}

	if tile == board.Tiles[tile.Id] && !tile.isPicked {
		tile.isPicked = true
		return tile.Id
	}

	return -1
}

func (board *Board) getTileAt(position Vector3) *Tile {
	for _, pos := range position.getHasTileCheckPosition() {
		if tile := board.getTileByTopLeftAt(pos); tile != nil {
			return tile
		}
	}
	return nil
}

func (board *Board) getTileByTopLeftAt(position Vector3) *Tile {
	if board == nil {
		return nil
	}

	sizeX, sizeY, sizeZ := board.Size.ToUint8()
	x, y, z := position.ToUint8()

	if sizeX <= x || sizeY <= y || sizeZ <= z {
		return nil
	}
	idx := vector3TotilesIteratorIndex(board.Size, position)
	return board.Tiles[idx]
}

func (board *Board) getTouchNotPickedTileWithPositions(positions []Vector3) TilesIterator {
	tiles := make(TilesIterator, 0)
	for _, pos := range positions {
		if tile := board.getTileByTopLeftAt(pos); tile != nil && !tile.isPicked {
			tiles = append(tiles, tile)
		}
	}
	return tiles
}

func (board *Board) getTouchTileLeft(position Vector3) TilesIterator {
	return board.getTouchNotPickedTileWithPositions(position.getTouchPositionLeft())
}

func (board *Board) getTouchTileRight(position Vector3) TilesIterator {
	return board.getTouchNotPickedTileWithPositions(position.getTouchPositionRight())
}

func (board *Board) getTouchTileAbove(position Vector3) TilesIterator {
	return board.getTouchNotPickedTileWithPositions(position.getTouchPositionAbove())
}

func (board *Board) getTouchTileBelow(position Vector3) TilesIterator {
	return board.getTouchNotPickedTileWithPositions(position.getTouchPositionBelow())
}

func (board *Board) getObstructors(position Vector3, fun func(Vector3) TilesIterator) TilesIterator {
	queue := fun(position)
	idx := 0

	checkMap := make(map[*Tile]struct{})
	tiles := make(TilesIterator, 0)

	for idx < len(queue) {
		currentTile := queue[idx]
		if _, ok := checkMap[currentTile]; !ok {
			tiles = append(tiles, currentTile)
			checkMap[currentTile] = struct{}{}
		}

		queue = append(queue, fun(currentTile.Position)...)
		idx += 1
	}
	return tiles
}

func (board *Board) GetObstructorsLeft(position Vector3) TilesIterator {
	return board.getObstructors(position, board.getTouchTileLeft)
}

func (board *Board) GetObstructorsRight(position Vector3) TilesIterator {
	return board.getObstructors(position, board.getTouchTileRight)
}

func (board *Board) GetObstructorsAbove(position Vector3) TilesIterator {
	return board.getObstructors(position, board.getTouchTileAbove)
}

func (board *Board) GetObstructorsBelow(position Vector3) TilesIterator {
	return board.getObstructors(position, board.getTouchTileBelow)
}
