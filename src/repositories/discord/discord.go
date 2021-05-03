package discord

import (
	"errors"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Discord struct{}

var session *discordgo.Session

func (this Discord) Auth(token string) (interface{}, error) {
	_session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, errors.New("Init `discordgo` module")
	}

	session = _session
	return nil, nil
}

type ReactionHandler = func(s *discordgo.Session, m *discordgo.MessageCreate) error
type Reaction struct {
	ChannelName string
	TriggerWord string
	Handler     ReactionHandler
}

func (this Discord) AddSimpleReactions(r []Reaction) error {
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		channel, err := session.Channel(m.ChannelID)
		if err != nil {
			return
		}

		for _, v := range r {
			if strings.Index(m.Content, v.TriggerWord) != -1 && (strings.EqualFold(v.ChannelName, channel.Name) || strings.EqualFold(v.ChannelName, "*")) {
				err := v.Handler(s, m)
				if err != nil {
					log.Println("Error sending message: ", err)
				}
			}
		}
	})

	return nil
}

func (this Discord) Listen() error {
	return session.Open()
}

func (this Discord) Close() {
	session.Close()
}
