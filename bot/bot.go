// bot is a very simple wrapper around the discordgo package with a few extra
// features added for writing Bots with the Disgord Bot "framework".
package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	// include discordgo Session
	*discordgo.Session

	// The account ID that the Bot is authenticated as
	ID string

	// The name of the Bot account
	Name string

	// The nick of the bot account
	Nick string

	// The account of a "admin" or "owner" of this Bot
	Owner *discordgo.User

	// Commands map holds all bot commands
	// not sure if I want to use a map or a slice
	Commands       commands
	DefaultCommand *Command
}

func New() *Bot {

	bot := &Bot{}
	bot.Session, _ = discordgo.New()

	return bot
}

// Opens the connection to Disgord
func (b *Bot) Open() error {

	var err error

	if b.Token == "" {
		return fmt.Errorf("no token")
	}

	// verify the token and grab user ID
	user, err := b.User("@me")
	if err != nil {
		return err
	}
	b.ID = user.ID
	b.Name = user.Username

	b.AddHandler(b.onMessageCreate)

	err = b.Session.Open()

	return err
}
