package main

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var theGuild = "689959312684220633"
var theRole = "827662098426363916"
var theAppID = "827658340591075439"

func main() {
	loadData()

	token, err := os.ReadFile("token.txt")
	if err != nil {
		panic(err.Error())
	}

	discord, err := discordgo.New("Bot " + strings.TrimSpace(string(token)))

	err = discord.Open()
	if err != nil {
		panic(err.Error())
	}

	// cmds, err := discord.ApplicationCommands(theAppID, theGuild)
	// if err != nil {
	// 	fmt.Println("failed to get", cmds)
	// 	return
	// }
	// fmt.Printf("they are: %#v\n", *cmds[0])

	putCommands(discord)
	discord.AddHandler(onInteraction)
	pause()
}

func onInteraction(s *discordgo.Session, e *discordgo.InteractionCreate) {
	if e.Data.Name == "horny" {
		horny(s, e)
		return
	}
}
