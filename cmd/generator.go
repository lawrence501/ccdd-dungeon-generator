package main

import (
	"fmt"
	"math/rand"
)

func generateDungeon() Dungeon {
	dungeon := Dungeon{}
	switch randSelect(DUNGEON_TYPES) {
	case "bastion":
		dungeon.Type = "bastion"
		dungeon.Functions = &BastionFunctions{}
		fmt.Println("This dungeon is a bastion, so it gains the Bastion Alert mechanic.\nAfter the first 3 times the party engages in a combat in the bastion, future monster packs contain 1 additional monster, after CR is calculated.")
	}
	return dungeon
}

func generateRoom(d Dungeon) string {
	room := d.Functions.room()
	dimensions := generateSize(room)
	contents := room.contents(d)
	exits := d.Functions.exits()
	return fmt.Sprintf("ROOM:\nSize: %s\nExits: %s\nContents: %s", dimensions, exits, contents)
}

func (f *BastionFunctions) room() Room {
	shapeRoll := rand.Intn(20)
	if shapeRoll < 8 {
		return PassageRoom{}
	} else if shapeRoll < 14 {
		return SmallRoom{}
	}
	return LargeRoom{}
}

func generateSize(r Room) string {
	x, y := r.size()
	axisRoll := rand.Intn(2)
	if axisRoll == 0 {
		x, y = y, x
	}
	return fmt.Sprintf("%d x %d", x, y)
}

func (r PassageRoom) size() (int, int) {
	shortSide := rand.Intn(3) + 1
	longSide := 0
	for longSide < (shortSide * 3) {
		longSide = rand.Intn(8) + 3
	}
	return shortSide, longSide
}

func (r SmallRoom) size() (int, int) {
	side1 := rand.Intn(10) + 1
	side2 := 100
	for side2 >= (side1*3) || side1 >= (side2*3) || side1*side2 > 30 {
		side2 = rand.Intn(10) + 1
	}
	return side1, side2
}

func (r LargeRoom) size() (int, int) {
	side1 := rand.Intn(10) + 1
	side2 := 100
	for side2 >= (side1*3) || side1 >= (side2*3) || side1*side2 < 31 {
		side2 = rand.Intn(10) + 1
	}
	return side1, side2
}

func (f *BastionFunctions) exits() string {
	exitRoll := rand.Intn(20)
	if exitRoll < 3 {
		return "None"
	} else if exitRoll < 5 {
		return "Left"
	} else if exitRoll < 7 {
		return "Forward"
	} else if exitRoll < 9 {
		return "Right"
	} else if exitRoll < 11 {
		return "Left and right"
	} else if exitRoll < 13 {
		return "Left and forward"
	} else if exitRoll < 15 {
		return "Forward and right"
	} else if exitRoll < 18 {
		return "Forward, left, and right"
	}
	stairsRoll := rand.Intn(8)
	if stairsRoll < 2 {
		return "Stone stairs down"
	} else if stairsRoll < 3 {
		return "Stone spiral staircase down"
	} else if stairsRoll < 4 {
		options := []string{
			"Trapdoor down",
			"Trapdoor down hidden beneath rug. Investigation DC 13 to discover",
		}
		return randSelect(options)
	} else if stairsRoll < 5 {
		options := []string{
			"Ladder up",
			"Ladder down",
		}
		return randSelect(options)
	} else if stairsRoll < 6 {
		return "Stone spiral staircase up"
	} else if stairsRoll < 7 {
		return "Trapdoor up"
	}
	return "Stairs going 1d4 levels up and 1d4 levels down"
}

func (r PassageRoom) contents(d Dungeon) string {
	contentRoll := rand.Intn(20)
	if contentRoll < 10 {
		return "Empty"
	} else if contentRoll < 14 {
		return generatePassageScenery(d)
	} else if contentRoll < 18 {
		return generateEscalation(d)
	}
	return generateObstacle(d)
}

