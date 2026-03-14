package main

import (
	"fmt"
	"os"
	"routine-tracker/routine-tracker-gRPC/internal/server"
	"routine-tracker/routine-tracker-gRPC/log"
)

const port = 28710

func main() {
	lgr := log.New(os.Stdout)
	srv := server.New(lgr)
	err := srv.ListenAndServe(port)
	if err != nil {
		lgr.Logf("Error while running the server: %s", err)
		os.Exit(1)
	}
}
