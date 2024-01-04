package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
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

// HelloWorld prints "world!" if anyone types "hello" into chat
func HelloWorld(s *discordgo.Session) {
	s.AddHandler(
		func(s *discordgo.Session, m *discordgo.MessageCreate) {

			// checks if the author of the message is not the bot
			if m.Author.ID == s.State.User.ID {
				return
			}

			if m.Content == "hello" || m.Content == "HELLO" {
				_, err := s.ChannelMessageSend(m.ChannelID, "world!")

				if err != nil {
					log.Fatal(err)
				}
			}
		})
}
