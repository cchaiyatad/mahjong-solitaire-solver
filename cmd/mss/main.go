package main

import (
	"log"

	"github.com/cchaiyatad/mss/internal/api"
	"github.com/cchaiyatad/mss/internal/utils"
)

func main() {
	controller := api.CreateAPIController(nil)
	log.Fatalln(controller.StartServer(utils.GetArgs()))
}
