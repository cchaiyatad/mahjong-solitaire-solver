package main

import (
	"log"

	"github.com/cchaiyatad/mss/internal/api"
)

func main() {
	controller := api.CreateAPIController()
	log.Fatalln(controller.StartServer())
}
