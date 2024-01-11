package discord

import (
	"discord-bot-2/pkg/rpg"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
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

			// DM logic
			//if m.GuildID == "" {
			//	answers, ok := responses[m.ChannelID]
			//	if !ok {
			//		return
			//	}
			//
			//	if answers.FavFood == "" {
			//		answers.FavFood = m.Content
			//		answers.User = m.Author.Username
			//		s.ChannelMessageSend(m.ChannelID, "Great! What's your favorite game now?")
			//
			//		responses[m.ChannelID] = answers
			//		return
			//
			//	} else {
			//		answers.FavGame = m.Content
			//		embed := answers.ToMessageEmbed()
			//		s.ChannelMessageSendEmbed(answers.OriginChannelId, &embed)
			//
			//		delete(responses, m.ChannelID)
			//	}
			//}

			if m.GuildID == "" {
				rpgAnswers, ok := rpg.Responses[m.ChannelID]
				if !ok {
					return
				}

				if rpgAnswers.CharacterName == "" {
					rpgAnswers.CharacterName = m.Content
					s.ChannelMessageSend(m.ChannelID, "Choose a class [healer, mage, ranger, warrior]:")
					rpgAnswers.CharacterClass = m.Content
					s.ChannelMessageSend(m.ChannelID, "Choose an adventure [exploration, dungeon, farming]:")

					rpg.Responses[m.ChannelID] = rpgAnswers
					return

				} else {
					rpgAnswers.CharacterName = m.Content
					embed := rpgAnswers.ToMessageEmbed()
					s.ChannelMessageSendEmbed(rpgAnswers.OriginChannelId, &embed)

					delete(responses, m.ChannelID)
				}
			}

			// server logic
			args := strings.Split(m.Content, " ")

			if args[0] != prefix {
				return
			}

			if len(args) < 2 {
				return
			}

			if args[1] == strings.ToUpper(word) || args[1] == word {
				err := HelloWorldHandler(s, m)

				if err != nil {
					log.Fatal(err)
				}
			}

			if args[1] == "prompt" {
				UserPromptHandler(s, m)
			}

			if args[1] == "roll" {
				diceSides, err := strconv.Atoi(args[2])

				if err != nil {
					log.Fatal(err)
				}

				err = RollDiceHandler(s, m, diceSides)
				if err != nil {
					return
				}
			}

			if args[1] == "rpg" {
				err := RpgHandler(s, m)
				if err != nil {
					log.Fatal(err)
				}
			}

			if args[1] == "mudkip" {
				u := "https://youtu.be/3DkqMjfqqPc?t=79"
				SendVideo(s, m, u)
			}
		})
}
