package connectors

import (
	"github.com/JackMaarek/go-bot-utils/env"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func InitializeBot() (*discordgo.Session, error) {
	discordToken := env.GetVariable("DISCORD_TOKEN")
	if "" == discordToken {
		log.Fatal("Missing environment variable : DISCORD_TOKEN")
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Error("Error creating Discord session, ", err)
		return nil, err
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Error("Error opening discord connection, ", err)
		return nil, err
	}

	return dg, nil
}
