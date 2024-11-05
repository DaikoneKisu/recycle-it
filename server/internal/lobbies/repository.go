package lobbies

import (
	"context"
	"errors"
	"fmt"

	"github.com/DaikoneKisu/recycle-it/server/internal/db"
	"github.com/DaikoneKisu/recycle-it/server/internal/players"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db               *gorm.DB
	playerRepository players.Repository
}

func NewRepository(db *gorm.DB, playerRepository players.Repository) Repository {
	return Repository{db: db, playerRepository: playerRepository}
}

func (r Repository) CreateLobby(ctx context.Context, ownerNickname string, settings Settings) (Lobby, error) {
	return db.EnsureTx(ctx, r.db, func(ctx context.Context, tx *gorm.DB) (Lobby, error) {
		player, ownerExists := r.playerRepository.GetPlayer(ctx, ownerNickname)
		if !ownerExists {
			return Lobby{}, fmt.Errorf("owner not found: there is no player with nickname `%v`", ownerNickname)
		}
		if isValid, errMsg := isSettingsFormatValid(settings); !isValid {
			return Lobby{}, fmt.Errorf("invalid settings: %v", errMsg)
		}

		dbLobby := db.Lobby{
			Members:               []db.Player{db.Player(player)},
			PlayerCapacity:        settings.PlayerCapacity,
			GameDurationInSeconds: settings.GameDurationInSeconds,
		}
		tx.Create(&dbLobby)

		lobby, _ := r.GetLobbyByID(ctx, dbLobby.ID)
		return lobby, nil
	})
}

func isSettingsFormatValid(settings Settings) (bool, string) {
	const MAX_PLAYER_CAPACITY = 4
	if settings.PlayerCapacity < 1 || settings.PlayerCapacity > MAX_PLAYER_CAPACITY {
		return false, fmt.Sprintf("the lobby must be able to contain 1 to %v players", MAX_PLAYER_CAPACITY)
	}

	const MIN_GAME_DURATION_IN_SECONDS = 30
	if settings.GameDurationInSeconds < MIN_GAME_DURATION_IN_SECONDS {
		return false, fmt.Sprintf("the game must last %v seconds or more", MIN_GAME_DURATION_IN_SECONDS)
	}

	return true, ""
}

func (r Repository) GetLobbyByID(ctx context.Context, id uuid.UUID) (Lobby, bool) {
	lobby, err := db.EnsureTx(ctx, r.db, func(ctx context.Context, tx *gorm.DB) (Lobby, error) {
		var dbLobby db.Lobby
		result := tx.Joins("Player").First(&dbLobby, id)
		if result.RowsAffected == 0 {
			return Lobby{}, errors.New("")
		}
		return NewFromDBModel(dbLobby), nil
	})

	if err != nil {
		return Lobby{}, false
	}
	return lobby, true
}

func (r Repository) JoinLobby(ctx context.Context, lobbyID uuid.UUID, guestNickname string) (Lobby, error) {
	return db.EnsureTx(ctx, r.db, func(ctx context.Context, tx *gorm.DB) (Lobby, error) {
		var dbLobby db.Lobby
		result := tx.Joins("Player").First(&dbLobby, lobbyID)
		if result.RowsAffected == 0 {
			return Lobby{}, fmt.Errorf("lobby not found: there is no lobby with ID `%v`", lobbyID)
		}

		_, playerExists := r.playerRepository.GetPlayer(ctx, guestNickname)
		if !playerExists {
			return Lobby{}, fmt.Errorf("player not found: there is no player with nickname `%v`", guestNickname)
		}

		tx.Model(&dbLobby).Association("Members").Append(&db.Player{Nickname: guestNickname})
		updatedLobby, _ := r.GetLobbyByID(ctx, lobbyID)
		return updatedLobby, nil
	})
}

type Lobby struct {
	ID       uuid.UUID
	Members  []Member
	Settings Settings
}

func NewFromDBModel(dbLobby db.Lobby) Lobby {
	members := make([]Member, len(dbLobby.Members))
	for i, m := range dbLobby.Members {
		members[i] = Member{Nickname: m.Nickname, Membership: *m.LobbyMembership}
	}

	return Lobby{
		ID:      dbLobby.ID,
		Members: members,
		Settings: Settings{
			PlayerCapacity:        dbLobby.PlayerCapacity,
			GameDurationInSeconds: dbLobby.GameDurationInSeconds,
		},
	}
}

type Member struct {
	Nickname   string
	Membership string
}

type Settings struct {
	PlayerCapacity        int
	GameDurationInSeconds int
}
