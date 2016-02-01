# Disgord

Disgord is primarily used as a test tool for the [Discordgo](https://github.com/bwmarrin/discordgo) package.  Discordgo provides Go bindings for the Discord Chat API.

This is currently a placeholder.

Join [#go_discordgo](https://discord.gg/0SBTUU1wZTWT6sqd) Discord API channel on Discord for support.



# Goals

The main goal is to develop a scriptable and easily extendable Discord client 
that can be used as a bot framework, a console client, a Discord bot, and as the
primary test bot for the [Discordgo](https://github.com/bwmarrin/discordgo) package.

Disgord will handle connections to Discord and provide an easy to use interface
for scripting actions based on Discord events.  Scripts can be loaded from the
filesystem or feed to Disgord on-the-fly while it is running.

Disgord will have a command interface that can be used from the console or
though Discord chat (channels or private messages) to control all aspects of
Disgord.
