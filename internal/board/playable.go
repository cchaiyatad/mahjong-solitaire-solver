package board

func (board *Board) HasPlayableTile() bool {
	return len(board.GetPlayableTiles()) != 0
}

func (board *Board) GetPlayableTiles() []*Tile {
	playableTiles := make([]*Tile, 0)
	for _, tile := range board.Tiles {
		if tile != nil && board.isTilePlayable(tile) {
			playableTiles = append(playableTiles, tile)
		}
	}

	return playableTiles
}

func (board *Board) isTilePlayable(tile *Tile) bool {
	if tile == nil || tile.isPicked {
		return false
	}

	topTiles := board.getTouchTileAbove(tile.Position)

	if len(topTiles) != 0 {
		return false
	}

	leftTiles := board.getTouchTileLeft(tile.Position)
	rightTiles := board.getTouchTileRight(tile.Position)

	if len(leftTiles) != 0 && len(rightTiles) != 0 {
		return false
	}
	return true
}
