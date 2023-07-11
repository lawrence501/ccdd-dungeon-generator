package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	rand.Seed(time.Now().UnixNano())

	for {
		log.Println()
		baseP := promptui.Prompt{
			Label:    "Press enter to generate a new dungeon",
			Validate: validateAny,
		}
		_, err := baseP.Run()
		if err != nil {
			log.Fatal(err)
			return
		}

		dungeon := generateDungeon()
		for {
			roomP := promptui.Prompt{
				Label:    "Press enter to generate a new dungeon room, or type exit to start a new dungeon",
				Validate: validateAny,
			}
			roomInput, err := roomP.Run()
			if err != nil {
				log.Fatal(err)
				return
			}
			if strings.ToLower(roomInput) == "exit" {
				break
			}
			generateRoom(dungeon)
		}

		if err != nil {
			log.Println("Error occurred during dungeon generation")
			log.Fatal(err)
		}
		log.Println()
	}
}
