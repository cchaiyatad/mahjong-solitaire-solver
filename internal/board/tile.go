package board

// 0 is not used
// [1-38] represent every faces in mahjong tile
// 春 is the same as 梅 in the same way with 夏秋冬 and 兰竹菊
type Face uint8

const MAX_FACE = 38

// each tile take 2*2*1 space
type Tile struct {
	Id       int     `json:"id"`
	Face     Face    `json:"face"`
	Position Vector3 `json:"position"`
	isPicked bool
}

func getFacesInSuits() []Face {
	faces := make([]Face, MAX_FACE*2)
	for idx := range faces {
		faces[idx] = Face(idx/2) + 1
	}

	return faces
}
