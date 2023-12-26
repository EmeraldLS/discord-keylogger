package keys

import (
	"log"
	"logger/keylogger/saver"

	"github.com/eiannone/keyboard"
)

func KeyStrokes() {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		log.Println(err)
	}

	log.Println("Press ESC to exit")
	for {
		event := <-keysEvents
		if event.Err != nil {
			log.Fatal(err)
		}
		log.Printf("You pressed: %c ", event.Rune)
		if event.Key == keyboard.KeyEnter {
			event.Rune = '\n'
		}

		if event.Key == keyboard.KeyBackspace {
			event.Rune = '-'
		}

		saver.SaveCharacterPress(event.Rune)
	}

}
