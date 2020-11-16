package types

import "os"

// Update is the type of request that telegram sends once u send message to the bot
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message is the structure of the message sent to the bot
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	Date int    `json:"date"`
}

// Chat indicates the conversation to which the message belongs.
type Chat struct {
	ID int `json:"id"`
}

// User is a telegram user
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

const telegramAPIBaseURL string = "https://api.telegram.org/bot"
const telegramAPISendMessage string = "/sendMessage"
const telegramTokenEnv string = "TELEGRAM_BOT_TOKEN"

// TelegramAPI is the api to which we should send the message to
var TelegramAPI string = telegramAPIBaseURL + os.Getenv(telegramTokenEnv) + telegramAPISendMessage
