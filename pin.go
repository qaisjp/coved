package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func pin(s *discordgo.Session, e *discordgo.InteractionCreate, unpin bool) {
	msgID := e.ApplicationCommandData().TargetID
	channelID := e.ApplicationCommandData().Resolved.Messages[msgID].ChannelID

	var err error
	if unpin {
		err = s.ChannelMessageUnpin(channelID, msgID)
	} else {
		err = s.ChannelMessagePin(channelID, msgID)
	}

	msg := "Pinned"
	if unpin {
		msg = "Unpinned"
	}
	msg += " " + fmt.Sprintf("[message](https://discord.com/channels/%s/%s/%s)", e.GuildID, channelID, msgID)

	if err != nil {
		msg = "Failed: " + err.Error()
	}

	err = s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
	if err != nil {
		fmt.Printf("pin - Failed to respond to %#v, error: %v\n", *e.Interaction, err)
		return
	}
}
