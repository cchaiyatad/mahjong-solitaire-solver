package utils

import (
	"flag"
	"os"
)

const (
	PORT_KEY = "port"
)

type Args struct {
	Port string
}

func GetArgs() *Args {
	flagPort := flag.String(PORT_KEY, "", "a port to run api server (if not provide a env PORT will be use, if both env and flag not provide 8080 will be use")
	flag.Parse()

	if *flagPort != "" {
		return &Args{Port: *flagPort}
	}

	envPort := os.Getenv("PORT")
	if envPort != "" {
		return &Args{Port: envPort}
	}

	return &Args{Port: "8080"}
}
