package shared

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() {
	loggerOptions := &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true, // TODO: add fn to trim the source file path to omit base path upto the root directory
	}

	Logger = slog.New(slog.NewJSONHandler(os.Stdout, loggerOptions))
}