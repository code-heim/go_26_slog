package main

import (
	"log/slog"
	"os"
)

func main() {

	file, err := os.OpenFile(
		"/tmp/slog_demo.log", // Better to get this from an env variable
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666)

	if err != nil {
		panic("Error opening the log file")
	}

	defer file.Close()

	logger := slog.New(slog.NewJSONHandler(multiWriter,
		&slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))
	slog.SetDefault(logger)

	slog.Info("Go rocks!")
}

