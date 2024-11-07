package db

type Game struct {
	ID                          string   `gorm:"default:gen_random_uuid()::text;primaryKey"`
	RequiredPlayerAmount        int32    `gorm:"notNull"`
	GameDurationInSeconds       int32    `gorm:"notNull"`
	Players                     []Player `gorm:"foreignKey:GameID;references:ID"`
	UncollectedGarbage          Garbage  `gorm:"notNull"`
	UncollectedGarbageLocationX int32    `gorm:"notNull"`
	UncollectedGarbageLocationY int32    `gorm:"notNull"`
	TimeRemainingInSeconds      int32    `gorm:"notNull"`
	HasGameStarted              bool     `gorm:"notNull"`
}

type Player struct {
	GameID           string             `gorm:"primaryKey"`
	Nickname         string             `gorm:"primaryKey"`
	Role             PlayerRole         `gorm:"notNull"`
	GarbageToCollect Garbage            `gorm:"notNull"`
	PaddleLocationX  int32              `gorm:"notNull"`
	PaddleLocationY  int32              `gorm:"notNull"`
	GarbageCollected []GarbageCollected `gorm:"foreignKey:GameID,PlayerNickname;references:GameID,Nickname"`
}

type PlayerRole = string

const (
	PLAYER_ROLE_HOST  PlayerRole = "host"
	PLAYER_ROLE_GUEST PlayerRole = "guest"
)

type GarbageCollected struct {
	GameID         string `gorm:"primaryKey"`
	PlayerNickname string `gorm:"primaryKey"`
	GarbageNo      int32  `gorm:"primaryKey;autoIncrement"`
	Garbage        Garbage
}

type Garbage = string

const (
	GARBAGE_UNKNOWN Garbage = ""
	GARBAGE_PLASTIC Garbage = "plastic"
	GARBAGE_GLASS   Garbage = "glass"
	GARBAGE_PAPER   Garbage = "paper"
	GARBAGE_ORGANIC Garbage = "organic"
	GARBAGE_METAL   Garbage = "metal"
)
