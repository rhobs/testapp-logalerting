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
		time.Sleep(10 * time.Second)
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Debug("Nothing to worry about here")
		logger.Info("Testing logging alerting")
		logger.Warn("Generating logs, lots of logs")
		logger.Error("Panic")
	}
}
