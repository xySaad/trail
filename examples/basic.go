package examples

import (
	"fmt"
	"net/http"

	bot "github.com/xySaad/gocord"
	"github.com/xySaad/trail"
)

const BOT_TOKEN = "MTXXX0ODMXXwNDXXXA2NA.GXXXNf.Mt9XXXXXRWo1wXXXXXXvNn4RAq-WXX4"
const SERVER_ADDRESS = "0.0.0.0:8080"
const CHANNEL_ID = "1490400969990666"

type PullRequestPayload struct {
	Number      int                          `json:"number"`
	PullRequest struct{ Title, Body string } `json:"pull_request"`
}

func Webhook(c *trail.Context[*bot.Bot]) {
	fmt.Println("GitHub Event:", c.Header("X-GitHub-Event"))
	var payload PullRequestPayload
	c.Json(&payload)

	switch c.Header("X-GitHub-Event") {
	case "pull_request":
		postTitle := fmt.Sprintf("PR #%d: %s", payload.Number, payload.PullRequest.Title)
		c.Dep.CreatePost(CHANNEL_ID, postTitle, payload.PullRequest.Body, nil)
	}

	c.Response.WriteHeader(http.StatusOK)
	c.Write([]byte("OK"))
}

func Basic() {
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
