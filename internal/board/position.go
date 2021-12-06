package board

type Size struct {
	XSize uint8 `json:"x_size"`
	YSize uint8 `json:"y_size"`
	ZSize uint8 `json:"z_size"`
}

func uint8ArrayToSize(array [3]uint8) Size {
	return Size{array[0], array[1], array[2]}
}

func (s Size) getFlattenSize() int {
	return int(s.XSize) * int(s.YSize) * int(s.ZSize)
}
func (s Size) ToUint8() (uint8, uint8, uint8) {
	return s.XSize, s.YSize, s.ZSize
}
func (s Size) ToInt() (int, int, int) {
	return int(s.XSize), int(s.YSize), int(s.ZSize)
}

type Vector3 struct {
	X uint8 `json:"x"` // Left is -; Right is +
	Y uint8 `json:"y"` // Up is	-; Down is +
	Z uint8 `json:"z"` // Above Below; if Z == 0 mean the bottom layer
}

func (vector Vector3) ToUint8() (uint8, uint8, uint8) {
	return vector.X, vector.Y, vector.Z
}
func (vector Vector3) ToInt() (int, int, int) {
	return int(vector.X), int(vector.Y), int(vector.Z)
}

func (vector Vector3) getHasTileCheckPosition() []Vector3 {
	x, y, z := vector.ToUint8()
	return []Vector3{
		vector,
		{x - 1, y, z},
		{x, y - 1, z},
		{x - 1, y - 1, z},
	}
}

func (vector Vector3) getTouchPositionLeft() []Vector3 {
	x, y, z := vector.ToUint8()
	return []Vector3{
		{x - 2, y, z},
		{x - 2, y - 1, z},
		{x - 2, y + 1, z},
	}
}

func (vector Vector3) getTouchPositionRight() []Vector3 {
	x, y, z := vector.ToUint8()
	return []Vector3{
		{x + 2, y, z},
		{x + 2, y - 1, z},
		{x + 2, y + 1, z},
	}

}

func (vector Vector3) getTouchPositionAbove() []Vector3 {
	x, y, z := vector.ToUint8()
	return []Vector3{
		{x - 1, y - 1, z + 1},
		{x, y - 1, z + 1},
		{x + 1, y - 1, z + 1},
		{x - 1, y, z + 1},
		{x, y, z + 1},
		{x + 1, y, z + 1},
		{x - 1, y + 1, z + 1},
		{x, y + 1, z + 1},
		{x + 1, y + 1, z + 1},
	}
}

func (vector Vector3) getTouchPositionBelow() []Vector3 {
	x, y, z := vector.ToUint8()
	return []Vector3{
		{x - 1, y - 1, z - 1},
		{x, y - 1, z - 1},
		{x + 1, y - 1, z - 1},
		{x - 1, y, z - 1},
		{x, y, z - 1},
		{x + 1, y, z - 1},
		{x - 1, y + 1, z - 1},
		{x, y + 1, z - 1},
		{x + 1, y + 1, z - 1},
	}
}
