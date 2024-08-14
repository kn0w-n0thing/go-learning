package main

import (
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"log"
	"myslack/message"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	id := os.Getenv("SLACK_CHANNEL_ID")
	log.Println("app token:", appToken)
	log.Println("app bot token:", botToken)
	log.Println("channel id:", id)

	//message.StartSocketModeHandler(appToken, botToken)

	blockOptions := message.BuildLeaveRequestMessageOptions()
	api := slack.New(botToken)
	channelID, timestamp, err := api.PostMessage(
		id,
		slack.MsgOptionText("Some text", false),
		slack.MsgOptionAsUser(false), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
		blockOptions,
	)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	}
}
