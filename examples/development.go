package main

import go_logger "github.com/BaronBonet/go-logger"

func main() {
	log := go_logger.NewSlogLogger()
	log.Info("Hello world", "key", "value")
}
