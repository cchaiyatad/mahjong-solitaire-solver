package utils

import (
	"flag"
)

const (
	PORT_KEY = "port"
)

type Args struct {
	Port string
}

func GetArgs() *Args {
	port := flag.String(PORT_KEY, "8080", "a port to run api server")
	flag.Parse()
	return &Args{Port: *port}
}
