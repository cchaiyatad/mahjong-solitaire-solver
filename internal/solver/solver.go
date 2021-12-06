package solver

import (
	"fmt"

	"github.com/cchaiyatad/mss/internal/board"
)

type Solver struct {
	Strategy  Strategy
	Heuristic Heuristic
}

func (solver Solver) Solve(board *board.Board, callBack func(pickedTileIDs []int)) {
	solver.Strategy.Solve(board, solver.Heuristic, callBack)
}

func CreateSolver(strategyName, heuristicName string) (*Solver, error) {
	strategy := getStrategy(strategyName)
	heuristic := getHeuristic(heuristicName)

	if strategy == nil || heuristic == nil {
		return nil, fmt.Errorf("error: can not create slover strategy: %s heuristic: %s", strategyName, heuristicName)
	}
	return &Solver{
		Strategy:  strategy,
		Heuristic: heuristic,
	}, nil
}

func getStrategy(strategy string) Strategy {
	switch strategy {
	case "random":
		return &RandomStrategy{}
	case "multipleFirst":
		return &MultipleFirst{}
	default:
		return nil
	}
}

func getHeuristic(heuristic string) Heuristic {
	switch heuristic {
	case "random":
		return &RandomHeuristic{}
	case "maxBlock":
		return &MaxBlock{}
	default:
		return nil
	}
}
