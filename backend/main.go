package main

import (
	"math/rand"
	"time"

	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	gocron.Every(2).Seconds().Do(task)

	<-gocron.Start()
}

func task() {
	logEntry := log.WithFields(log.Fields{
		"service": "backend",
	})

	durationMilliseconds := int(2000 + rand.NormFloat64()*1000)

	duration := time.Duration(durationMilliseconds) * time.Millisecond

	logEntry.WithField("duration", duration).Info()
}
