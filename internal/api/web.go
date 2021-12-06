package api

import (
	"fmt"
	"net/http"

	"github.com/cchaiyatad/mss/internal/utils"
)

type APIController struct {
	args *utils.Args
}

func CreateAPIController() *APIController {
	return &APIController{args: utils.GetArgs()}
}

func (con *APIController) StartServer() error {
	fmt.Printf("Listening on port %s\n", con.args.Port)

	server := &http.Server{
		Addr:    getPort(con.args),
		Handler: getServerMux(),
	}

	return server.ListenAndServe()
}

func getPort(args *utils.Args) string {
	return fmt.Sprintf(":%s", args.Port)
}

func getServerMux() *http.ServeMux {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/layout", layout)
	serverMux.HandleFunc("/solve", solve)
	return serverMux
}
