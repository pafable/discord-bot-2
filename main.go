package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(*discord)
	var inputVar string
	scan, err := fmt.Scan(&inputVar)
	if err != nil {
		return
	}

	log.Println(scan, inputVar)
}
