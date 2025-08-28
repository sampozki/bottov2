package utils

import (
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CheckNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func LogText(ID string, text string) {
	log.Println(ID + ": " + text)
}

// Returns true if some []words in content string, otherwise false
func ContainsAny(content string, words []string) bool {
	for _, w := range words {
		if strings.Contains(content, w) {
			return true
		}
	}
	return false
}

func Msg(discord *discordgo.Session, message *discordgo.MessageCreate, toSend string) {
	msg := toSend
	discord.ChannelMessageSend(message.ChannelID, msg)
	LogText(message.ChannelID, msg)
}

func Regex(pattern string, input string) bool {
	return regexp.MustCompile(pattern).MatchString(input)
}
