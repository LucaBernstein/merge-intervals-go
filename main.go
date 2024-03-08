package main

import (
	"log/slog"
	"os"
	"strings"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelWarn) // For debugging purposes, set the level to LevelDebug

	inputArgs := os.Args[1:] // skip first arg (program)
	// sanitize input: remove all blanks and join probably multiple args together
	sanitizedInput := strings.ReplaceAll(strings.Join(inputArgs, ""), " ", "")

	slog.Debug("Input", slog.Any("args", sanitizedInput))

	// TODO: Pass input on to parsing and merging logic
}
