package interaction

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func DefaultHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Content not found.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		_, _ = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong",
			Flags:   discordgo.MessageFlagsEphemeral,
		})
	}

	// Delete the slash command when it doesn't have any registered handler
	log.Debug().
		Str("app-id", i.AppID).
		Str("guild-id", i.GuildID).
		Str("id", i.ID).
		Str("interaction-id", i.Interaction.ID).
		Msg("deleting this")

	err = s.ApplicationCommandDelete(i.AppID, i.GuildID, i.ID)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete default command")
	}
}
