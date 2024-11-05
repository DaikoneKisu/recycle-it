package players

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/DaikoneKisu/recycle-it/server/internal/db"
	"github.com/DaikoneKisu/recycle-it/server/internal/passwords"
	"github.com/google/uuid"
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
		_, playerAlreadyExists := r.GetPlayer(ctx, creationData.Nickname)
		if playerAlreadyExists {
			return Player{}, fmt.Errorf("already exists a player with the nickname `%v`", creationData.Nickname)
		}
		if isValid, errMsg := isNicknameFormatValid(creationData.Nickname); !isValid {
			return Player{}, fmt.Errorf("invalid nickname: %v", errMsg)
		}
		if isValid, errMsg := isPasswordFormatValid(creationData.Password); !isValid {
			return Player{}, fmt.Errorf("invalid password: %v", errMsg)
		}

		dbPlayer := db.Player{
			Nickname:        creationData.Nickname,
			HashedPassword:  passwords.HashPassword(creationData.Password),
			LobbyID:         nil,
			LobbyMembership: nil,
		}
		result := tx.Create(&dbPlayer)

		if result.Error != nil {
			return Player{}, result.Error
		}
		return Player(dbPlayer), nil
	})
}

type PlayerCreation struct {
	Nickname string
	Password string
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

func isPasswordFormatValid(password string) (bool, string) {
	const WHITESPACE_CHARS = " \t\r\n"
	if strings.ContainsAny(password, WHITESPACE_CHARS) {
		return false, "the nickname cannot contain any form of whitespace"
	}

	const MIN_PASSWORD_LENGTH = 8
	if len(password) < MIN_PASSWORD_LENGTH {
		return false, fmt.Sprintf("the password must contain at least %v characters", MIN_PASSWORD_LENGTH)
	}

	return true, ""
}

func (r Repository) GetPlayer(ctx context.Context, nickname string) (Player, bool) {
	player, err := db.EnsureTx(ctx, r.db, func(ctx context.Context, tx *gorm.DB) (Player, error) {
		var dbPlayer db.Player
		result := tx.First(&dbPlayer, "nickname = ?", nickname)

		if result.RowsAffected == 0 {
			return Player{}, errors.New("")
		}
		return Player(dbPlayer), nil
	})
	if err != nil {
		return Player{}, false
	}
	return player, true
}

type Player struct {
	Nickname        string
	HashedPassword  string
	LobbyID         *uuid.UUID
	LobbyMembership *string
}

var LobbyMemberships = struct {
	GUEST string
	OWNER string
}{
	GUEST: "guest",
	OWNER: "owner",
}
