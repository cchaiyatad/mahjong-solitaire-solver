package controller

import (
	"github.com/cchaiyatad/mss/internal/board"
	"github.com/cchaiyatad/mss/internal/solver"
	"github.com/cchaiyatad/mss/internal/utils"
)

type orders [][]int

type Controller struct {
	Solver    *solver.Solver
	Board     *board.Board
	layout    string
	strategy  string
	heuristic string
}

func CreateController(layout, strategy, heuristic string) *Controller {
	return &Controller{
		layout:    layout,
		strategy:  strategy,
		heuristic: heuristic,
	}
}

func (ctrl *Controller) Start() ([]byte, []byte, bool) {
	err := ctrl.initSolverAndBoard()
	if err != nil {
		return []byte("null"), []byte("null"), false
	}

	startBoardBytes, err := ctrl.randomSolvableFaceAndSave()
	if err != nil {
		return []byte("null"), []byte("null"), false
	}

	orderList := make(orders, 0)

	saveOrder := func(pickPairIDs []int) {
		if len(pickPairIDs) == 2 && pickPairIDs[0] != -1 && pickPairIDs[1] != -1 {
			orderList = append(orderList, pickPairIDs)
		}
	}

	ctrl.solve(saveOrder)
	return startBoardBytes, orderList.ToJSON(), true
}

func (ctrl *Controller) initSolverAndBoard() error {
	err := ctrl.initialSolver()
	if err != nil {
		return err
	}
	return ctrl.initialBoard()
}

func (ctrl *Controller) initialSolver() error {
	var err error
	ctrl.Solver, err = solver.CreateSolver(ctrl.strategy, ctrl.heuristic)
	return err
}

func (ctrl *Controller) initialBoard() error {
	var err error
	ctrl.Board, err = board.CreateBoard(ctrl.layout)
	return err
}

func (ctrl *Controller) randomSolvableFaceAndSave() ([]byte, error) {
	err := ctrl.Board.RandomSolvableFace()
	if err != nil {
		return []byte{}, err
	}
	return ctrl.Board.ToJSON(), nil
}

func (ctrl *Controller) solve(callBack func([]int)) {
	ctrl.Solver.Solve(ctrl.Board, callBack)
}

func (ordersList orders) ToJSON() []byte {
	return utils.ToJSON(ordersList)
}
