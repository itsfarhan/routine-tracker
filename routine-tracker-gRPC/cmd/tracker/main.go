package main

import (
	"os"

	"github.com/itsfarhan/routine-tracker/internal/repository"
	"github.com/itsfarhan/routine-tracker/internal/server"
	"github.com/itsfarhan/routine-tracker/log"
)

const port = 28710

func main() {
	lgr := log.New(os.Stdout)
	repo := repository.New(lgr)  // create the in-memory store
	srv := server.New(repo, lgr) // inject it into the server
	if err := srv.ListenAndServe(port); err != nil {
		lgr.Logf("server error: %s", err)
		os.Exit(1)
	}
}
