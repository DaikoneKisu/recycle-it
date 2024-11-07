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
			"PaddleLocationX",
			"PaddleLocationY",
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

func (gm GameManager) UpdateStage(ctx context.Context, gameID string, updateData GameStageUpdate) (Game, error) {
	return db.EnsureTx(ctx, gm.db, func(ctx context.Context, tx *gorm.DB) (Game, error) {
		game, gameExists := gm.gameRepository.GetGameByID(ctx, gameID)
		if !gameExists {
			return Game{}, fmt.Errorf("there is no game with ID `%v`", gameID)
		}
		if !game.HasGameStarted {
			return Game{}, errors.New("cannot move any paddle, the game has not started yet")
		}

		err := tx.Model(
			&db.Game{ID: gameID},
		).Select(
			"UncollectedGarbage",
			"UncollectedGarbageLocationX",
			"UncollectedGarbageLocationY",
		).Updates(&db.Game{
			UncollectedGarbage:          string(updateData.UncollectedGarbage),
			UncollectedGarbageLocationX: updateData.UncollectedGarbageLocation.X,
			UncollectedGarbageLocationY: updateData.UncollectedGarbageLocation.Y,
		}).Error
		if err != nil {
			return Game{}, fmt.Errorf("could not update the game stage: %w", err)
		}

		for _, gc := range updateData.GarbageCollectors {
			err := tx.Model(
				&db.Player{Nickname: gc.PlayerNickname},
			).Select(
				"PaddleLocationX",
				"PaddleLocationY",
			).Updates(&db.Player{
				PaddleLocationX: gc.PaddleLocation.X,
				PaddleLocationY: gc.PaddleLocation.Y,
			}).Error
			if err != nil {
				return Game{}, fmt.Errorf("could not update the player data of `%v`: %w", gc.PlayerNickname, err)
			}

			garbageCollected := make([]db.GarbageCollected, len(gc.GarbageCollected))
			for i, g := range gc.GarbageCollected {
				garbageCollected[i] = db.GarbageCollected{Garbage: g}
			}

			err = tx.Session(&gorm.Session{FullSaveAssociations: true}).Model(
				&db.Player{GameID: gameID, Nickname: gc.PlayerNickname},
			).Association(
				"GarbageCollected",
			).Replace(
				garbageCollected,
			)
			if err != nil {
				return Game{}, fmt.Errorf("could not update the garbage collected by `%v`: %w", gc.PlayerNickname, err)
			}
		}

		game, _ = gm.gameRepository.GetGameByID(ctx, gameID)
		return game, nil
	})
}

type GameStageUpdate struct {
	GarbageCollectors          []GarbageCollectorUpdate
	UncollectedGarbage         Garbage
	UncollectedGarbageLocation Point2D
}

type GarbageCollectorUpdate struct {
	PlayerNickname   string
	PaddleLocation   Point2D
	GarbageCollected []Garbage
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
