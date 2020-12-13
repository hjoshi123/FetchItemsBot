# Fetch Items Bot ![Golang](https://img.shields.io/badge/language-Go-green)

This repo contains a telegram bot to fetch your regular items like getting the top news or getting the word of the day. This bot can be found [here](https://t.me/fetchitemsbot).

## Installation

For building the bot, I have used Golang as the language for handling the business logic. To develop or deploy your own handler on this, golang must be installed. To deploy the handler, I have used `heroku`. Once you create an account on heroku this will be pretty straight forward.

* `brew install go`
* `heroku create telebot-go`
* `heroku stack:set container`
* `git add .`
* `git commit -m "[]"`
* `git push heroku master`

Note that `heroku stack:set container` enables docker to be used for deployment. This allows us to use Dockerfile which simplifies our deployment.
Also note that for this handler to work, both the API keys of **`NewsAPI`** and **`Wordnik`** are required. Details of the websites are mentioned at the [end](#api-keys) of the README.

## API Keys

* **`NewsAPI`**

    The API key for NewsAPI is available [here](https://newsapi.org/). Signup here and you will get an API key immediately. There are restrictions for basic account but for personal use this would be enough.

* **`Wordnik`**

    This is again a free service. But in case of `Wordnik` developers would have to wait for a week in order to receive their API keys. Head over to [wordnik](https://wordnik.com) and sign up. Once signed up, you can visit [here](https://developer.wordnik.com/) where you will be asked to enter your wordnik username, upon which you can submit and wait for a week before you get the key.
