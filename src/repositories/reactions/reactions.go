package reactions

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nasustim/d2-bot/src/repositories/discord"
)

func Reactions() []discord.Reaction {
	reactions := []discord.Reaction{
		discord.Reaction{
			ChannelName: "general",
			TriggerWord: "abc",
			Handler: func(s *discordgo.Session, m *discordgo.MessageCreate) error {
				s.ChannelMessageSend(m.ChannelID, "cba")
				return nil
			},
		},
	}
	return reactions
}
