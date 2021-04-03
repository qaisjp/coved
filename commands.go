package main

import (
	"github.com/bwmarrin/discordgo"
)

func putCommands(s *discordgo.Session) {
	_, err := ApplicationCommandsCreate(s, theAppID, theGuild, []*discordgo.ApplicationCommand{
		{
			ID:          "827733837282803772",
			Name:        "horny",
			Description: "wawaweewa",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Description: "who to move to the horny channel",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    true,
				},
			},
		},
	})

	if err != nil {
		panic(err.Error())
	}
}
