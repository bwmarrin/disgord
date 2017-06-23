# dgMux 
[![Go report](http://goreportcard.com/badge/bwmarrin/dgmux)](http://goreportcard.com/report/bwmarrin/dgmux) [![Discord Gophers](https://img.shields.io/badge/Discord%20Gophers-%23info-blue.svg)](https://discord.gg/0f1SbxBZjYq9jLBk)
<img align="right" src="https://raw.githubusercontent.com/wiki/bwmarrin/disgord/images/gourd.jpg">

dgMux is a message router (or muxer) that can be used with Disgord or the
[DiscordGo](https://github.com/bwmarrin/discordgo) library.

dgMux allows you to register "commands" or "keywords" that if found in a 
message then dgMux will call the corresponding message handler function.

If you would like to help the Disgord or DiscordGo package please use 
[this link](https://discordapp.com/oauth2/authorize?client_id=173113690092994561&scope=bot)
to add the official DiscordGo test bot **dgo** to your server. This provides 
indispensable help to this project.

**For help with this program or general Go discussion, please join the [Discord 
Gophers](https://discord.gg/0f1SbxBZjYq9jLBk) chat server.**

**NOTE** : This is an experimental package and it's likely to have large changes
breaking it's API and compatibility with previous versions.

Some parts of dgMux were inspired by the [chi][https://github.com/go-chi/chi] 
and [httprouter][https://github.com/julienschmidt/httprouter] routers.
