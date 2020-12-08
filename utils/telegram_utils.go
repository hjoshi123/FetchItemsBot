package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hjoshi123/fetchitemsbot/types"
)

// ParseTelegramUpdate takes in the request from telegram and parses Update from it
func ParseTelegramUpdate(r *http.Request) (*types.Update, error) {
	var update types.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		return nil, err
	}

	return &update, nil
}

// SendTextToTelegram sends text to the user
func SendTextToTelegram(chatID int, text string, keyboard []byte) (string, error) {
	log.Printf("Sending to chat_id: %d", chatID)
	log.Printf(string(keyboard))
	log.Printf(text)

	response, err := http.PostForm(
		types.TelegramAPI,
		url.Values{
			"chat_id":      {strconv.Itoa(chatID)},
			"text":         {text},
			"parse_mode":   {"HTML"},
			"reply_markup": {string(keyboard)},
		},
	)

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}

// GetNewsForResponse gets news for the selected button
func GetNewsForResponse(source string) (string, error) {
	output := ""
	article, err := types.GetTopHeadlines(source)
	if err != nil {
		output += "Sorry... couldnt fetch headlines. Please try again later.."
		return output, err
	}

	output += "Here's your headlines for the day \n\n"
	output += article.Title + "\n"
	output += article.Description

	return output, nil
}
