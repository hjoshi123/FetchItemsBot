package types

// Commands is the command array
var Commands = map[string]string{
	"/start": "Starting the bot help",
	"/news":  "Getting the top headlines",
	"/port":  "Getting portfolio rates right now",
}

// BotName is the name of the bot
const BotName = "@fetchitemsbot"

// ParseCommand parses the string to get the appropriate command
func ParseCommand(userCommand string) string {
	for key := range Commands {
		if len(userCommand) >= len(key) {
			if userCommand[:len(key)] == key {
				userCommand = userCommand[len(key):]
			}
		}
	}

	if len(userCommand) >= len(BotName) {
		if userCommand[:len(BotName)] == BotName {
			userCommand = userCommand[len(BotName):]
		}
	}

	return userCommand
}
