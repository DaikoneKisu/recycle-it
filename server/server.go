package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/DaikoneKisu/recycle-it/server/internal/db"
	"github.com/DaikoneKisu/recycle-it/server/internal/game"
	"github.com/DaikoneKisu/recycle-it/server/internal/players"
	pbGame "github.com/DaikoneKisu/recycle-it/server/internal/protos/game"
	"google.golang.org/grpc"
)

func Serve() error {
	const SERVER_ADDRESS = "0.0.0.0:8080"
	networkListener, err := net.Listen("tcp", SERVER_ADDRESS)
	if err != nil {
		return fmt.Errorf("couldn't listen to the network: %w", err)
	}

	recycleItDB, err := db.NewConnection()
	if err != nil {
		return fmt.Errorf("couldn't connect to the database: %w", err)
	}
	db.RunMigrations(recycleItDB)

	grpcServer := grpc.NewServer()
	gameController := game.NewController(
		game.NewGameManager(
			recycleItDB,
			game.NewRepository(recycleItDB, players.NewRepository(recycleItDB)),
		),
	)
	pbGame.RegisterGameControllerServer(grpcServer, gameController)

	slog.Info(fmt.Sprintf("listening on %v", SERVER_ADDRESS))
	grpcServer.Serve(networkListener)

	return nil
}
