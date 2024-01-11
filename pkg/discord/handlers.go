package discord

import (
	"discord-bot-2/pkg/rpg"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
)

var responses = map[string]Answers{}

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

func HelloWorldHandler(s *discordgo.Session, m *discordgo.MessageCreate) error {
	log.Println("sending \"world\" message to server")
	_, err := s.ChannelMessageSend(m.ChannelID, "world!")

	if err != nil {
		return err
	}

	return nil
}

func RollDiceHandler(s *discordgo.Session, m *discordgo.MessageCreate, numSides int) error {
	diceResult := rand.Intn(numSides)

	msg := fmt.Sprintf("rolled a %d sided dice and got %d", numSides, diceResult+1)

	_, err := s.ChannelMessageSendReply(m.ChannelID, msg, m.Reference())

	if err != nil {
		return err
	}

	return nil
}

func RpgHandler(s *discordgo.Session, m *discordgo.MessageCreate) error {
	channel, err := s.UserChannelCreate(m.Author.ID)

	if err != nil {
		return err
	}

	if _, ok := rpg.Responses[channel.ID]; !ok {
		rpg.Responses[channel.ID] = rpg.Answers{
			OriginChannelId: m.ChannelID,
			CharacterName:   "",
			CharacterClass:  "",
			AdventureChoice: "",
		}
		s.ChannelMessageSend(channel.ID, "Let's begin your RPG character!")
		s.ChannelMessageSend(channel.ID, "Enter your character's name:")
	}

	return nil
}
