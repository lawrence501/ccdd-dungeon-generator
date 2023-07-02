package main

import (
	"fmt"
	"math/rand"
)

func randSelect[S []E, E any](s S) E {
	return s[rand.Intn(len(s))]
}

func generateDungeon() Dungeon {
	var dungeon Dungeon
	switch randSelect(DUNGEON_TYPES) {
	case "bastion":
		dungeon = Bastion{}
	}
	return dungeon
}

func (d Bastion) generateRoom() string {
	var shape RoomShape
	shapeRoll := rand.Intn(20)
	if shapeRoll < 8 {
		shape = PassageRoom{}
	} else if shapeRoll < 14 {
		shape = SmallRoom{}
	} else {
		shape = LargeRoom{}
	}

	size := shape.size()

}

func (r PassageRoom) size() string {
	shortSide := rand.Intn(3) + 1
	longSide := 0
	for longSide < (shortSide * 3) {
		longSide = rand.Intn(8) + 3
	}
	axisRoll := rand.Intn(2)
	x := shortSide
	y := longSide
	if axisRoll == 0 {
		x = longSide
		y = shortSide
	}
	return fmt.Sprintf("%d x %d", x, y)
}

func (r SmallRoom) size() string {
	side1 := rand.Intn(10) + 1
	side2 := 100
	for side2 >= (side1 * 3) {
		side2 = rand.Intn(5) + 1
	}
	axisRoll := rand.Intn(2)
	x := side1
	y := side2
	if axisRoll == 0 {
		x = side2
		y = side1
	}
	return fmt.Sprintf("%d x %d", x, y)
}

func (r LargeRoom) size() string {
	side1 := rand.Intn(5) + 1
	side2 := 100
	for side2 >= (side1 * 3) {
		side2 = rand.Intn(5) + 1
	}
	axisRoll := rand.Intn(2)
	x := side1
	y := side2
	if axisRoll == 0 {
		x = side2
		y = side1
	}
	return fmt.Sprintf("%d x %d", x, y)
}

func (r PassageRoom) contents() string {
	percentile := rand.Intn(100)
	if percentile < 10 {
		return passageScenery()
	}
}

func passageScenery() string {
	scenery := [][]string{
		[]string{
			"Dusty suit of armour ($bodyArmour)",
			"Mounted monster head (random monster)",
		},
		[]string{
			"Weapons crossed on the wall ($weapon and $weapon)",
		},
		[]string{
			"Racks hold spare weapons and ammunition (1d4x $weapon and 1d20x $ammo)",
			"Barrels hold spare weapons and ammunition (1d4x $weapon and 1d20x $ammo)",
		},
		[]string{
			"Three camp stools gathered around a small table",
		},
		[]string{
			"Tapestries telling the dungeon's story line the walls",
		},
		[]string{
			"Brackets hold torches (1 per side of longest wall every 15')",
		},
	}
	return randSelect(randSelect(scenery))
}
