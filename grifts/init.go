package grifts

import (
  "github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/go_with_jwt/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
