package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func horny(s *discordgo.Session, e *discordgo.InteractionCreate) {
	data := e.ApplicationCommandData()
	user := data.Options[0].UserValue(s)
	sourceUserID := e.Member.User.ID

	if sourceUserID != user.ID {
		// check time
		t, ok := savedData.Horny[sourceUserID]
		if ok {
			now := time.Now()
			allowedAfter := t.Add(time.Hour * 8)
			if allowedAfter.After(now) {
				err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "You can use this again after " + allowedAfter.Sub(now).String(),
						Flags:   64, // ephemeral
					},
				})
				fmt.Printf("Failed to respond to %#v, error: %v\n", *e.Interaction, err)
				return
			}
		}
	}

	c := "808809400667209820"
	err := s.GuildMemberMove(theGuild, user.ID, &c)
	if err != nil {
		err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Something went wrong: " + err.Error(),
				Flags:   64, // ephemeral
			},
		})
		fmt.Printf("Failed to respond to %#v, error: %v\n", *e.Interaction, err)
		return
	}

	// update time
	if sourceUserID != user.ID {
		savedData.Horny[e.Member.User.ID] = time.Now()
		saveData()
	}

	// push response
	err = s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Looks like <@%s> is horny ;)", user.ID),
		},
	})
	if err != nil {
		fmt.Printf("Failed to respond to %#v, error: %v\n", *e.Interaction, err)
		return
	}
}
