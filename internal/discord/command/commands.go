package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/erosdesire/discord-nestjs-utility-bot/internal/discord/command/npm"
)

const (
	NpmInspect           = "npm-inspect"
	NpmInspectNameOption = "name"
)

var (
	StaticCommands = []*discordgo.ApplicationCommand{
		npm.Search,
	}
	// RegisteredCommands stores both static and dynamic commands
	// that can be easily configured after the bot is launched
	RegisteredCommands []*discordgo.ApplicationCommand
)