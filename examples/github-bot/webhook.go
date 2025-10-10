package main

import (
	"fmt"
	"github-bot/types"
	"net/http"

	"github.com/xySaad/trail"
)

const CHANNEL_ID = "1490400969990666"

func Webhook(c *trail.Context[types.Dependecies]) {
	fmt.Println("GitHub Event:", c.Header("X-GitHub-Event"))
	var payload types.PullRequestPayload
	c.Json(&payload)

	switch c.Header("X-GitHub-Event") {
	case "pull_request":
		postTitle := fmt.Sprintf("PR #%d: %s", payload.Number, payload.PullRequest.Title)
		c.Dep.Bot.CreatePost(CHANNEL_ID, postTitle, payload.PullRequest.Body, nil)
	}

	c.Response.WriteHeader(http.StatusOK)
	c.Write([]byte("OK"))
}
