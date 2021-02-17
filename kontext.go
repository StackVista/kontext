//go:generate pkger
package main

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/stackvista/kontext/cmd"
)

func main() {
	ctx := log.Logger.WithContext(context.Background())
	cmd.Execute(ctx)
}
