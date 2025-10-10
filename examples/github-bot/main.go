package main

import (
	"fmt"
	"github-bot/middlewares"
	"net/http"

	bot "github.com/xySaad/gocord"
	"github.com/xySaad/trail"
)

const BOT_TOKEN = "MTXXX0ODMXXwNDXXXA2NA.GXXXNf.Mt9XXXXXRWo1wXXXXXXvNn4RAq-WXX4"
const SERVER_ADDRESS = "0.0.0.0:8080"

func main() {
	bot, err := bot.New(BOT_TOKEN)
	if err != nil {
		fmt.Println(err)
		// don't return to skip invalid bot token
		// return
	}
	// defer bot.Close()
	fmt.Println("Bot is running...")

	router := trail.New(bot)
	router.Add("GET /", Webhook, middlewares.GithubSignature)

	err = http.ListenAndServe(SERVER_ADDRESS, router)
	if err != nil {
		fmt.Println("http server:", err)
		return
	}
}
