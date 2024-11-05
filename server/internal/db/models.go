package db

import "github.com/google/uuid"

type Player struct {
	Nickname        string `gorm:"primaryKey"`
	HashedPassword  string `gorm:"notNull"`
	LobbyID         *uuid.UUID
	LobbyMembership *string
}

type Lobby struct {
	ID                    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Members               []Player  `gorm:"foreignKey:LobbyID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PlayerCapacity        int       `gorm:"notNull"`
	GameDurationInSeconds int       `gorm:"notNull"`
}
