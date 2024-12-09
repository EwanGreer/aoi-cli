/*
Copyright © 2024 Ewan Greer <ewanja.greer@gmail.com>
*/
package main

import (
	"log/slog"
	"os"

	"github.com/EwanGreer/aoi-cli/cmd"
)

func main() {
	slog.SetDefault(
		slog.New(slog.NewTextHandler(os.Stderr, nil)),
	)

	cmd.Execute()
}
