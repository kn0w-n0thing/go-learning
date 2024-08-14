package message

import (
	"encoding/json"
	"github.com/slack-go/slack/socketmode"
	"log"
	"os"

	"github.com/slack-go/slack"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	TeamID   string `json:"team_id"`
}

type Container struct {
	Type        string `json:"type"`
	MessageTS   string `json:"message_ts"`
	ChannelID   string `json:"channel_id"`
	IsEphemeral bool   `json:"is_ephemeral"`
}

type Team struct {
	ID     string `json:"id"`
	Domain string `json:"domain"`
}

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"`
}

type Field struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Verbatim bool   `json:"verbatim"`
}

type Block struct {
	Type     string    `json:"type"`
	BlockID  string    `json:"block_id"`
	Text     *Text     `json:"text,omitempty"`
	Fields   []Field   `json:"fields,omitempty"`
	Elements []Element `json:"elements,omitempty"`
}

type Element struct {
	Type     string `json:"type"`
	ActionID string `json:"action_id"`
	Text     *Text  `json:"text,omitempty"`
	Style    string `json:"style,omitempty"`
	Value    string `json:"value,omitempty"`
}

type Message struct {
	User   string  `json:"user"`
	Type   string  `json:"type"`
	Ts     string  `json:"ts"`
	BotID  string  `json:"bot_id,omitempty"`
	AppID  string  `json:"app_id,omitempty"`
	Text   string  `json:"text"`
	Team   string  `json:"team"`
	Blocks []Block `json:"blocks"`
}

type Action struct {
	ActionID string `json:"action_id"`
	BlockID  string `json:"block_id"`
	Text     *Text  `json:"text,omitempty"`
	Value    string `json:"value"`
	Style    string `json:"style,omitempty"`
	Type     string `json:"type"`
	ActionTS string `json:"action_ts"`
}

type Payload struct {
	Type                string      `json:"type"`
	User                User        `json:"user"`
	APIAppID            string      `json:"api_app_id"`
	Token               string      `json:"token"`
	Container           Container   `json:"container"`
	TriggerID           string      `json:"trigger_id"`
	Team                Team        `json:"team"`
	Enterprise          interface{} `json:"enterprise"` // or use specific type if known
	IsEnterpriseInstall bool        `json:"is_enterprise_install"`
	Channel             Channel     `json:"channel"`
	Message             Message     `json:"message"`
	State               struct {
		Values map[string]interface{} `json:"values"` // Adjust as needed
	} `json:"state"`
	ResponseURL string   `json:"response_url"`
	Actions     []Action `json:"actions"`
}

func StartSocketModeHandler(appToken string, botToken string) {
	api := slack.New(
		botToken,
		slack.OptionDebug(false),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
		slack.OptionAppLevelToken(appToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(false),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	userInteractionHandler := socketmode.NewSocketmodeHandler(client)
	userInteractionHandler.HandleInteraction(slack.InteractionTypeBlockActions, middlewareInteractionTypeBlockActions)
	err := userInteractionHandler.RunEventLoop()
	if err != nil {
		log.Fatalln("Failed to start a user interaction handler, error:", err)
	}
}

func middlewareInteractionTypeBlockActions(evt *socketmode.Event, client *socketmode.Client) {
	client.Debugf("button clicked!")
	//var payload Payload
	//err := json.Unmarshal(evt.Request.Payload, &payload)
	//if err != nil {
	//	fmt.Println("Error unmarshalling JSON:", err)
	//	return
	//}
	//log.Println("response url:", payload)

	request, err := json.Marshal(evt.Request.Payload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return
	}
	log.Println(string(request))

}
