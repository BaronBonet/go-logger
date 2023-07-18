package main

import "github.com/BaronBonet/go-logger/logger"

func main() {
	log := logger.NewZapLogger(false, "1")
	log.Info("Hello world", "key", "value")
}
