//go:build !js

package main

import (
	"github.com/pion/turn/v5"
)

func newTURNServer() *turn.Server { _ = "STUB: not implemented"; return nil }
