package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const Version = "v0.0.0-alpha"

// Global DiscordGo Session
// In this use case, there is no error that would be returned.
var Session, _ = discordgo.New()

func init() {

	// Get Bot Token
	Session.Token = os.Getenv("DG_TOKEN")
	if Session.Token == "" {
		flag.StringVar(&Session.Token, "t", "", "Discord Authentication Token")
	}
}

func main() {

	fmt.Printf(` 
	________  .__                               .___
	\______ \ |__| ______ ____   ___________  __| _/
	||    |  \|  |/  ___// ___\ /  _ \_  __ \/ __ | 
	||    '   \  |\___ \/ /_/  >  <_> )  | \/ /_/ | 
	||______  /__/____  >___  / \____/|__|  \____ | 
	\_______\/        \/_____/   %-16s\/`+"\n\n", Version)

	// Parse command line arguments
	flag.Parse()

	// Connect to Discord if a Token was provided.
	if Session.Token == "" {
		log.Println("You must provide a Discord authentication token.")
		return
	}

	// Verify the Token and grab user information
	u, err := Session.User("@me")
	if err != nil {
		log.Printf("error fetching user information, %s\n", err)
	}
	Session.State.User = u

	// Open a websocket connection to Discord
	err = Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
	}

	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	Session.Close()

	// Exit Normally.
}
