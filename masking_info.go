package main

import (
	"log/slog"
	"os"
	"runtime"
	"strings"
)

func main() {

	replaceAttr := func(groups []string, a slog.Attr) slog.Attr {
		keysToMask := []string{"password", "channel"}
		exists := false
		keyToFind := strings.ToLower(a.Key)
		for _, key := range keysToMask {
			if key == keyToFind {
				exists = true
				break
			}
		}
		if exists {
			a.Value = slog.StringValue("<<MASKED>>")
		}
		return a
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr,
		&slog.HandlerOptions{Level: slog.LevelDebug,
			ReplaceAttr: replaceAttr}))
	slog.SetDefault(logger)

	slog.Info(
		"Golang rocks!",
		slog.String("version", runtime.Version()),
		slog.String("password", "G0r0ck$"),
	)

}

