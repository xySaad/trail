package models

import (
	"github.com/xySaad/gocord"
	"github.com/xySaad/trail"
)

type Context struct {
	trail.Context
	Bot *gocord.Bot
}
