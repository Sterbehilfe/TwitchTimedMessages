package twitch

import (
	"fmt"
	"strings"
	"time"

	console "TwitchTimedMessages/Console"
	settings "TwitchTimedMessages/Settings"

	linq "github.com/ahmetb/go-linq"
	color "github.com/fatih/color"
	irc "github.com/gempir/go-twitch-irc/v2"
)

const (
	SecInNano = 1000000000
)

type Client struct {
	settings  settings.Settings
	ircClient *irc.Client
}

func NewTwitchClient(s settings.Settings) *Client {
	return &Client{
		settings:  s,
		ircClient: irc.NewClient(s.Username, s.OAuthToken),
	}
}

func (client *Client) Initialize() {
	client.checkMessagesForRateLimiting()
	client.setEvents()

	channels := client.getChannels()

	client.joinChannels(channels)
	console.WriteLine("Connecting...")
	go client.Connect()
	sleep := SecInNano * len(channels)
	time.Sleep(time.Duration(sleep))
	console.WriteLine("Connected!")
	client.createTimers()
}

func (client *Client) Connect() {
	err := client.ircClient.Connect()
	if err != nil {
		panic(err)
	}
}

func (client *Client) Send(msg settings.Message) {
	client.ircClient.Say(msg.Channel, msg.Content)
	console.WriteLine("Sent message to <#" + msg.Channel + ">: " + msg.Content)
}

func (client *Client) getChannels() []string {
	var result []string
	linq.From(client.settings.Messages).SelectT(func(m settings.Message) string {
		return strings.ToLower(m.Channel)
	}).Distinct().ToSlice(&result)
	return result
}

func (client *Client) joinChannels(channels []string) {
	for _, c := range channels {
		client.ircClient.Join(c)
		console.WriteLine("Joined channel <#" + c + ">")
	}
}

func (client *Client) setEvents() {
	client.ircClient.OnConnect(func() {
		console.WriteLine("Client connected")
	})
}

func (client *Client) createTimers() {
	for _, msg := range client.settings.Messages {
		ticker := time.NewTicker(time.Duration(msg.Interval) * time.Millisecond)
		go client.waitForTick(ticker, msg)
	}
}

func (client *Client) waitForTick(ticker *time.Ticker, msg settings.Message) {
	for {
		client.Send(msg)
		<-ticker.C
	}
}

func (client *Client) checkMessagesForRateLimiting() {
	var intervalTooSmall []settings.Message
	linq.From(client.settings.Messages).WhereT(func(m settings.Message) bool {
		return m.Interval < 1300
	}).ToSlice(&intervalTooSmall)
	for _, m := range intervalTooSmall {
		color.Red("[WARNING] Interval for", m.Channel+":", `"`+m.Content+`"`, "too small.")
	}
	if len(intervalTooSmall) > 0 {
		fmt.Println("The interval shouldn't be smaller than 1300ms")
	}
	var messageTooLong []settings.Message
	linq.From(client.settings.Messages).WhereT(func(m settings.Message) bool {
		return len(m.Content) > 500
	}).ToSlice(&messageTooLong)
	for _, m := range messageTooLong {
		color.Red("[WARNING] Message for", m.Channel+":", `"`+m.Content+`"`, "too long.")
	}
	if len(messageTooLong) > 0 {
		fmt.Println("The length of the message can't exceed 500 chars")
	}
}
