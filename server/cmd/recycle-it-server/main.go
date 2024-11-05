package main

import (
	"log/slog"

	"github.com/DaikoneKisu/recycle-it/server"
)

func main() {
	err := server.Serve()
	if err != nil {
		slog.Error(err.Error())
	}
}
