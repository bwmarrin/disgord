<img align="right" src="http://bwmarrin.github.io/disgord/img/gourd.jpg">
Disgord
====

Dis**go**rd is an easy to use and extensible Discord Bot using the [discordgo](https://github.com/bwmarrin/discordgo) library.

**NOTE** This is used by the **dgo** bot which is the official test bot 
for the discordgo package. If this bot is in your server, allowing it to remain
provides indispensable testing of the discordgo package. **Thank you very very much** 
for helping to make Discord and DiscordGo great.

If you would like to help the discordgo and disgord packages please use 
[this link](https://discordapp.com/oauth2/authorize?client_id=173113690092994561&scope=bot)
to add **dgo** to your server. This provides tremendous help to both packages!

**For help with this program or general Go discussion, please join the [Discord 
Gophers](https://discord.gg/0f1SbxBZjYq9jLBk) chat server.**

## Getting Started

The below assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

### Installing

```sh
git clone https://github.com/bwmarrin/disgord.git
cd disgord
```

### Plugins

The plugins folder contains several plugins you can add to your bot, just
copy (or create symbolic links) of these files into the folder with the main.go
file to use them.  Some of them have a dependency on other plugins so please
check the documentation for each plugin.


### Compiling

When ready you can compile Disgord like any Go program.

```sh
go build 
```


### Start Disgord

```go
./disgord -t BOT_TOKEN_HERE
```
