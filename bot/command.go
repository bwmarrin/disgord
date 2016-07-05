package bot

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Holds a message received over Discord
type Message struct {
	*discordgo.Message
	Fields          []string
	Command         string
	GuildID         string
	IsDirected      bool
	IsPrivate       bool
	HasPrefix       bool
	HasMention      bool
	HasMentionFirst bool
}

type Command struct {
	Word        string // Command Word that uniquely identifies this command.
	Description string // Description of this command
	Help        string // Detailed help information for this Command
	Run         func(*Bot, *Message) bool
}

type commands []*Command

func (c commands) find(word string) *Command {

	var fuzzy *Command
	var i int
	var l int = len(word)

	for _, v := range c {

		// If we find an exact match, return that immediately.
		if v.Word == word {
			return v
		}

		// Some "Fuzzy" searching..
		if l > i && strings.HasPrefix(v.Word, word) {
			fuzzy = v
		}
		i++
	}

	return fuzzy
}

// AddCommand allows you to register a bot command
func (b *Bot) AddCommand(word, desc string, cb func(*Bot, *Message) bool) (*Command, error) {

	cmd := Command{}
	cmd.Word = word
	cmd.Description = desc
	cmd.Run = cb
	b.Commands = append(b.Commands, &cmd)

	return &cmd, nil
}

func (b *Bot) onMessageCreate(ds *discordgo.Session, mc *discordgo.MessageCreate) {

	var err error
	var c *discordgo.Channel

	// Ignore all messages created by the Bot account itself
	if mc.Author.ID == b.ID {
		return
	}

	// Fetch the channel for this Message
	c, err = b.State.Channel(mc.ChannelID)
	if err != nil {
		// Try fetching via REST API
		c, err = b.Channel(mc.ChannelID)
		if err != nil {
			log.Printf("unable to fetch Channel for Message")
			return
		}
		// Attempt to add this channel into our State
		err = b.State.ChannelAdd(c)
		if err != nil {
			log.Printf("error updating State with Channel")
		}
	}

	m := &Message{Message: mc.Message, Command: strings.TrimSpace(mc.Content), GuildID: c.GuildID, IsPrivate: c.IsPrivate}

	// Detect Private Message
	if c.IsPrivate {
		m.IsDirected = true
	}

	// Detect @name or @nick mentions
	if !m.IsDirected {

		// Detect if Bot was @mentioned
		for _, v := range m.Mentions {

			if v.ID == b.ID {

				m.IsDirected, m.HasMention = true, true

				reg := regexp.MustCompile(fmt.Sprintf("<@!?(%s)>", b.ID))

				// Was the @mention the first part of the string?
				if reg.FindStringIndex(m.Command)[0] == 0 {
					m.HasMentionFirst = true
				}

				// strip bot mention tags from content string
				m.Command = reg.ReplaceAllString(m.Command, "")

				break
			}
		}
	}

	// Detect prefix mention
	if !m.IsDirected {

		// TODO : Must be changed to support a per-guild user defined prefix
		if strings.HasPrefix(m.Command, "-dg ") {
			m.IsDirected, m.HasPrefix = true, true
			m.Command = strings.TrimPrefix(m.Command, "-dg ")
		}
	}

	// For now, if we're not specifically mentioned we do nothing.
	// later I might add an option for global non-mentioned command words
	if !m.IsDirected {
		return
	}

	// Tokenize the Command string into a slice of words
	m.Fields = strings.Fields(m.Command)

	// no point to continue if there's no fields
	if len(m.Fields) == 0 {
		return
	}

	// Search though the command list for a match
	// first match wins
	var cmd *Command
	for k, v := range m.Fields {

		cmd = b.Commands.find(v)
		if cmd != nil {

			// strip out leading Fields
			m.Fields = m.Fields[k:]

			// run the command.
			cmd.Run(b, m)
			return
		}
	}

	// If no command match was found, call the default.
	// Ignore if only @mentioned in the middle of a message
	if b.DefaultCommand != nil && (m.IsPrivate || m.HasPrefix || m.HasMentionFirst) {

		// TODO: This needs a ratelimit
		b.DefaultCommand.Run(b, m)
	}

}
