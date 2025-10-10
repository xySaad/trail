package main

import (
	"fmt"
	"net/http"

	"github.com/xySaad/gocord"
	"github.com/xySaad/trail"
)

const CHANNEL_ID = "1490400969990666"

type PullRequestPayload struct {
	Number      int                          `json:"number"`
	PullRequest struct{ Title, Body string } `json:"pull_request"`
}

func Webhook(c *trail.Context[*gocord.Bot]) {
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
