package bot

import (
	"fmt"

	"github.com/OctavianUrsu/go-discord-cryptobot/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string

func Start() {
	//creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Making our bot a user using User function
	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Storing our id from u to BotId
	BotId = u.ID

	// Adding handler function to handle our messages using AddHandler from discordgo package. We will declare messageHandler function later
	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content[0:1] == config.BotPrefix {
		r := getPrices(m.Content)
		s.ChannelMessageSend(m.ChannelID, r)
	}
}
