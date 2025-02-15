package polling

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func runoff(poll Poll, session *discordgo.Session, message *discordgo.MessageCreate) {
	var split []string = strings.Split(getResult(poll, session), ", ")
	var messageOutput string = "Runoff poll:\n"
	var emotes []string
	for _, key := range split {
		var splitEmote []string = strings.Split(key, "(")
		var splitEmoteFin = strings.Split(splitEmote[1], ")")
		emotes = append(emotes, splitEmoteFin[0])

		messageOutput += key + "\n"
	}
	//for each result, add one
	poll.runoffMessage, _ = session.ChannelMessageSend(poll.channel, messageOutput)

	for _, emote := range emotes {
		go session.MessageReactionAdd(poll.channel, poll.runoffMessage.ID, emote)
	}
}

func runoffRes(poll Poll, session *discordgo.Session) string {
	if poll.runoffMessage != nil {
		return getResultHelper(poll, session)
	}
	return ""
}
