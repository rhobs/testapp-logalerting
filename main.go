package main

import (
    "time"
    "os"
    "log/slog"   
)

func main() {

  sum := 0
  for {
	sum++ // repeated forever
        time.Sleep(5 * time.Second)
        logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
        logger.Debug("This is a Debug message")
        logger.Info("This is a Info message")
        logger.Warn("This is a Warning message")
        logger.Error("This is a Error message")
    }
}
