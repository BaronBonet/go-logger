package main

import (
	go_logger "github.com/BaronBonet/go-logger"
)

func main() {
	log := go_logger.NewZapLogger(false, "1")
	log.Info("Hello world", "key", "value")
}
