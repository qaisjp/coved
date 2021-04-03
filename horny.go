package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func horny(s *discordgo.Session, e *discordgo.InteractionCreate) {
	// check time
	t, ok := savedData.Horny[e.Member.User.ID]
	if ok {
		now := time.Now()
		allowedAfter := t.Add(time.Hour * 8)
		if allowedAfter.After(now) {
			err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionApplicationCommandResponseData{
					Content: "You can use this again after " + allowedAfter.Sub(now).String(),
					Flags:   64, // ephemeral
				},
			})
			fmt.Printf("Failed to respond to %#v, error: %v\n", *e.Interaction, err)
			return
		}
	}

	c := "808809400667209820"
	target := e.Data.Options[0].StringValue()
	err := s.GuildMemberMove(theGuild, target, &c)
	if err != nil {
		err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionApplicationCommandResponseData{
				Content: "Something went wrong: " + err.Error(),
				Flags:   64, // ephemeral
			},
		})
		fmt.Printf("Failed to respond to %#v, error: %v\n", *e.Interaction, err)
		return
	}

	// update time
	savedData.Horny[e.Member.User.ID] = time.Now()
	saveData()

	// push response
	err = s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: fmt.Sprintf("Looks like <@%s> is horny ;)", target),
		},
	})
	if err != nil {
		fmt.Printf("Failed to respond to %#v, error: %v\n", *e.Interaction, err)
		return
	}
}
