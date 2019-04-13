package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	commandPrefix string
	botID         string
)

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

func main() {
	discord, err := discordgo.New("Bot ")
	errCheck("Error Creating Disord session", err)
	user, err := discord.User("@me")
	errCheck("Error Retrieving Account", err)

	botID := user.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err := discord.UpdateStatus(0, "Discord.go")
		if err != nil {
			fmt.Println("Error Updating The Status")
			servers := discord.State.Guilds
			fmt.Printf("The Bot is online on %d servers ", len(servers))
			fmt.Println(botID)
		}
	})

	err = discord.Open()
	errCheck("Connection eror", err)
	defer discord.Close()

	commandPrefix = ")"
	<-make(chan struct{})
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the bot is talking
		return
	}

	fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author)
}
