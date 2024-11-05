package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/DaikoneKisu/recycle-it/server/internal/db"
	"github.com/DaikoneKisu/recycle-it/server/internal/passwords"
	"github.com/DaikoneKisu/recycle-it/server/internal/players"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Authenticator struct {
	db *gorm.DB
}

func NewAuthenticator(db *gorm.DB) Authenticator {
	return Authenticator{db: db}
}

func (a Authenticator) SignUp(ctx context.Context, nickname, password string) (string, error) {
	return db.EnsureTx(ctx, a.db, func(ctx context.Context, tx *gorm.DB) (string, error) {
		playerRepository := players.NewRepository(tx)

		player, err := playerRepository.CreatePlayer(ctx, players.PlayerCreation{
			Nickname: nickname,
			Password: password,
		})
		if err != nil {
			return "", fmt.Errorf("an error occurred while creating the player: %w", err)
		}

		jwt, err := buildJWT(player)
		if err != nil {
			return "", fmt.Errorf("an unexpected error occurred while building the jwt: %w", err)
		}
		return jwt, nil
	})
}

func (a Authenticator) SignIn(ctx context.Context, nickname, password string) (string, error) {
	return db.EnsureTx(ctx, a.db, func(ctx context.Context, tx *gorm.DB) (string, error) {
		playerRepository := players.NewRepository(tx)

		player, playerExists := playerRepository.GetPlayer(ctx, nickname)
		if !playerExists {
			return "", errors.New("invalid credentials")
		}
		if !passwords.ArePasswordsEqual(password, player.HashedPassword) {
			return "", errors.New("invalid credentials")
		}

		jwt, err := buildJWT(player)
		if err != nil {
			return "", fmt.Errorf("an unexpected error occurred while building the jwt: %w", err)
		}
		return jwt, nil
	})
}

func buildJWT(player players.Player) (string, error) {
	JWT_SIGNING_KEY := []byte("k+E{puZhk7UfrxyL'G'+hL/EcgblHq")
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": player.Nickname,
		},
	)
	return unsignedToken.SignedString(JWT_SIGNING_KEY)
}
