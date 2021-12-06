package board

type TilesIterator []*Tile

func tilesIteratorIndexToVector3(size Size, idx int) Vector3 {
	sizeX, sizeY, _ := size.ToInt()

	z := idx / (sizeX * sizeY)
	y := (idx % (sizeX * sizeY)) / sizeX
	x := (idx % (sizeX * sizeY)) % sizeX
	return Vector3{X: uint8(x), Y: uint8(y), Z: uint8(z)}
}

func vector3TotilesIteratorIndex(size Size, vector Vector3) int {
	sizeX, sizeY, _ := size.ToInt()
	x, y, z := vector.ToInt()
	return (z * sizeX * sizeY) + (y * sizeX) + x
}
