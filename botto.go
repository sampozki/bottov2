package main

import (
	"botto/utils"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// We don't want these words in our cristian server
var banList = []string{"neekeri", "nigger"}

var garglList = []string{"gargl", "_gargl_", "GARGL", "GARGLL............", "Gargl 💀", "come on parti lets go gargli"}

func main() {

	// Try to get bot token from env -> fallback to file called "env"
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Println("Env empty, fallbacking to file")
		b, err := os.ReadFile("env")
		if err != nil {
			panic("No token read from env or file: " + err.Error())
		}
		token = strings.TrimSpace(string(b))
	}

	discord, err := discordgo.New("Bot " + token)
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

	// Pre message hook switch here:

	// Switch case for chatting functions
	switch {

	// Test
	case utils.Match(message.Content, "!hello"):
		utils.Msg(discord, message, "hola")
		utils.UpdateStatus(discord, "asd")

	// banlist
	case utils.ContainsAny(message.Content, banList):
		discord.ChannelMessageSend(message.ChannelID, "mur")

	// gargl
	case utils.Match(message.Content, "gargl"):
		if rand.IntN(50) == 2 {
			utils.Msg(discord, message, garglList[rand.IntN(len(garglList))])
		}

	// hakemus
	case utils.Match(message.Content, "hakemus"):
		if rand.IntN(10) == 1 {
			utils.Msg(discord, message, "Hyy-vä")
		} else {
			utils.Msg(discord, message, "Tapan sut")
		}

	// simpsons faces
	case utils.Regex("(?i)(sotd|rotd|fotd)", message.Content):
		site := ""
		url := ""

		switch {
		case strings.Contains(message.Content, "sotd"):
			site = "frinkiac"
		case strings.Contains(message.Content, "fotd"):
			site = "morbotron"
		case strings.Contains(message.Content, "rotd"):
			site = "masterofallscience"
		}

		if len(message.Content) > 5 {
			url = utils.SendTagFace(site, message.Content[5:])
		} else {
			url = utils.SendFace(site)
		}

		utils.Msg(discord, message, url)

	// hyvä botti & paska botti
	case utils.Match(message.Content, "hyvä botti"):
		utils.Msg(discord, message, ":3")

	case utils.Match(message.Content, "paska botti"):
		utils.Msg(discord, message, "HAISTA VITTU!")

	// mau & hau
	case utils.Regex("(?i)^(m+(a+u|o+u|i+u|ä+y|ö+y|i+a+u|j+ä+y))$", message.Content):
		url := utils.Mau()
		utils.Msg(discord, message, url)

	case utils.Regex("(?i)^(?:h+)(?:a+)(?:u+)$", message.Content):
		url := utils.Hau()
		utils.Msg(discord, message, url)

	// yawn & bark
	case utils.Regex("^yawn", message.Content):
		utils.Msg(discord, message, "https://sampozki.fi/yawn.png")

	case utils.Regex("^bark", message.Content):
		utils.Msg(discord, message, "https://sampozki.fi/barkmanul.gif")

	// tulin
	case utils.Regex("^(tu(un|li|ut|le))", message.Content):
		if rand.IntN(3) == 1 {
			utils.Msg(discord, message, "tirsk")
		}

	// ping
	case utils.Regex("!ping", message.Content):
		utils.Msg(discord, message, "pong")
	}

}
