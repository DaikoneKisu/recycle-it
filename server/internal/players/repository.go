package players

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/DaikoneKisu/recycle-it/server/internal/db"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) CreatePlayer(ctx context.Context, creationData PlayerCreation) (Player, error) {
	return db.EnsureTx(ctx, r.db, func(ctx context.Context, tx *gorm.DB) (Player, error) {
		dbGame := db.Game{
			ID: creationData.GameID,
		}
		rowsAffected := tx.Preload("Players").First(&dbGame).RowsAffected
		if rowsAffected == 0 {
			return Player{}, fmt.Errorf("there is no game with ID `%v`", creationData.GameID)
		}

		_, playerAlreadyExists := r.GetPlayer(ctx, creationData.GameID, creationData.Nickname)
		if playerAlreadyExists {
			return Player{}, fmt.Errorf("already exists a player with the nickname `%v` in the game with ID `%v`",
				creationData.Nickname, creationData.GameID,
			)
		}
		if isValid, errMsg := isNicknameFormatValid(creationData.Nickname); !isValid {
			return Player{}, fmt.Errorf("invalid nickname: %v", errMsg)
		}
		if isValid, errMsg := isRoleFormatValid(creationData.Role); !isValid {
			return Player{}, fmt.Errorf("invalid role: %v", errMsg)
		}
		if isAvailable, errMsg := isRoleAvailable(creationData.Role, dbGame); !isAvailable {
			return Player{}, fmt.Errorf("role not available: %v", errMsg)
		}

		dbPlayer := db.Player{
			GameID:           creationData.GameID,
			Nickname:         creationData.Nickname,
			Role:             creationData.Role,
			GarbageToCollect: r.pickRandomAvailableGarbage(ctx, creationData.GameID),
		}
		err := tx.Create(&dbPlayer).Error
		if err != nil {
			return Player{}, err
		}
		return Player{
			GameID:   dbPlayer.GameID,
			Nickname: dbPlayer.Nickname,
			Role:     dbPlayer.Role,
		}, nil
	})
}

type PlayerCreation struct {
	GameID   string
	Nickname string
	Role     PlayerRole
}

func isNicknameFormatValid(nickname string) (bool, string) {
	const WHITESPACE_CHARS = " \t\r\n"
	if strings.ContainsAny(nickname, WHITESPACE_CHARS) {
		return false, "the nickname cannot contain any form of whitespace"
	}

	if len(nickname) == 0 {
		return false, "the nickname cannot be empty"
	}

	return true, ""
}

func isRoleFormatValid(role PlayerRole) (bool, string) {
	if role == PLAYER_ROLE_HOST || role == PLAYER_ROLE_GUEST {
		return true, ""
	}
	return false, fmt.Sprintf("the player role must be one of the following: %v, %v",
		PLAYER_ROLE_HOST, PLAYER_ROLE_GUEST,
	)
}

func isRoleAvailable(role db.PlayerRole, game db.Game) (bool, string) {
	if role == db.PLAYER_ROLE_GUEST {
		return true, ""
	}
	if len(game.Players) != 0 {
		return false, fmt.Sprintf("there already is a host on the game with ID %v", game.ID)
	}
	return true, ""
}

func (r Repository) pickRandomAvailableGarbage(ctx context.Context, gameID string) db.Garbage {
	garbage, _ := db.EnsureTx(ctx, r.db, func(ctx context.Context, tx *gorm.DB) (db.Garbage, error) {
		allGarbages := []db.Garbage{
			db.GARBAGE_PLASTIC,
			db.GARBAGE_GLASS,
			db.GARBAGE_PAPER,
			db.GARBAGE_ORGANIC,
			db.GARBAGE_METAL,
		}
		var alreadyPickedGarbages []db.Garbage
		tx.Where("game_id = ?", gameID).Select("garbage_to_collect").Find(&alreadyPickedGarbages)

		var availableGarbages []db.Garbage
		for _, g := range allGarbages {
			if isGarbageAvailable(g, alreadyPickedGarbages) {
				availableGarbages = append(availableGarbages, g)
			}
		}
		return availableGarbages[rand.IntN(len(availableGarbages))], nil
	})
	return garbage
}

func isGarbageAvailable(garbage db.Garbage, alreadyPickedGarbages []db.Garbage) bool {
	for _, g := range alreadyPickedGarbages {
		if g == garbage {
			return false
		}
	}
	return true
}

func (r Repository) GetPlayer(ctx context.Context, gameID string, nickname string) (Player, bool) {
	player, err := db.EnsureTx(ctx, r.db, func(ctx context.Context, tx *gorm.DB) (Player, error) {
		var dbPlayer db.Player
		rowsAffected := tx.First(&dbPlayer, "game_id = ? AND nickname = ?", gameID, nickname).RowsAffected

		if rowsAffected == 0 {
			return Player{}, errors.New("")
		}
		return Player{
			GameID:   dbPlayer.GameID,
			Nickname: dbPlayer.Nickname,
			Role:     dbPlayer.Role,
		}, nil
	})
	if err != nil {
		return Player{}, false
	}
	return player, true
}

type Player struct {
	GameID   string
	Nickname string
	Role     PlayerRole
}

type PlayerRole = string

const (
	PLAYER_ROLE_HOST  PlayerRole = "host"
	PLAYER_ROLE_GUEST PlayerRole = "guest"
)