func (r SmallRoom) contents(d Dungeon) string {
	contentRoll := rand.Intn(20)
	if contentRoll < 3 {
		return "Empty"
	} else if contentRoll < 8 {
		return generateSmallRoomScenery(d)
	} else if contentRoll < 11 {
		return generateNovelty(d)
	} else if contentRoll < 14 {
		return generateObstacle(d)
	} else if contentRoll < 16 {
		return generateDiscovery(d)
	}
}

func generatePassageScenery(d Dungeon) string {
	if d.PendingEvents[0].Category == "any" || d.PendingEvents[0].Category == "scenery" {
		var event Event
		event, d.PendingEvents = pop(d.PendingEvents)
		return event.Message
	}

	return d.Functions.passageScenery()
}

func (f *BastionFunctions) passageScenery() string {
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

func generateSmallRoomScenery(d Dungeon) string {
	if d.PendingEvents[0].Category == "any" || d.PendingEvents[0].Category == "scenery" {
		var event Event
		event, d.PendingEvents = pop(d.PendingEvents)
		return event.Message
	}

	return d.Functions.smallRoomScenery()
}

func (f *BastionFunctions) smallRoomScenery() string {
	sceneryRoll := rand.Intn(12)
	if sceneryRoll < 1 {
		return "Tables bearing a total of 1d4x candles and an equal number of decks of cards are scattered around the room."
	} else if sceneryRoll < 2 {
		options := [][]string{
			[]string{
				"Armoury containing a spear, a shortsword, a crossbow, and a suit of studded leather armour.",
			},
			[]string{
				"Armoury containing a spear, a shortsword, a crossbow, a suit of studded leather armour, and 1x vial of holy water.",
				"Armoury containing a spear, a shortsword, a crossbow, a suit of studded leather armour, and 1x vial of acid.",
				"Armoury containing a spear, a shortsword, a crossbow, a suit of studded leather armour, and 1x vial of alchemist's fire.",
				"Armoury containing a spear, a shortsword, a crossbow, a suit of studded leather armour, and 1x box of caltrops.",
			},
		}
		return randSelect(randSelect(options))
	} else if sceneryRoll < 3 {
		return "Elegant dining room for 6 guests. Table settings are worth 200gp in total. Two bottles of vintage wine, worth 50gp each, stand on a side table."
	} else if sceneryRoll < 4 {
		return "Barracks with neatly-made beds, weapon racks (1d4x $weapon), and a table heaped with armour (1d4x $bodyArmour), game boards, and personal possessions."
	} else if sceneryRoll < 5 {
		return "Comfortable barracks used by high-level guards. On the walls are weapon racks (1d4x $weapon), armour stands (1d4x $bodyArmour), paintings, and a full-length mirror."
	} else if sceneryRoll < 6 {
		return "A shrine on which are laid fresh food offerings (Religion DC 13 for +1 dmg barrier for the rest of the dungeon)"
	} else if sceneryRoll < 7 {
		return "A guardroom containing benches and tables, cards and game boards, wine jugs, and 1d4x plates of half-eaten food."
	} else if sceneryRoll < 8 {
		return "A small kitchen containing a fireplace, wine ready to mull, 1d4x blocks of cheese, and 1d4x barrels of biscuits."
	} else if sceneryRoll < 9 {
		return "A pantry stocked with flour and 1d4x cans of beans, and hung with herbs."
	} else if sceneryRoll < 10 {
		return "A library containing works of strategy and history, as well as atlases and sheafs of papers."
	} else if sceneryRoll < 11 {
		options := []string{
			"A latrine.",
			"A latrine. The smelly privy leads down to the floor below.",
		}
		return randSelect(options)
	}
	return "Contains spy holes that let you see into adjoining rooms, allowing you to generate the rooms without opening the doors to them. All exits are locked without keys."
}

func generateNovelty(d Dungeon) string {
	if d.PendingEvents[0].Category == "any" || d.PendingEvents[0].Category == "novelty" {
		var event Event
		event, d.PendingEvents = pop(d.PendingEvents)
		return event.Message
	}

	return d.Functions.novelty()
}

func (f *BastionFunctions) novelty() string {
	noveltyRoll := rand.Intn(12)
	if noveltyRoll < 1 {
		return "Cannon with 1d10x cannonballs and barrels of powder. If operated, the cannon deals Trap bludgeoning damage."
	} else if noveltyRoll < 2 {
		return "Magic map of the area around the bastion, allowing the party to find 1x Treasure on their way back to town."
	} else if noveltyRoll < 3 {
		return "Arched bridge leading to an iron door halfway up a wall."
	} else if noveltyRoll < 4 {
		return "Immense, monstrous statues on either side of all doors in the room."
	} else if noveltyRoll < 5 {
		return "Drawbridge made of wall of force"
	} else if noveltyRoll < 6 {
		return "Miniature model of the bastion, populated by tiny illusions of its inhabitants. For the remainder of the dungeon, you can always generate rooms before opening their doors."
	} else if noveltyRoll < 7 {
		return "Beasts heads mounted on the wall. Although bodiless, they are alive and can bite for Trap piercing damage."
	} else if noveltyRoll < 8 {
		return "A marble table around which sit the spirits of dead warriors re-enacting an ancient feast."
	} else if noveltyRoll < 9 {
		return "Hundreds of life-sized, sculpted warriors standing in battle array."
	} else if noveltyRoll < 10 {
		options := []string{
			"A war banner 20 feet on each side.",
			"A war banner 30 feet on each side.",
		}
		return randSelect(options)
	} else if noveltyRoll < 11 {
		return "Portrait gallery. Each portrait is enchanted with a permanent Magic Mouth spell."
	}
	return "Narrow shafts that carry sound. It is perfect for eavesdropping or communicating between distant chambers. You can generate adjoining rooms without opening their doors."
}

func generateDiscovery(d Dungeon) string {
	if d.PendingEvents[0].Category == "any" || d.PendingEvents[0].Category == "discovery" {
		var event Event
		event, d.PendingEvents = pop(d.PendingEvents)
		return event.Message
	}

	ret, newEvents := d.Functions.discovery()
	for _, e := range newEvents {
		d.PendingEvents = append(d.PendingEvents, e)
	}

	return ret
}

func (f *BastionFunctions) discovery() (string, []Event) {
	discoveryRoll := rand.Intn(12)
	if discoveryRoll < 1 {
		return "[QUEST ITEM] A list of passwords, each next to a day of the week.", []Event{{
			Category: "obstacle",
			Message:  "A locked chest, which demands today's password. Inside is 1x Treasure. Lock cannot be picked.",
		}}
	} else if discoveryRoll < 2 {
		return "[QUEST ITEM] A key bearing a family crest.", []Event{{
			Category: "obstacle",
			Message:  "A locked chest, bearing a family crest. Inside is 1x Treasure. Lock cannot be picked.",
		}}
	} else if discoveryRoll < 3 {
		return "[QUEST ITEM] A gold-plated gauntlet.", []Event{{
			Category: "obstacle",
			Message:  "A locked chest, with an indentation in the shape of a gauntlet. Inside is 1x Treasure. Lock cannot be picked.",
		}}
	} else if discoveryRoll < 4 {
		return "[QUEST ITEM] A keyring with 7 keys.", []Event{{
			Category: "obstacle",
			Message:  "A locked chest, with 7 locks. Inside is 1x Treasure. Lock cannot be picked.",
		}}
	} else if discoveryRoll < 5 {
		msg, newEvents := f.escalation()
		msg += "\n[ENCOUNTER MODIFIER] Only one random creature from this group is present, and it is not particularly loyal, so is willing to talk."
		return msg, newEvents
	} else if discoveryRoll < 6 {
		msg, newEvents := f.escalation()
		options := []string{
			"\n[ENCOUNTER MODIFIER] This group of creatures are dissatisfied with their commander and willing to turn a blind eye to intruders.",
			"\n[ENCOUNTER MODIFIER] This group of creatures are dissatisfied with their commander and willing to aid intruders.",
		}
		msg += randSelect(options)
		return msg, newEvents
	} else if discoveryRoll < 7 {
		msg, newEvents := f.escalation()
		msg += "\n[ENCOUNTER MODIFIER] This group of creatures are exchanging revealing gossip about their commander and not paying attention to surroundings. One of the dungeon boss' modifiers is generated now."
		return msg, newEvents
	} else if discoveryRoll < 8 {
		return "A messenger with urgent news approaches the party, telling them that if they immediately return to town they can get the quest reward.", []Event{}
	} else if discoveryRoll < 9 {
		options := [][]string{
			[]string{
				"An armoury containing ranged weapons (1d4x $rangedWeapon), ammunition (1d10x $ammo), and ballistas (deal Trap damage of a type based on their ammunition if operated). It also contains 12x +1 Bow Arrows.",
			},
			[]string{
				"An armoury containing ranged weapons (1d4x $rangedWeapon), ammunition (1d10x $ammo), and ballistas (deal Trap damage of a type based on their ammunition if operated). It also contains 12x +1 Bow Arrows and a Javelin of Lightning.",
				"An armoury containing ranged weapons (1d4x $rangedWeapon), ammunition (1d10x $ammo), and ballistas (deal Trap damage of a type based on their ammunition if operated). It also contains 12x +1 Bow Arrows and 1x Treasure.",
			},
		}
		return process(randSelect(randSelect(options))), []Event{}
	} else if discoveryRoll < 10 {
		return "A richly furnished bed chamber containing a four-poster bed, a desk, wardrobes, and treasure chests containing 1x Treasure for each player.", []Event{}
	} else if discoveryRoll < 11 {
		return "A chest containing officers' armour and uniforms (1d4x suits of $bodyArmour).", []Event{}
	}
	return "Treasure vault filled with several splintered chests and one locked (no key) iron chest containing 1x Treasure for each player.", []Event{}
}

func generateEscalation(d Dungeon) string {
	if d.PendingEvents[0].Category == "any" || d.PendingEvents[0].Category == "escalation" {
		var event Event
		event, d.PendingEvents = pop(d.PendingEvents)
		return event.Message
	}

	ret, newEvents := d.Functions.escalation()
	for _, e := range newEvents {
		d.PendingEvents = append(d.PendingEvents, e)
	}

	lootRoll := rand.Intn(2)
	if lootRoll == 0 {
		ret += "\n[LOOT] 1x Treasure"
	}
	if d.PendingEvents[0].Category == "combat reward" {
		var event Event
		event, d.PendingEvents = pop(d.PendingEvents)
		ret += fmt.Sprintf("\n%s", event.Message)
	}
	return ret
}

func (f *BastionFunctions) escalation() (string, []Event) {
	escalationRoll := rand.Intn(10)
	if escalationRoll < 1 {
		return "1d6x random monster (MEDIUM), then +x from Bastion Alert level. On patrol, if they meet creatures not dressed as guards, one of them sounds the alarm (+1 Bastion Alert level) as an action and then they attack.", []Event{}
	} else if escalationRoll < 2 {
		return "1d6x random monster (MEDIUM), then +x from Bastion Alert level. Off-duty but alert.", []Event{}
	} else if escalationRoll < 5 {
		return "1d6x random monster (MEDIUM), then +x from Bastion Alert level. Guarding this room.", []Event{}
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
		event := Event{}
		if optionRoll < 3 {
			event = Event{
				Category: "any",
				Message:  f.setPiece(),
			}
		}
		return options[optionRoll], []Event{event}
	}
	options := []string{
		"You find monster tracks leading into the next room. Survival DC 13 to generate the next room before opening the door.",
		"You hear monster noises from the next room. Perception DC 13 to generate the next room before opening the door.",
		"You see the flickering torchlight of monsters under the doorway to the next room. Perception DC 13 to generate the next room before opening the door.",
	}
	newEscalationMessage, newEscalationEvents := f.escalation()
	event := Event{
		Category: "any",
		Message:  newEscalationMessage,
	}
	retEvents := []Event{event}
	retEvents = append(retEvents, newEscalationEvents...)
	return randSelect(options), retEvents
}

func generateObstacle(d Dungeon) string {
	if d.PendingEvents[0].Category == "any" || d.PendingEvents[0].Category == "obstacle" {
		var event Event
		event, d.PendingEvents = pop(d.PendingEvents)
		return event.Message
	}

	ret, newEvents := d.Functions.obstacle()
	for _, e := range newEvents {
		d.PendingEvents = append(d.PendingEvents, e)
	}

	return ret
}

func (f *BastionFunctions) obstacle() (string, []Event) {
	obstacleRoll := rand.Intn(20)
	if obstacleRoll < 1 {
		return "A locked chest, which demands today's password. Inside is 1x Treasure. Lock cannot be picked.", []Event{{
			Category: "discovery",
			Message:  "[QUEST ITEM] A list of passwords, each next to a day of the week.",
		}}
	} else if obstacleRoll < 2 {
		return "A locked chest, bearing a family crest. Inside is 1x Treasure. Lock cannot be picked.", []Event{{
			Category: "discovery",
			Message:  "[QUEST ITEM] A key bearing a family crest.",
		}}
	} else if obstacleRoll < 3 {
		return "A locked chest, with an indentation in the shape of a gauntlet. Inside is 1x Treasure. Lock cannot be picked.", []Event{{
			Category: "discovery",
			Message:  "[QUEST ITEM] A gold-plated gauntlet.",
		}}
	} else if obstacleRoll < 4 {
		return "A locked chest, with 7 locks. Inside is 1x Treasure. Lock cannot be picked.", []Event{{
			Category: "discovery",
			Message:  "[QUEST ITEM] A keyring with 7 keys.",
		}}
	} else if obstacleRoll < 5 {
		return "A locked wardrobe. Inside is 1x Treasure. Has no key.", []Event{}
	} else if obstacleRoll < 6 {
		return "An arcane locked chest bearing a bronze face with a unique facial expression. The door is unlocked to any creature perfectly imitating the expression (Performance DC 13). Inside is 1x Treasure.", []Event{}
	} else if obstacleRoll < 7 {
		return "A brick chest, requiring a DC 13 Strength check to smash open. Inside is 1x Treasure. There is also a magically-lit fireplace (can only be extinguished with water and magic) that contains a key to the brick chest. Pulling it out of the lit fire deals Trap fire damage.", []Event{}
	} else if obstacleRoll < 8 {
		return "A chest hidden behind a fingerprint-stained mirror. DC 13 Investigation to open. Inside is 1x Treasure.", []Event{}
	} else if obstacleRoll < 9 {
		return "A chest on balcony high up the wall. There is no ladder up to the balcony. Inside is 1x Treasure.", []Event{}
	} else if obstacleRoll < 10 {
		return "A chest hidden behind a mounted, bronze deer head, which is dull except for one shiny antler. Turning the shiny antler reveals the chest. DC 13 Perception to notice the antler. Inside chest is 1x Treasure.", []Event{}
	} else if obstacleRoll < 11 {
		return "Throne room, there is a button on each armrest of the throne. 50% for each option, one opens a hidden pit trap (presser makes DEX save DC 13 or falls 10 x level feet), other reveals a hidden chest that contains 1x Treasure.", []Event{}
	} else if obstacleRoll < 12 {
		return "A harmless ghost of the bestion's former seneschal who knows where a secret chest is in this room (contains 1x Treasure). He, by default, does not want intruders to have the treasure. It can't be found without his help.", []Event{}
	}
	return "A crossbow trap goes off, firing at the first creature to enter the room. Attack of prof + 5 vs AC, dealing Trap piercing damage.", []Event{}
}
