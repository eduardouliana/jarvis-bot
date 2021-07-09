package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const token string = ""

var BotID string

func main() {
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	user, err := discord.User("@me")

	if err != nil {
		fmt.Printf(err.Error())
	}

	BotID = user.ID

	discord.AddHandler(messageHandler)

	err = discord.Open()

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Println("Bot is running")

	<-make(chan struct{})
	return
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.ID == BotID {
		return
	}

	if message.Content == "ping" {
		_, _ = session.ChannelMessageSend(
			message.ChannelID,
			"pong",
		)
	}
}
