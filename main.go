package main

import (
	"github.com/mpedrozoduran/go-orchestrator/internal/runner"
)

func main() {
	srv := runner.NewServiceRestAPI()
	srv.Run()
}
