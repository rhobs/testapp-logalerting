package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {

	sum := 0
	for {
		sum++ // repeated forever
		time.Sleep(5 * time.Second)
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Debug("Nothing to worry about here")
		logger.Info("Someone might want to look into this")
		logger.Warn("Someone should probably look into this")
		logger.Error("Panic")
	}
}
