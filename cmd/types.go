package main

type Room interface {
	size() string
	contents(d Dungeon) string
}

type PassageRoom struct{}
type SmallRoom struct{}
type LargeRoom struct{}

type Dungeon interface {
	generateRoom() string
	passageScenery() string
	escalation() string
	obstacle() string
	setPiece() string
}

type DungeonData struct {
	Type          string
	PendingEvents []Event
}

type Event struct {
	Category string
	Message  string
}

type Bastion struct {
	DungeonData
	AlertCounter int
}

var DUNGEON_TYPES []string = []string{
	"bastion",
	"laboratory",
	"mine",
	"ruin",
	"sewer",
	"temple",
	"tomb",
}

var BODY_ARMOURS []string = []string{
	"Padded Armour",
	"Leather Armour",
	"Studded Leather Armour",
	"Hide Armour",
	"Chain Shirt",
	"Scale Mail",
	"Breastplate",
	"Half Plate",
	"Ring Mail",
	"Chain Mail",
	"Splint",
	"Plate",
}

var WEAPONS []string = []string{
	"Club",
	"Dagger",
	"Gauntlet",
	"Longspear",
	"Mace",
	"Morningstar",
	"Sickle",
	"Spear",
	"Staff",
	"Blowgun",
	"Crossbow",
	"Hand Crossbow",
	"Sling",
	"Wand",
	"Bastard Sword",
	"Battle Axe",
	"Boomerang",
	"Bo Staff",
	"Falchion",
	"Flail",
	"Glaive",
	"Greataxe",
	"Greatclub",
	"Greatpick",
	"Greatsword",
	"Guisarme",
	"Halberd",
	"Handaxe",
	"Lance",
	"Light Hammer",
	"Light Pick",
	"Longsword",
	"Main-Gauche",
	"Maul",
	"Nunchaku",
	"Pick",
	"Rapier",
	"Ring Blade",
	"Sap",
	"Scimitar",
	"Scythe",
	"Shortsword",
	"Small Shield",
	"Starknife",
	"Tower Shield",
	"Trident",
	"Warhammer",
	"Whip",
	"Longbow",
	"Net",
	"Shortbow",
}

var AMMO []string = []string{
	"Blowgun Dart",
	"Crossbow Bolt",
	"Sling Stone",
	"Wand Charge",
	"Bow Arrow",
}
