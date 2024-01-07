package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
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

type Answers struct {
	OriginChannelId string
	FavFood         string
	FavGame         string
	User            string
}

func (a *Answers) ToMessageEmbed() discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{

		{
			Name:  "User",
			Value: a.User,
		},
		{
			Name:  "Favorite food",
			Value: a.FavFood,
		},
		{
			Name:  "Favorite game",
			Value: a.FavGame,
		},
	}

	return discordgo.MessageEmbed{
		Title:  "New responses!",
		Fields: fields,
	}
}

var responses map[string]Answers = map[string]Answers{}

// UserPromptHandler creates user channel in discord
func UserPromptHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, err := s.UserChannelCreate(m.Author.ID)

	if err != nil {
		log.Fatal(err)
	}

	// if the user is already answers question, ignore it, otherwise ask questions
	if _, ok := responses[channel.ID]; !ok {
		responses[channel.ID] = Answers{
			OriginChannelId: m.ChannelID,
			FavFood:         "",
			FavGame:         "",
		}
		s.ChannelMessageSend(channel.ID, "Hey there! Here are some questions")
		s.ChannelMessageSend(channel.ID, "What's your favorite food?")
	} else {
		s.ChannelMessageSend(channel.ID, "We're still waiting... :)")
	}
}

func HelloWorldHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println("sending \"world\" message to server")
	_, err := s.ChannelMessageSend(m.ChannelID, "world!")

	if err != nil {
		log.Fatal(err)
	}
}

func RollDice(s *discordgo.Session, m *discordgo.MessageCreate, numSides int) error {
	diceResult := rand.Intn(numSides)

	msg := fmt.Sprintf("rolled a %d sided dice and got %d", numSides, diceResult+1)
	_, err := s.ChannelMessageSend(m.ChannelID, msg)

	if err != nil {
		return err
	}

	return nil
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
			if m.GuildID == "" {
				answers, ok := responses[m.ChannelID]
				if !ok {
					return
				}

				if answers.FavFood == "" {
					answers.FavFood = m.Content
					answers.User = m.Author.Username
					s.ChannelMessageSend(m.ChannelID, "Great! What's your favorite game now?")

					responses[m.ChannelID] = answers
					return

				} else {
					answers.FavGame = m.Content
					embed := answers.ToMessageEmbed()
					s.ChannelMessageSendEmbed(answers.OriginChannelId, &embed)

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
				HelloWorldHandler(s, m)
			}

			if args[1] == "prompt" {
				UserPromptHandler(s, m)
			}

			if args[1] == "roll" {
				diceSides, err := strconv.Atoi(args[2])

				if err != nil {
					log.Fatal(err)
				}

				RollDice(s, m, diceSides)
			}
		})
}
