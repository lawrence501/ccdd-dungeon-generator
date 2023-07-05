package main

import (
	"fmt"
	"math/rand"
)

func generateDungeon() Dungeon {
	var dungeon Dungeon
	switch randSelect(DUNGEON_TYPES) {
	case "bastion":
		dungeon = &Bastion{}
		fmt.Println("This dungeon is a bastion, so it gains the Bastion Alert mechanic.\nWhenever the party engages in a combat in the bastion, future monster packs contain 1 additional monster, after CR is calculated.")
	}
	return dungeon
}

func (d *Bastion) generateRoom() string {
	var room Room
	shapeRoll := rand.Intn(20)
	if shapeRoll < 8 {
		room = PassageRoom{}
	} else if shapeRoll < 14 {
		room = SmallRoom{}
	} else {
		room = LargeRoom{}
	}

	dimensions := room.size()
	contents := room.contents(d)
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
	for side2 >= (side1*3) || side1 >= (side2*3) || side1*side2 > 30 {
		side2 = rand.Intn(10) + 1
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
	side1 := rand.Intn(10) + 1
	side2 := 100
	for side2 >= (side1*3) || side1 >= (side2*3) || side1*side2 < 31 {
		side2 = rand.Intn(10) + 1
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

func (r PassageRoom) contents(d Dungeon) string {
	percentile := rand.Intn(100)
	if percentile < 10 {
		return "Empty"
	} else if percentile < 14 {
		return d.passageScenery()
	} else if percentile < 18 {
		return d.escalation()
	}
	return d.obstacle()
}

func (d *Bastion) passageScenery() string {
	if d.PendingEvents[0].Category == "any" || d.PendingEvents[0].Category == "scenery" {
		event, d.PendingEvents := pop(d.PendingEvents)
		return event.Message
	}

	scenery := [][]string{
		[]string{
			"Dusty suit of armour ($bodyArmour)",
			"Mounted monster head (random monster)",
		},
		[]string{
			"Weapons crossed on the wall ($weapon and $weapon)",
		},
		[]string{
			"Racks hold spare weapons and ammunition (1d4x $weapon and 1d10x $ammo)",
			"Barrels hold spare weapons and ammunition (1d4x $weapon and 1d10x $ammo)",
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
		[]string{
			"Pushed into a corner is an archery target bristling with arrows (1d10x $ammo)",
		},
		[]string{
			"A shrine to a war god (Religion DC 13 for +1 global damage for the rest of the dungeon)",
			"A statue of a war god",
		},
	}
	return process(randSelect(randSelect(scenery)))
}

func (d *Bastion) escalation() string {
	if d.PendingEvents[0].Category == "any" || d.PendingEvents[0].Category == "escalation" {
		event, d.PendingEvents := pop(d.PendingEvents)
		return event.Message
	}

	escalationRoll := rand.Intn(10)
	var ret string
	if escalationRoll < 1 {
		ret = "1d6x random monster (MEDIUM), then +x from Bastion Alert level. On patrol, if they meet creatures not dressed as guards, one of them sounds the alarm (+1 Bastion Alert level) as an action and then they attack."
	} else if escalationRoll < 2 {
		ret = "1d6x random monster (MEDIUM), then +x from Bastion Alert level. Off-duty but alert."
	} else if escalationRoll < 5 {
		ret = "1d6x random monster (MEDIUM), then +x from Bastion Alert level. Guarding this room."
	} else if escalationRoll < 6 {
		options := []string{
			"1d6x random monster (MEDIUM), then +x from Bastion Alert level. Dungeon inhabitants planning to ambush the next/current HARD combat due to disloyalty",
			"1d6x random monster (MEDIUM), then +x from Bastion Alert level. Dungeon inhabitants planning to ambush the next/current HARD combat due to ambition",
			"1d6x random monster (MEDIUM), then +x from Bastion Alert level. Dungeon inhabitants planning to ambush the next/current HARD combat due to revenge for a past betrayal",
			"1d6x random monster (MEDIUM), then +x from Bastion Alert level. Dungeon inhabitants planning to ambush the dungeon boss due to disloyalty",
			"1d6x random monster (MEDIUM), then +x from Bastion Alert level. Dungeon inhabitants planning to ambush the dungeon boss due to ambition",
			"1d6x random monster (MEDIUM), then +x from Bastion Alert level. Dungeon inhabitants planning to ambush the dungeon boss due to revenge for a past betrayal",
		}
		optionRoll := rand.Intn(6)
		if optionRoll < 3 {
			d.PendingEvents = append(d.PendingEvents, Event{
				Category: "any",
				Message:  d.setPiece(),
			})
		}
		ret = options[optionRoll]
	} else {
		options := []string {
			"You find monster tracks leading into the next room. Survival DC 13 to generate the next room before opening the door.",
			"You hear monster noises from the next room. Perception DC 13 to generate the next room before opening the door.",
			"You see the flickering torchlight of monsters under the doorway to the next room. Perception DC 13 to generate the next room before opening the door."
		}
		d.PendingEvents = append(d.PendingEvents, Event{
			Category: "any",
			Message: d.escalation(),
		})
		ret = randSelect(options)
	}
	lootRoll = rand.Intn(2)
	if lootRoll == 0 {
		ret += "\n[LOOT] 1x Treasure"
	}
	if d.PendingEvents[0].Category == "combat reward" {
		event, d.PendingEvents := pop(d.PendingEvents)
		ret += fmt.Sprintf("\n%s", event)
	}
	return ret
}

func (d *Bastion) obstacle() string {
	if d.PendingEvents[0].Category == "any" || d.PendingEvents[0].Category == "obstacle" {
		event, d.PendingEvents := pop(d.PendingEvents)
		return event.Message
	}

	escalationRoll := rand.Intn(20)
	if escalationRoll < 1 {
		ret = "A locked chest, which demands today's password. Inside is 1x Treasure. Lock cannot be picked."
		d.PendingEvents = append(d.PendingEvents, Event{
			Category: "combat reward",
			Message:  "[QUEST ITEM] A list of passwords, each next to a day of the week.",
		})
	} else if escalationRoll < 2 {
		ret = "A locked chest, bearing a family crest. Inside is 1x Treasure. Lock cannot be picked."
		d.PendingEvents = append(d.PendingEvents, Event{
			Category: "combat reward",
			Message:  "[QUEST ITEM] A key bearing a family crest.",
		})
	} else if escalationRoll < 3 {
		ret = "A locked chest, with an indentation in the shape of a gauntlet. Inside is 1x Treasure. Lock cannot be picked."
		d.PendingEvents = append(d.PendingEvents, Event{
			Category: "combat reward",
			Message:  "[QUEST ITEM] A gold-plated gauntlet.",
		})
	} else if escalationRoll < 4 {
		ret = "A locked chest, with 7 locks. Inside is 1x Treasure. Lock cannot be picked."
		d.PendingEvents = append(d.PendingEvents, Event{
			Category: "combat reward",
			Message:  "[QUEST ITEM] A keyring with 7 keys.",
		})
	} else if escalationRoll < 5 {
		ret = "A locked wardrobe. Inside is 1x Treasure. Has no key."
	} else if escalationRoll < 6 {
		ret = "An arcane locked chest bearing a bronze face with a unique facial expression. The door is unlocked to any creature perfectly imitating the expression (Performance DC 13). Inside is 1x Treasure."
	} else if escalationRoll < 7 {
		ret = "A brick chest, requiring a DC 13 Strength check to smash open. Inside is 1x Treasure. There is also a magically-lit fireplace (can only be extinguished with water and magic) that contains a key to the brick chest. Pulling it out of the lit fire deals Trap fire damage."
	}
}
