package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

const (
	prefix string = "!dbot"
	word   string = "hello"
)

// Auth authenticates bot with discord.
// Pass in discord bot token
func Auth(token string) (*discordgo.Session, error) {
	sess, err := discordgo.New(token)
	if err != nil {
		return nil, err
	}

	return sess, nil

}

// CreateHandler prints "world!" if anyone types "hello" into chat
func CreateHandler(s *discordgo.Session) {
	s.AddHandler(
		func(s *discordgo.Session, m *discordgo.MessageCreate) {

			// checks if the author of the message is not the bot
			if m.Author.ID == s.State.User.ID {
				return
			}

			args := strings.Split(m.Content, " ")

			log.Println("args:", args)

			if args[0] != prefix {
				return
			}

			if args[1] == strings.ToUpper(word) || args[1] == word {
				log.Println("sending message to server")
				_, err := s.ChannelMessageSend(m.ChannelID, "world!")

				if err != nil {
					log.Fatal(err)
				}
			}
		})
}
