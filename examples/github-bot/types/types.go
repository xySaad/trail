package types

import "github.com/xySaad/gocord"

type PullRequestPayload struct {
	Number      int                          `json:"number"`
	PullRequest struct{ Title, Body string } `json:"pull_request"`
}

type Dependecies struct {
	Bot           *gocord.Bot
	SkipSignature bool
}
