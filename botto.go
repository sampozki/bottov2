package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"github.com/bwmarrin/discordgo"
	"botto/utils"
	"regexp"
	"math/rand/v2"
)

// We don't want these words in our cristian server
var banList = []string{"neekeri", "kielletty"}

var garglList = []string{"gargl", "_gargl_", "GARGL", "GARGLL............", "Gargl 💀", "come on parti lets go gargli"}

func main() {

	// Read bot token
	b, err := os.ReadFile("env")
	utils.CheckNilErr(err)

	botToken := string(b)

	discord, err := discordgo.New("Bot " + botToken)
	utils.CheckNilErr(err)

	discord.AddHandler(newMessage)
	discord.Open()

	defer discord.Close()

	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}


// Handler for all new messages
func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

// Doesn't reply to its own messages
 	if message.Author.ID == discord.State.User.ID {
 		return
 	}


 	// Switch case for functions
 	switch {

	// Test
 	case strings.Contains(message.Content, "!hello"):
 		msg := "Hello from Go!"
 		discord.ChannelMessageSend(message.ChannelID, msg)
 		utils.LogText(message.ChannelID, msg)


	// banlist
	case utils.ContainsAny(message.Content, banList):
		discord.ChannelMessageSend(message.ChannelID, "mur")

	// gargl
	case strings.Contains(strings.ToLower(message.Content), "gargl"):
		if rand.IntN(3) == 2 {
			msg := garglList[rand.IntN(len(garglList))]
			discord.ChannelMessageSend(message.ChannelID, msg)
 			utils.LogText(message.ChannelID, msg)
		}


	// hakemus


	// simpsons faces



	// hyvä botti & paska botti


	// mau & hau


	// yawn & bark
	

	// tulin
	case regexp.MustCompile("^(tu(un|li|ut|le))").MatchString(message.Content):
		discord.ChannelMessageSend(message.ChannelID, "tirsk")


	// ping
	case regexp.MustCompile("!ping").MatchString(message.Content):
		msg := "pong"
		discord.ChannelMessageSend(message.ChannelID, msg)
		utils.LogText(message.ChannelID, msg)
 	}

}