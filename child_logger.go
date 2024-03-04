package main

import (
	"log/slog"
	"os"
	"runtime"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stderr,
		&slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	childLogger := slog.With(
		slog.Group("OS Info",
			slog.String("OS", runtime.GOOS),
			slog.Int("CPUs", runtime.NumCPU()),
			slog.String("arch", runtime.GOARCH),
		),
	)

	childLogger.Info(
		"Golang rocks!",
		slog.String("version", runtime.Version()),
	)

}

