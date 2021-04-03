package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func ApplicationCommandsCreate(s *discordgo.Session, appID string, guildID string, cmds []*discordgo.ApplicationCommand) (ccmd []*discordgo.ApplicationCommand, err error) {
	endpoint := discordgo.EndpointApplicationGlobalCommands(appID)
	if guildID != "" {
		endpoint = discordgo.EndpointApplicationGuildCommands(appID, guildID)
	}

	body, err := s.RequestWithBucketID("PUT", endpoint, cmds, endpoint)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &ccmd)
	if err != nil {
		return nil, discordgo.ErrJSONUnmarshal
	}

	return
}

func pause() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
