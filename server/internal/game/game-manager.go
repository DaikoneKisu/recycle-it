package game

import (
	"context"
	"errors"
	"fmt"

	"github.com/DaikoneKisu/recycle-it/server/internal/db"
	"gorm.io/gorm"
)

type GameManager struct {
	db             *gorm.DB
	gameRepository Repository
}

func NewGameManager(db *gorm.DB, gameRepository Repository) GameManager {
	return GameManager{db: db, gameRepository: gameRepository}
}

func (gm GameManager) HostGame(ctx context.Context, hostNickname string, gameSettings Settings) (chan Lobby, error) {
	game, err := gm.gameRepository.CreateGame(ctx, hostNickname, gameSettings)
	if err != nil {
		return nil, err
	}

	lobbyUpdates := make(chan Lobby)
	go func() {
		lobby, _ := gm.GetLobby(ctx, game.ID)

	updateLoop:
		for !lobby.HasGameStarted {
			select {
			case <-ctx.Done():
				break updateLoop
			case lobbyUpdates <- lobby:
				lobby, _ = gm.GetLobby(ctx, game.ID)
			}
		}
		close(lobbyUpdates)
	}()

	return lobbyUpdates, nil
}

func (gm GameManager) JoinGame(ctx context.Context, gameID string, guestNickname string) (chan Lobby, error) {
	err := gm.gameRepository.AddGuest(ctx, gameID, guestNickname)
	if err != nil {
		return nil, err
	}

	lobbyUpdates := make(chan Lobby)
	go func() {
		lobby, _ := gm.GetLobby(ctx, gameID)

	updateLoop:
		for !lobby.HasGameStarted {
			select {
			case <-ctx.Done():
				break updateLoop
			case lobbyUpdates <- lobby:
				lobby, _ = gm.GetLobby(ctx, gameID)
			}
		}
		close(lobbyUpdates)
	}()

	return lobbyUpdates, nil
}

func (gm GameManager) MovePaddle(ctx context.Context, gameID string, playerNickname string, newLocation Point2D) (Game, error) {
	return db.EnsureTx(ctx, gm.db, func(ctx context.Context, tx *gorm.DB) (Game, error) {
		game, gameExists := gm.gameRepository.GetGameByID(ctx, gameID)
		if !gameExists {
			return Game{}, fmt.Errorf("there is no game with ID `%v`", gameID)
		}
		if !game.HasGameStarted {
			return Game{}, errors.New("cannot move any paddle, the game has not started yet")
		}

		err := tx.Model(
			&db.Player{},
		).Select(
			"paddle_location_x",
			"paddle_location_y",
		).Where(
			"game_id = ? AND nickname = ?", gameID, playerNickname,
		).Updates(
			db.Player{PaddleLocationX: newLocation.X, PaddleLocationY: newLocation.Y},
		).Error
		if err != nil {
			return Game{}, fmt.Errorf("could not update the paddle location: %w", err)
		}

		game, _ = gm.gameRepository.GetGameByID(ctx, gameID)
		return game, nil
	})
}

func (gm GameManager) GetLobby(ctx context.Context, gameID string) (Lobby, error) {
	game, gameExists := gm.gameRepository.GetGameByID(ctx, gameID)
	if !gameExists {
		return Lobby{}, fmt.Errorf("there is no game with ID `%v`", gameID)
	}

	players := make([]PlayerWithoutGameID, len(game.Stage.GarbageCollectors))
	for i, gc := range game.Stage.GarbageCollectors {
		players[i] = gc.Player
	}
	return Lobby{
		GameID:         game.ID,
		Players:        players,
		Settings:       game.Settings,
		HasGameStarted: game.HasGameStarted,
	}, nil
}

type Lobby struct {
	GameID         string
	Players        []PlayerWithoutGameID
	Settings       Settings
	HasGameStarted bool
}
