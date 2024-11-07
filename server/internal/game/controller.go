package game

import (
	"io"
	"time"

	"github.com/DaikoneKisu/recycle-it/server/internal/players"
	pb "github.com/DaikoneKisu/recycle-it/server/internal/protos/game"
	pbMath "github.com/DaikoneKisu/recycle-it/server/internal/protos/math"
	"google.golang.org/grpc"
)

type Controller struct {
	pb.UnimplementedGameControllerServer
	gameManager GameManager
}

func NewController(gameManager GameManager) Controller {
	return Controller{gameManager: gameManager}
}

func (c Controller) HostGame(
	request *pb.HostGameRequest,
	responseStream grpc.ServerStreamingServer[pb.HostGameResponse],
) error {
	lobbyUpdates, err := c.gameManager.HostGame(
		responseStream.Context(),
		request.HostNickname,
		Settings{
			RequiredPlayerAmount:  request.GameSettings.RequiredPlayerAmount,
			GameDurationInSeconds: request.GameSettings.GameDurationInSeconds,
		},
	)
	if err != nil {
		return err
	}

	var lobbyUpdate Lobby
updateLoop:
	for {
		select {
		case <-responseStream.Context().Done():
			break updateLoop
		case lobbyUpdate = <-lobbyUpdates:
			responseStream.Send(buildHostGameResponse(lobbyUpdate))
		}
		if lobbyUpdate.HasGameStarted {
			break updateLoop
		}
		const NUMBER_OF_UPDATES_PER_SECOND = 60
		time.Sleep(time.Second / NUMBER_OF_UPDATES_PER_SECOND)
	}

	return nil
}

func buildHostGameResponse(lobby Lobby) *pb.HostGameResponse {
	pbPlayers := make([]*pb.Player, len(lobby.Players))
	for i, p := range lobby.Players {
		pbPlayers[i] = &pb.Player{
			Nickname: p.Nickname,
			Role:     playerRoleToPB(p.Role),
		}
	}

	return &pb.HostGameResponse{
		GameID: lobby.GameID,
		GameSettings: &pb.GameSettings{
			RequiredPlayerAmount:  lobby.Settings.RequiredPlayerAmount,
			GameDurationInSeconds: lobby.Settings.GameDurationInSeconds,
		},
		Players: pbPlayers,
	}
}

func (c Controller) JoinGame(
	request *pb.JoinGameRequest,
	responseStream grpc.ServerStreamingServer[pb.JoinGameResponse],
) error {
	lobbyUpdates, err := c.gameManager.JoinGame(
		responseStream.Context(),
		request.GameID,
		request.GuestNickname,
	)
	if err != nil {
		return err
	}

	var lobbyUpdate Lobby
updateLoop:
	for {
		select {
		case <-responseStream.Context().Done():
			break updateLoop
		case lobbyUpdate = <-lobbyUpdates:
			responseStream.Send(buildJoinGameResponse(lobbyUpdate))
		}
		if lobbyUpdate.HasGameStarted {
			break updateLoop
		}
		const NUMBER_OF_UPDATES_PER_SECOND = 60
		time.Sleep(time.Second / NUMBER_OF_UPDATES_PER_SECOND)
	}

	return nil
}

func buildJoinGameResponse(lobby Lobby) *pb.JoinGameResponse {
	pbPlayers := make([]*pb.Player, len(lobby.Players))
	for i, p := range lobby.Players {
		pbPlayers[i] = &pb.Player{
			Nickname: p.Nickname,
			Role:     playerRoleToPB(p.Role),
		}
	}

	return &pb.JoinGameResponse{
		GameID: lobby.GameID,
		GameSettings: &pb.GameSettings{
			RequiredPlayerAmount:  lobby.Settings.RequiredPlayerAmount,
			GameDurationInSeconds: lobby.Settings.GameDurationInSeconds,
		},
		Players: pbPlayers,
	}
}

func (c Controller) PlayGame(communicationStream grpc.BidiStreamingServer[pb.PlayGameRequest, pb.PlayGameResponse]) error {
	for {
		request, err := communicationStream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		updatedGame, err := c.gameManager.MovePaddle(
			communicationStream.Context(),
			request.GameID,
			request.GuestNickname,
			Point2D{X: request.PaddleLocation.X, Y: request.PaddleLocation.Y},
		)
		if err != nil {
			return err
		}

		communicationStream.Send(buildPlayGameResponse(updatedGame))
	}
}

func buildPlayGameResponse(game Game) *pb.PlayGameResponse {
	pbGarbageCollectors := make([]*pb.GarbageCollector, len(game.Stage.GarbageCollectors))
	for i, gc := range game.Stage.GarbageCollectors {
		pbGarbageCollected := make([]pb.Garbage, len(gc.GarbageCollected))
		for j, g := range gc.GarbageCollected {
			pbGarbageCollected[j] = garbageToPB(g)
		}

		pbGarbageCollectors[i] = &pb.GarbageCollector{
			Player: &pb.Player{
				Nickname: gc.Player.Nickname,
				Role:     playerRoleToPB(gc.Player.Role),
			},
			GarbageToCollect: garbageToPB(gc.GarbageToCollect),
			GarbageCollected: pbGarbageCollected,
			PaddleLocation: &pbMath.Point2D{
				X: gc.PaddleLocation.X,
				Y: gc.PaddleLocation.Y,
			},
		}
	}

	return &pb.PlayGameResponse{
		Game: &pb.Game{
			Id: game.ID,
			Settings: &pb.GameSettings{
				RequiredPlayerAmount:  game.Settings.RequiredPlayerAmount,
				GameDurationInSeconds: game.Settings.GameDurationInSeconds,
			},
			Stage: &pb.GameStage{
				GarbageCollectors:  pbGarbageCollectors,
				UncollectedGarbage: garbageToPB(game.Stage.UncollectedGarbage),
				UncollectedGarbageLocation: &pbMath.Point2D{
					X: game.Stage.UncollectedGarbageLocation.X,
					Y: game.Stage.UncollectedGarbageLocation.Y,
				},
			},
			TimeRemainingInSeconds: game.TimeRemainingInSeconds,
		},
	}
}

func playerRoleToPB(role players.PlayerRole) pb.PlayerRole {
	switch role {
	case players.PLAYER_ROLE_HOST:
		return pb.PlayerRole_PLAYER_ROLE_HOST
	case players.PLAYER_ROLE_GUEST:
		return pb.PlayerRole_PLAYER_ROLE_GUEST
	default:
		return pb.PlayerRole_PLAYER_ROLE_UNKNOWN
	}
}

func garbageToPB(garbage Garbage) pb.Garbage {
	switch garbage {
	case GARBAGE_GLASS:
		return pb.Garbage_GARBAGE_GLASS
	case GARBAGE_METAL:
		return pb.Garbage_GARBAGE_METAL
	case GARBAGE_ORGANIC:
		return pb.Garbage_GARBAGE_ORGANIC
	case GARBAGE_PAPER:
		return pb.Garbage_GARBAGE_PAPER
	case GARBAGE_PLASTIC:
		return pb.Garbage_GARBAGE_PLASTIC
	default:
		return pb.Garbage_GARBAGE_UNKNOWN
	}
}
