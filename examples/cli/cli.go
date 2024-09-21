package main

import "github.com/BaronBonet/go-logger/logger"

func main() {
	log := logger.NewSlogLogger()
	log.Info("Hello world", "key", "value")
}
