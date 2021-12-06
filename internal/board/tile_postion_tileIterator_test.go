package board

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestTilesIteratorIndexToVector3(t *testing.T) {
	cases := []struct {
		GivenSize  Size
		GivenIndex int
		Expected   Vector3
	}{
		{Size{3, 3, 3}, 0, Vector3{0, 0, 0}},
		{Size{3, 3, 3}, 1, Vector3{1, 0, 0}},
		{Size{3, 3, 3}, 4, Vector3{1, 1, 0}},
		{Size{3, 3, 3}, 10, Vector3{1, 0, 1}},
		{Size{3, 3, 3}, 13, Vector3{1, 1, 1}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("tilesIteratorIndexToVector3(%v, %d) should get %v", tc.GivenSize, tc.GivenIndex, tc.Expected), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, tilesIteratorIndexToVector3(tc.GivenSize, tc.GivenIndex))
		})
	}
}

func TestVector3TotilesIteratorIndex(t *testing.T) {
	cases := []struct {
		GivenSize    Size
		GivenVector3 Vector3
		Expected     int
	}{
		{Size{3, 3, 3}, Vector3{0, 0, 0}, 0},
		{Size{3, 3, 3}, Vector3{1, 0, 0}, 1},
		{Size{3, 3, 3}, Vector3{1, 1, 0}, 4},
		{Size{3, 3, 3}, Vector3{1, 0, 1}, 10},
		{Size{3, 3, 3}, Vector3{1, 1, 1}, 13},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("vector3TotilesIteratorIndex(%v, %v) should get %d", tc.GivenSize, tc.GivenVector3, tc.Expected), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, vector3TotilesIteratorIndex(tc.GivenSize, tc.GivenVector3))
		})
	}
}

func TestGetShuffleFace(t *testing.T) {
	faces := getFacesInSuits()
	expected := []Face{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 23, 24, 24, 25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38}
	assert.Equal(t, expected, faces)
}

func TestGetFlatternSize(t *testing.T) {
	given := Size{10, 10, 2}
	expected := 200
	assert.Equal(t, expected, given.getFlattenSize())
}

func TestGetHasTileCheckPosition(t *testing.T) {
	cases := []struct {
		GivenVector Vector3
		Expected    []Vector3
	}{
		{Vector3{2, 2, 2}, []Vector3{
			{2, 2, 2},
			{1, 2, 2},
			{2, 1, 2},
			{1, 1, 2},
		}},
		{Vector3{0, 0, 0}, []Vector3{
			{0, 0, 0},
			{255, 0, 0},
			{0, 255, 0},
			{255, 255, 0},
		}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("getHasTileCheckPosition(%v) should get %v", tc.GivenVector, tc.Expected), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, tc.GivenVector.getHasTileCheckPosition())
		})
	}
}

func TestGetTouchPositionLeft(t *testing.T) {
	cases := []struct {
		GivenVector Vector3
		Expected    []Vector3
	}{
		{Vector3{2, 2, 2}, []Vector3{
			{0, 2, 2},
			{0, 1, 2},
			{0, 3, 2},
		}},
		{Vector3{0, 0, 0}, []Vector3{
			{254, 0, 0},
			{254, 255, 0},
			{254, 1, 0},
		}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("getGetTouchPositionLeft(%v) should get %v", tc.GivenVector, tc.Expected), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, tc.GivenVector.getTouchPositionLeft())
		})
	}
}
func TestGetTouchPositionRight(t *testing.T) {
	cases := []struct {
		GivenVector Vector3
		Expected    []Vector3
	}{
		{Vector3{2, 2, 2}, []Vector3{
			{4, 2, 2},
			{4, 1, 2},
			{4, 3, 2},
		}},
		{Vector3{0, 0, 0}, []Vector3{
			{2, 0, 0},
			{2, 255, 0},
			{2, 1, 0},
		}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("getGetTouchPositionRight(%v) should get %v", tc.GivenVector, tc.Expected), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, tc.GivenVector.getTouchPositionRight())
		})
	}
}
func TestGetTouchPositionAbove(t *testing.T) {
	cases := []struct {
		GivenVector Vector3
		Expected    []Vector3
	}{
		{Vector3{2, 2, 2}, []Vector3{
			{1, 1, 3},
			{2, 1, 3},
			{3, 1, 3},
			{1, 2, 3},
			{2, 2, 3},
			{3, 2, 3},
			{1, 3, 3},
			{2, 3, 3},
			{3, 3, 3},
		}},
		{Vector3{0, 0, 0}, []Vector3{
			{255, 255, 1},
			{0, 255, 1},
			{1, 255, 1},
			{255, 0, 1},
			{0, 0, 1},
			{1, 0, 1},
			{255, 1, 1},
			{0, 1, 1},
			{1, 1, 1},
		}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("getGetTouchPositionAbove(%v) should get %v", tc.GivenVector, tc.Expected), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, tc.GivenVector.getTouchPositionAbove())
		})
	}
}
func TestGetTouchPositionBelow(t *testing.T) {
	cases := []struct {
		GivenVector Vector3
		Expected    []Vector3
	}{
		{Vector3{2, 2, 2}, []Vector3{
			{1, 1, 1},
			{2, 1, 1},
			{3, 1, 1},
			{1, 2, 1},
			{2, 2, 1},
			{3, 2, 1},
			{1, 3, 1},
			{2, 3, 1},
			{3, 3, 1},
		}},
		{Vector3{0, 0, 0}, []Vector3{
			{255, 255, 255},
			{0, 255, 255},
			{1, 255, 255},
			{255, 0, 255},
			{0, 0, 255},
			{1, 0, 255},
			{255, 1, 255},
			{0, 1, 255},
			{1, 1, 255},
		}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("getGetTouchPositionBelow(%v) should get %v", tc.GivenVector, tc.Expected), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, tc.GivenVector.getTouchPositionBelow())
		})
	}
}
