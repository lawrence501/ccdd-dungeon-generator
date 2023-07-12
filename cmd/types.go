package main

type Room interface {
	size() (int, int)
	contents(d *Dungeon) string
}

type PassageRoom struct{}
type SmallRoom struct{}
type LargeRoom struct{}

type Dungeon struct {
	Type          string
	PendingEvents []Event
	Functions     DungeonFunctions
}

type Event struct {
	Category string
	Message  string
}

type DungeonFunctions interface {
	room() Room
	passageScenery() string
	smallRoomScenery() string
	escalation() (string, []Event)
	obstacle() (string, []Event)
	exits() string
	novelty() string
	discovery() (string, []Event)
	setPiece() string
}

type BastionFunctions struct{}

var DUNGEON_TYPES []string = []string{
	"bastion",
	// "laboratory",
	// "mine",
	// "ruin",
	// "sewer",
	// "temple",
	// "tomb",
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

var RANGED_WEAPONS []string = []string{
	"Blowgun",
	"Crossbow",
	"Hand Crossbow",
	"Sling",
	"Wand",
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
