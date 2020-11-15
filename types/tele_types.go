package types

// Update is the type of request that telegram sends once u send message to the bot
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message is the structure of the message sent to the bot
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	Date int `json:"date"`
}

// Chat indicates the conversation to which the message belongs.
type Chat struct {
	ID int `json:"id"`
}

// User is a telegram user
type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
}
