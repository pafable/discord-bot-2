package rpg

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

var BaseHealth = map[string]int{
	"healer":  8,
	"mage":    7,
	"ranger":  9,
	"warrior": 10,
}

var BaseDamage = map[string]int{
	"healer":  2,
	"mage":    6,
	"ranger":  5,
	"warrior": 5,
}

var BaseSpeed = map[string]int{
	"healer":  4,
	"mage":    3,
	"ranger":  6,
	"warrior": 5,
}

var Responses = map[string]Answers{}

type Answers struct {
	OriginChannelId string
	CharacterName   string
	CharacterClass  string
	AdventureChoice string
}

type Player struct {
	PlayerClass Class
	Name        string
}

type Class struct {
	ClassName string
	Health    int
	Dmg       int
	Speed     int
}

func (p *Player) GetDamage(dmg int) {
	log.Printf("%s took %d dmg", p.Name, dmg)
	p.PlayerClass.Health -= dmg
}

func (p *Player) GetHealth(hp int) {
	log.Printf("%s was healed by %d hp", p.Name, hp)
	p.PlayerClass.Health += hp
}

func (p *Player) getHealthCount() int {
	return p.PlayerClass.Health
}

func (p *Player) KillPlayer() string {
	p.PlayerClass.Health = 0
	return fmt.Sprintf("%s is dead!", p.Name)
}

func (p *Player) revivePlayer() {
	switch p.PlayerClass.ClassName {
	case "healer":
		p.PlayerClass.Health = BaseHealth["healer"]
		log.Printf("revived %s with %d", p.Name, BaseHealth["healer"])
	case "mage":
		p.PlayerClass.Health = BaseHealth["mage"]
		log.Printf("revived %s with %d", p.Name, BaseHealth["mage"])
	case "ranger":
		p.PlayerClass.Health = BaseHealth["ranger"]
		log.Printf("revived %s with %d", p.Name, BaseHealth["ranger"])
	case "warrior":
		p.PlayerClass.Health = BaseHealth["warrior"]
		log.Printf("revived %s with %d", p.Name, BaseHealth["warrior"])
	}
}

func (a *Answers) ToMessageEmbed() discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{
		{
			Name:  "Character Name",
			Value: a.CharacterName,
		},
		{
			Name:  "Class",
			Value: a.CharacterClass,
		},
		{
			Name:  "Adventure",
			Value: a.AdventureChoice,
		},
	}

	return discordgo.MessageEmbed{
		Title:  "New character created",
		Fields: fields,
	}
}
