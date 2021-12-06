package api

import (
	"fmt"
	"net/http"

	"github.com/cchaiyatad/mss/internal/utils"
)

type APIController struct {
	Args *utils.Args
}

func CreateAPIController(args *utils.Args) *APIController {
	return &APIController{Args: args}
}

func (*APIController) StartServer(args *utils.Args) error {
	fmt.Printf("Listening on port %s\n", args.Port)

	server := &http.Server{
		Addr:    getPort(args),
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
