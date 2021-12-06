package layout

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"

	"github.com/cchaiyatad/mss/internal/file"
)

type Size struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
	Z int `xml:"z,attr"`
}
type Tile struct {
	X uint8 `xml:"x,attr"`
	Y uint8 `xml:"y,attr"`
	Z uint8 `xml:"z,attr"`
}

type Tiles struct {
	Tile []Tile `xml:"tile"`
}

type Layout struct {
	XMLName xml.Name `xml:"layout"`
	Name    string   `xml:"name"`
	Size    Size     `xml:"size"`
	Tiles   Tiles    `xml:"tiles"`
}

const key = "LAYOUT_PATH"

var layout_path = ""

// layout_name => .xml_file_name
var layouts map[string]string

func getLayoutpath(layout string) string {
	return fmt.Sprintf("%s%s", layout_path, layout)
}

func init() {
	InitLayout()
}

func InitLayout() {
	layout_path = os.Getenv(key)
	layouts = make(map[string]string)
	filenames := file.GetAllFileNameinPath(layout_path)
	for _, filename := range filenames {
		layoutName, err := getLayoutnameFromFile(getLayoutpath(filename))
		if err == nil {
			layouts[layoutName] = filename
		}
	}
}

func getLayoutnameFromFile(filename string) (string, error) {
	layout, err := getLayoutFromFile(filename)
	if err != nil {
		return "", err
	}

	return layout.Name, nil
}

func getLayoutFromFile(filename string) (*Layout, error) {
	bytes, err := file.GetBytesFromFile(filename)
	if err != nil {
		return nil, err
	}

	var layout Layout
	err = xml.Unmarshal(bytes, &layout)

	if layout.Size.X >= 256 || layout.Size.Y >= 256 || layout.Size.Z >= 256 {
		return &Layout{}, errors.New("board can not have a size that larger than 256")
	}
	return &layout, err
}

func getLayout(layoutName string) (*Layout, error) {
	filename, ok := layouts[layoutName]
	if !ok {
		return nil, fmt.Errorf("error: layout %q not exist", layoutName)
	}
	return getLayoutFromFile(getLayoutpath(filename))
}

func IsLayoutValid(layoutName string) bool {
	_, ok := layouts[layoutName]
	return ok
}

func GetSize(layoutName string) ([3]uint8, error) {
	layout, err := getLayout(layoutName)

	if err != nil {
		return [3]uint8{}, err
	}
	return [3]uint8{uint8(layout.Size.X), uint8(layout.Size.Y), uint8(layout.Size.Z)}, nil
}

func GetHasTilesFunc(layoutName string) (func(x, y, z uint8) bool, error) {
	layout, err := getLayout(layoutName)
	if err != nil {
		return nil, err
	}

	tilePosition := make(map[string]struct{})
	for _, tile := range layout.Tiles.Tile {
		key := getVector3UniqueKey(tile.X, tile.Y, tile.Z)
		tilePosition[key] = struct{}{}
	}

	hasTilesFunc := func(x, y, z uint8) bool {
		key := getVector3UniqueKey(x, y, z)
		_, ok := tilePosition[key]
		return ok
	}

	return hasTilesFunc, nil
}

func GetLayoutOption() []string {
	layoutNames := make([]string, 0, len(layouts))
	for layoutName := range layouts {
		layoutNames = append(layoutNames, layoutName)
	}

	return layoutNames
}

func getVector3UniqueKey(x, y, z uint8) string {
	return fmt.Sprintf("%d_%d_%d", x, y, z)
}
