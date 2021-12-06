package api

import (
	l "github.com/cchaiyatad/mss/internal/layout"
	"github.com/cchaiyatad/mss/internal/utils"

	"net/http"
)

func layout(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write(getLayoutsOptionJSON())
}

func getLayoutsOptionJSON() []byte {
	return utils.ToJSON(l.GetLayoutOption())
}
