package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/hjoshi123/fetchitemsbot/types"
)

func parseTelegramUpdate(r *http.Request) (*types.Update, error) {
	var update types.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		return nil, err
	}

	return &update, nil
}

func handler(res http.ResponseWriter, r *http.Request) {
	update, err := parseTelegramUpdate(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}
	fmt.Println(update)

	userCmd := types.ParseCommand(update.Message.Text)
	log.Printf(userCmd)

	if userCmd == "/start" {
		var output string
		output = types.BotName + " has the following commands for you \n"
		for command, desc := range types.Commands {
			output += command + " : " + desc + "\n"
		}
		startResponse, errStart := startCmd(update.Message.Chat.ID, output)
		if errStart != nil {
			log.Printf("got error %s from telegram, response body is %s", errStart.Error(), startResponse)
		} else {
			log.Printf("punchline %s successfully distributed to chat id %d", output, update.Message.Chat.ID)
		}
	} else if userCmd == "/news" {

	} else if userCmd == "/port" {

	}
}

func startCmd(chatID int, text string) (string, error) {
	log.Printf("Sending %s to chat_id: %d", text, chatID)
	response, err := http.PostForm(
		types.TelegramAPI,
		url.Values{
			"chat_id": {strconv.Itoa(chatID)},
			"text":    {text},
		})

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

func main() {
	fmt.Println(os.Getenv("PORT"))
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), http.HandlerFunc(handler))
}
