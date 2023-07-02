package main

var DUNGEON_TYPES []string = []string{
	"bastion",
	"laboratory",
	"mine",
	"ruin",
	"sewer",
	"temple",
	"tomb",
}

type RoomShape interface {
	size() string
	contents() string
}

type PassageRoom struct{}
type SmallRoom struct{}
type LargeRoom struct{}

type Dungeon interface {
	generateRoom() string
}

type DungeonData struct {
	Type          string
	PendingEvents []string
}

type Bastion struct {
	DungeonData
	CombatCount int
}
