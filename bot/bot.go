// bot is a very simple wrapper around the discordgo package
package bot

import "fmt"
import "github.com/bwmarrin/discordgo"

type bot struct {

	// The account that the Bot is authenticated as
	ID string

	// The account of a "admin" or "owner" of this Bot
	Owner *discordgo.User

	// include discordgo Session
	*discordgo.Session
}

func New() *bot {

	bot := bot{}
	bot.Session, _ = discordgo.New()

	return &bot
}

// Opens the connection to Disgord
func (b *bot) Open() error {

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

	err = b.Session.Open()

	return err
}
