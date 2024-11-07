package game

import (
	"context"
	"errors"
	"fmt"

	"github.com/DaikoneKisu/recycle-it/server/internal/db"
	"github.com/DaikoneKisu/recycle-it/server/internal/players"
	"gorm.io/gorm"
)

type Repository struct {
	db               *gorm.DB
	playerRepository players.Repository
}

func NewRepository(db *gorm.DB, playerRepository players.Repository) Repository {
	return Repository{db: db, playerRepository: playerRepository}
}

func (r Repository) CreateGame(ctx context.Context, hostNickname string, gameSettings Settings) (Game, error) {
	game, err := db.EnsureTx(ctx, r.db, func(ctx context.Context, tx *gorm.DB) (Game, error) {
		if areValid, errMsg := areSettingsValid(gameSettings); !areValid {
			return Game{}, fmt.Errorf("invalid settings: %v", errMsg)
		}

		dbGame := db.Game{
			RequiredPlayerAmount:   gameSettings.RequiredPlayerAmount,
			GameDurationInSeconds:  gameSettings.GameDurationInSeconds,
			TimeRemainingInSeconds: gameSettings.GameDurationInSeconds,
			HasGameStarted:         false,
		}
		err := tx.Create(&dbGame).Error
		if err != nil {
			return Game{}, err
		}

		_, err = r.playerRepository.CreatePlayer(ctx, players.PlayerCreation{
			GameID:   dbGame.ID,
			Nickname: hostNickname,
			Role:     players.PLAYER_ROLE_HOST,
		})
		if err != nil {
			return Game{}, err
		}

		game, _ := r.GetGameByID(ctx, dbGame.ID)
		return game, nil
	})
	return game, err
}

func areSettingsValid(settings Settings) (bool, string) {
	if settings.RequiredPlayerAmount != 2 && settings.RequiredPlayerAmount != 4 {
		return false, "the required player amount must be 2 or 4"
	}
	if settings.GameDurationInSeconds < 30 {
		return false, "the game must last at least 30 seconds"
	}
	return true, ""
}

func (r Repository) AddGuest(ctx context.Context, gameID string, guestNickname string) error {
	_, err := db.EnsureTx(ctx, r.db, func(ctx context.Context, tx *gorm.DB) (struct{}, error) {
		game, gameExists := r.GetGameByID(ctx, gameID)
		if !gameExists {
			return struct{}{}, fmt.Errorf("game not found: there is no game with ID `%v`", gameID)
		}

		if int(game.Settings.RequiredPlayerAmount) == len(game.Stage.GarbageCollectors) {
			return struct{}{}, fmt.Errorf(
				"cannot add a new player to the game: there already are the %v players needed to play",
				game.Settings.RequiredPlayerAmount,
			)
		}

		_, err := r.playerRepository.CreatePlayer(ctx, players.PlayerCreation{
			GameID:   gameID,
			Nickname: guestNickname,
			Role:     players.PLAYER_ROLE_GUEST,
		})
		if err != nil {
			return struct{}{}, err
		}
		return struct{}{}, nil
	})
	return err
}

func (r Repository) GetGameByID(ctx context.Context, id string) (Game, bool) {
	game, err := db.EnsureTx(ctx, r.db, func(ctx context.Context, tx *gorm.DB) (Game, error) {
		dbGame := db.Game{
			ID: id,
		}
		rowsAffected := tx.Preload("Players").Preload("Players.GarbageCollected").First(&dbGame).RowsAffected
		if rowsAffected == 0 {
			return Game{}, errors.New("")
		}

		garbageCollectors := make([]GarbageCollector, len(dbGame.Players))
		for i, p := range dbGame.Players {
			garbageCollected := make([]Garbage, len(p.GarbageCollected))
			for j, gc := range p.GarbageCollected {
				garbageCollected[j] = Garbage(gc.Garbage)
			}

			garbageCollectors[i] = GarbageCollector{
				Player: PlayerWithoutGameID{
					Nickname: p.Nickname,
					Role:     players.PlayerRole(p.Role),
				},
				GarbageToCollect: Garbage(p.GarbageToCollect),
				PaddleLocation: Point2D{
					X: p.PaddleLocationX,
					Y: p.PaddleLocationY,
				},
				GarbageCollected: garbageCollected,
			}
		}

		return Game{
			ID: dbGame.ID,
			Settings: Settings{
				RequiredPlayerAmount:  dbGame.RequiredPlayerAmount,
				GameDurationInSeconds: dbGame.GameDurationInSeconds,
			},
			Stage: Stage{
				GarbageCollectors:  garbageCollectors,
				UncollectedGarbage: Garbage(dbGame.UncollectedGarbage),
				UncollectedGarbageLocation: Point2D{
					X: dbGame.UncollectedGarbageLocationX,
					Y: dbGame.UncollectedGarbageLocationY,
				},
			},
			TimeRemainingInSeconds: dbGame.TimeRemainingInSeconds,
			HasGameStarted:         dbGame.HasGameStarted,
		}, nil
	})
	if err != nil {
		return Game{}, false
	}
	return game, true
}

type Game struct {
	ID                     string
	Settings               Settings
	Stage                  Stage
	TimeRemainingInSeconds int32
	HasGameStarted         bool
}

type Settings struct {
	RequiredPlayerAmount  int32
	GameDurationInSeconds int32
}

type Stage struct {
	GarbageCollectors          []GarbageCollector
	UncollectedGarbage         Garbage
	UncollectedGarbageLocation Point2D
}

type GarbageCollector struct {
	Player           PlayerWithoutGameID
	GarbageToCollect Garbage
	PaddleLocation   Point2D
	GarbageCollected []Garbage
}

type PlayerWithoutGameID struct {
	Nickname string
	Role     players.PlayerRole
}

type Garbage string

const (
	GARBAGE_PLASTIC Garbage = "plastic"
	GARBAGE_GLASS   Garbage = "glass"
	GARBAGE_PAPER   Garbage = "paper"
	GARBAGE_ORGANIC Garbage = "organic"
	GARBAGE_METAL   Garbage = "metal"
)

type Point2D struct {
	X int32
	Y int32
}
