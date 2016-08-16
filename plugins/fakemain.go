// This file exists to allow this folder to compile without error.
// However, this fake main doesn't really do anything else :)
package main

import "github.com/bwmarrin/discordgo"

var Session, _ = discordgo.New()

func main() {
}
