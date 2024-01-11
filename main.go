package main

import (
	"discord-bot-2/pkg/discord"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	token := "Bot " + os.Getenv("DISCORD_TOKEN")

	// Creates discord session
	sess, err := discord.Auth(token)

	// processes incoming messages from discord
	discord.CreateHandler(sess)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func(sess *discordgo.Session) {
		err := sess.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(sess)

	log.Println("the bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
