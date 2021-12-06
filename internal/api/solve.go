package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/cchaiyatad/mss/internal/controller"
	l "github.com/cchaiyatad/mss/internal/layout"
	"github.com/cchaiyatad/mss/internal/utils"
)

var validStrategy = map[string]struct{}{
	"random":        {},
	"multipleFirst": {},
}
var validHeuristic = map[string]struct{}{
	"random":   {},
	"maxBlock": {},
}

const respond_template = `{"board":%s, "order":%s, "params":%s}`

type Params struct {
	Strategy  string `json:"strategy"`
	Heuristic string `json:"heuristic"`
	Layout    string `json:"layout"`
}

func (params *Params) ToJSON() []byte {
	return utils.ToJSON(params)
}

func solve(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	params, ok := getParamsFromQuery(query)
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	setSeed()
	controller := controller.CreateController(params.Layout, params.Strategy, params.Heuristic)
	board, order, ok := controller.Start()

	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, respond_template, board, order, params.ToJSON())
}

func getParamsFromQuery(query url.Values) (*Params, bool) {
	layout, ok := getParamFromQuery(query, "layout")
	if !ok || !isLayoutValid(layout) {
		return nil, false
	}

	strategy, ok := getParamFromQuery(query, "strategy")
	if !ok || !isParamValueValid(strategy, validStrategy) {
		return nil, false
	}

	heuristic, ok := getParamFromQuery(query, "heuristic")
	if !ok || !isParamValueValid(heuristic, validHeuristic) {
		return nil, false
	}

	return &Params{Strategy: strategy,
		Heuristic: heuristic,
		Layout:    layout,
	}, true
}

func getParamFromQuery(query url.Values, key string) (string, bool) {
	value, ok := query[key]
	if !ok || len(value[0]) < 1 {
		return "", false
	}
	return value[0], true
}

func isLayoutValid(layoutName string) bool {
	return l.IsLayoutValid(layoutName)
}

func isParamValueValid(value string, validValues map[string]struct{}) bool {
	if _, ok := validValues[value]; !ok {
		return false
	}
	return true
}

func setSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}
