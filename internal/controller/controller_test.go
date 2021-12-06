package controller

import (
	"os"
	"testing"

	"github.com/cchaiyatad/mss/internal/layout"

	"github.com/go-playground/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("LAYOUT_PATH", "./layout_test/")
	layout.InitLayout()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCreateController(t *testing.T) {
	t.Run("createController ", func(t *testing.T) {
		givenLayout := "layout"
		givenStrategy := "strategy"
		givenHeuristic := "heuristic"

		gotController := CreateController(givenLayout, givenStrategy, givenHeuristic)

		expectedLayout := "layout"
		expectedStrategy := "strategy"
		expectedHeuristic := "heuristic"

		assert.Equal(t, expectedLayout, gotController.layout)
		assert.Equal(t, expectedStrategy, gotController.strategy)
		assert.Equal(t, expectedHeuristic, gotController.heuristic)
	})
}

func TestStartController(t *testing.T) {
	t.Run("startController when everything is valid ", func(t *testing.T) {
		givenLayout := "two-tile"
		givenStrategy := "random"
		givenHeuristic := "random"

		givenController := &Controller{layout: givenLayout, strategy: givenStrategy, heuristic: givenHeuristic}

		gotBoard, gotOrder, gotOk := givenController.Start()

		expectedBoard := []byte(`{"tiles":[{"id":13,"face":34,"position":{"x":1,"y":2,"z":0}},{"id":15,"face":34,"position":{"x":3,"y":2,"z":0}}],"size":{"x_size":6,"y_size":6,"z_size":1},"layout":"two-tile"}`)
		expectedOrder := []byte(`[[13,15]]`)
		expectedOk := true

		assert.Equal(t, expectedBoard, gotBoard)
		assert.Equal(t, expectedOrder, gotOrder)
		assert.Equal(t, expectedOk, gotOk)
	})
	t.Run("startController when board has no tile ", func(t *testing.T) {
		givenLayout := "no-tile"
		givenStrategy := "random"
		givenHeuristic := "random"

		givenController := &Controller{layout: givenLayout, strategy: givenStrategy, heuristic: givenHeuristic}

		gotBoard, gotOrder, gotOk := givenController.Start()
		expectedBoard := []byte(`{"tiles":[],"size":{"x_size":8,"y_size":8,"z_size":2},"layout":"no-tile"}`)
		expectedOrder := []byte("[]")
		expectedOk := true

		assert.Equal(t, expectedBoard, gotBoard)
		assert.Equal(t, expectedOrder, gotOrder)
		assert.Equal(t, expectedOk, gotOk)
	})

	t.Run("startController when strategy not valid ", func(t *testing.T) {
		givenLayout := "two-tile"
		givenStrategy := "not-valid"
		givenHeuristic := "random"

		givenController := &Controller{layout: givenLayout, strategy: givenStrategy, heuristic: givenHeuristic}

		gotBoard, gotOrder, gotOk := givenController.Start()

		expectedBoard := []byte("null")
		expectedOrder := []byte("null")
		expectedOk := false

		assert.Equal(t, expectedBoard, gotBoard)
		assert.Equal(t, expectedOrder, gotOrder)
		assert.Equal(t, expectedOk, gotOk)
	})

	t.Run("startController when heuristic not valid ", func(t *testing.T) {
		givenLayout := "two-tile"
		givenStrategy := "random"
		givenHeuristic := "not-valid"

		givenController := &Controller{layout: givenLayout, strategy: givenStrategy, heuristic: givenHeuristic}

		gotBoard, gotOrder, gotOk := givenController.Start()

		expectedBoard := []byte("null")
		expectedOrder := []byte("null")
		expectedOk := false

		assert.Equal(t, expectedBoard, gotBoard)
		assert.Equal(t, expectedOrder, gotOrder)
		assert.Equal(t, expectedOk, gotOk)
	})
	t.Run("startController when layout not valid (not exist)", func(t *testing.T) {
		givenLayout := "not-valid"
		givenStrategy := "random"
		givenHeuristic := "random"

		givenController := &Controller{layout: givenLayout, strategy: givenStrategy, heuristic: givenHeuristic}

		gotBoard, gotOrder, gotOk := givenController.Start()

		expectedBoard := []byte("null")
		expectedOrder := []byte("null")
		expectedOk := false

		assert.Equal(t, expectedBoard, gotBoard)
		assert.Equal(t, expectedOrder, gotOrder)
		assert.Equal(t, expectedOk, gotOk)
	})
	t.Run("startController when layout not valid (have file but have only one tile) ", func(t *testing.T) {
		givenLayout := "one-tile"
		givenStrategy := "random"
		givenHeuristic := "random"

		givenController := &Controller{layout: givenLayout, strategy: givenStrategy, heuristic: givenHeuristic}

		gotBoard, gotOrder, gotOk := givenController.Start()

		expectedBoard := []byte("null")
		expectedOrder := []byte("null")
		expectedOk := false

		assert.Equal(t, expectedBoard, gotBoard)
		assert.Equal(t, expectedOrder, gotOrder)
		assert.Equal(t, expectedOk, gotOk)
	})
}
