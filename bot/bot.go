package bot

import (
	"fmt"

	"br.edu.sjc/jarvis/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var jarvis *discordgo.Session

func Start() {
	jarvis, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	user, err := jarvis.User("@me")

	if err != nil {
		fmt.Printf(err.Error())
	}

	BotID = user.ID

	jarvis.AddHandler(messageHandler)

	err = jarvis.Open()

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Println("Bot is running")
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
