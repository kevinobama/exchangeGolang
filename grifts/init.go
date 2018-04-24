package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kevingates/exchange/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
