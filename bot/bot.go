package bot

import (
	"fmt"

	"strings"

	"br.edu.sjc/jarvis/config"
	"github.com/bwmarrin/discordgo"
)

var JarvisID string
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

	JarvisID = user.ID

	jarvis.AddHandler(messageHandler)

	err = jarvis.Open()

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Println("Bot is running")
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if !strings.HasPrefix(message.Content, config.BotPrefix) {
		return
	}

	if message.ID == JarvisID {
		return
	}

	if message.Content == "ping" {
		_, _ = session.ChannelMessageSend(
			message.ChannelID,
			"pong",
		)
	}
}
