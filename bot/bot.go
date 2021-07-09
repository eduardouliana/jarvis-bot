package bot

import (
	"fmt"

	"strings"

	cmd "br.edu.sjc/jarvis/bot/commands"
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

func SendMessage(session *discordgo.Session, chanelId string, content string) {
	if content == "" {
		return
	}

	session.ChannelMessageSend(
		chanelId,
		content,
	)
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if !strings.HasPrefix(message.Content, config.BotPrefix) {
		return
	}

	if message.ID == JarvisID {
		return
	}

	messageWithoutPrefix := strings.TrimPrefix(message.Content, config.BotPrefix)

	switch {
	case strings.HasPrefix(messageWithoutPrefix, "ping"):
		SendMessage(session, message.ChannelID, cmd.ExecutePong())

	case strings.HasPrefix(messageWithoutPrefix, "build"):
		SendMessage(session, message.ChannelID, cmd.ExecuteBuildTest())

	default:
		SendMessage(session, message.ChannelID, "Command not found")
	}
}
