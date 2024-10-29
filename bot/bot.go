package bot

import (
	"discord-cfb-bot/config"
	"discord-cfb-bot/clients"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"strings"
)

var Bot *discordgo.Session

func Start() error {
	var err error
	Bot, err = discordgo.New("Bot " + config.BotToken)
	if err != nil {
		return fmt.Errorf("error creating a Discord session: %w", err)
	}

	fmt.Println("Discord session created successfully")

	Bot.AddHandler(commandHandler)
	err = Bot.Open()
	if err != nil {
		return fmt.Errorf("Error opening connection: %w", err)
	}
	fmt.Println("Bot is now running.")
	return nil
}

func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, "!s ") {
		teamName := strings.TrimSpace(strings.TrimPrefix(m.Content, "!s "))
		response := clients.GetGameInfo(teamName)
		s.ChannelMessageSend(m.ChannelID, response)
	}
}
