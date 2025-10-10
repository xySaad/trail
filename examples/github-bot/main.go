package main

import (
	"fmt"
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
		return
	}
	defer bot.Close()
	fmt.Println("Bot is running...")

	router := trail.New(bot)
	router.Add("GET /", Webhook)

	err = http.ListenAndServe(SERVER_ADDRESS, router)
	if err != nil {
		fmt.Println("http server:", err)
		return
	}
}
