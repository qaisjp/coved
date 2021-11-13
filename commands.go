package main

import (
	"github.com/bwmarrin/discordgo"
)

const CommandHornyID = "827733837282803772"
const CommandPinID = "908905946082844722"
const CommandUnpinID = "908905946678452245"

func putCommands(s *discordgo.Session) {
	// s.ApplicationCommandDelete(theAppID, theGuild, CommandPinID)

	// x, _ := s.ApplicationCommandCreate(theAppID, theGuild, &discordgo.ApplicationCommand{
	// 	ID:   CommandPinID,
	// 	Type: discordgo.MessageApplicationCommand,
	// 	Name: "Pin",
	// })
	// y, _ := s.ApplicationCommandCreate(theAppID, theGuild, &discordgo.ApplicationCommand{
	// 	ID:   CommandPinID,
	// 	Type: discordgo.MessageApplicationCommand,
	// 	Name: "Unpin",
	// })
	// fmt.Println("Pin", x.ID, "Unpin", y.ID)

	_, err := ApplicationCommandsCreate(s, theAppID, theGuild, []*discordgo.ApplicationCommand{
		{
			ID:          CommandHornyID,
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
		{
			ID:   CommandPinID,
			Type: discordgo.MessageApplicationCommand,
			Name: "Pin",
		},
		{
			ID:   CommandUnpinID,
			Type: discordgo.MessageApplicationCommand,
			Name: "Unpin",
		},
	})

	if err != nil {
		panic(err.Error())
	}
}
