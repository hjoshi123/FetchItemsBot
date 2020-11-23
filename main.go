package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/hjoshi123/fetchitemsbot/types"
	"github.com/hjoshi123/fetchitemsbot/utils"
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
	log.Printf(update.Message.Text)
	log.Println(update.CallbackQuery.Data)

	userCmd, err := types.ParseCommand(update.Message.Text)
	callbackData := update.CallbackQuery.Data

	log.Println(userCmd)
	// TODO Parse Arguments of the command too

	var output string
	var keyboard []byte = nil

	if err != nil && callbackData == "" {
		output = "Sorry you entered the wrong command. Here are the list of supported commands \n"
		for command, desc := range types.Commands {
			output += command + " : " + desc + "\n"
		}
	}

	// Check if there is a callback from an inline button
	if callbackData == "GN" {
		output, err = utils.GetNewsForResponse("the-times-of-india")
		log.Println(output)

		if err != nil {
			resp, err := utils.SendTextToTelegram(update.CallbackQuery.From.ID, output, keyboard)
			if err != nil {
				log.Printf("got error %s from telegram, response body is %s", err.Error(), resp)
			} else {
				log.Printf("punchline %s successfully distributed to chat id %d", output, update.Message.Chat.ID)
			}
			return
		}
	} else if callbackData == "BN" {
		output, err = utils.GetNewsForResponse("business-insider")
		if err != nil {
			resp, err := utils.SendTextToTelegram(update.CallbackQuery.From.ID, output, keyboard)
			if err != nil {
				log.Printf("got error %s from telegram, response body is %s", err.Error(), resp)
			} else {
				log.Printf("punchline %s successfully distributed to chat id %d", output, update.Message.Chat.ID)
			}
			return
		}
	} else {
		if userCmd == "/start" {
			output = "Hello I'm " + types.BotName + " I can do the following things for you \n\n"
			for command, desc := range types.Commands {
				output += command + " : " + desc + "\n"
			}
		} else if userCmd == "/news" {
			but := types.Buttons{}
			but.CreateInlineButtons(1, 2, "General News", "GN", "Business News", "BN")

			keyboard, err = json.Marshal(but)
			if err != nil {
				log.Printf(err.Error())
				return
			}

			output += "Great.. Almost there.. Please choose which kind of news you want\n"
		} else if userCmd == "/port" {

		}
	}

	var resp string
	if update.Message.Chat.ID != 0 {
		resp, err = utils.SendTextToTelegram(update.Message.Chat.ID, output, keyboard)
	} else if update.CallbackQuery.From.ID != 0 {
		resp, err = utils.SendTextToTelegram(update.CallbackQuery.From.ID, output, keyboard)
	}

	if err != nil {
		log.Printf("got error %s from telegram, response body is %s", err.Error(), resp)
	} else {
		log.Printf("punchline %s successfully distributed to chat id %d", output, update.Message.Chat.ID)
	}
}

func main() {
	log.Printf(os.Getenv("PORT"))
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), http.HandlerFunc(handler))
}
