package main

import (
	"golang-docker-demo/logger"
	"time"
)

func main() {
	forever := make(chan int)
	logger.Newlogger(logger.ConfigLogger{})
	log := logger.GetLogger()
	for tick := range time.Tick(time.Millisecond) {
		log.Debug(tick)
	}
	<-forever
}
