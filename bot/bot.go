package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	command "github.com/derfoh/discord-dog-bot/commands"
	"github.com/derfoh/discord-dog-bot/config"
)

// BotID is the id set by discord in the main function
var BotID string
var goBot *discordgo.Session

// Start starts opens connections and starts the bot
func Start() {
	// connect
	goBot, err := discordgo.New("Bot " + config.Token)
	checkExit(err)

	// get bot info and set bot id with user info
	u, err := goBot.User("@me")
	checkLog(err)

	BotID = u.ID

	// Add handlers
	goBot.AddHandler(messageHandler)

	// Open connection
	err = goBot.Open()
	checkExit(err)

	// Log bot is running
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.\nPress CTRL-C to exit.")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// check for prefix
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		// ignore the bots messages
		if m.Author.ID == BotID {
			return
		}

		if strings.Contains(m.Content, "dog") {
			response := command.Dog()
			s.ChannelMessageSend(m.ChannelID, response)
		}

		if strings.Contains(m.Content, "ping") {
			start := time.Now()
			s.ChannelMessageSend(m.ChannelID, "")
			elapsed := time.Since(start)
			s.ChannelMessageSend(m.ChannelID, "Pong!"+elapsed.String())
		}
	}
}

// Stop ends closes the connection
func Stop() {
	// Cleanly close down the Discord session.
	goBot.Close()
}

// basic error checker for go, logs then keeps running
func checkLog(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// basic error checker for go, logs then exits
func checkExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
