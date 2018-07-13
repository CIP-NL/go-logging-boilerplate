// Package sentry implements the logrus sentry hooks for sentry based logging.
package sentry

import (
	"errors"
	logrus "github.com/sirupsen/logrus"
	airbrake "gopkg.in/gemnasium/logrus-airbrake-hook.v3"
)

func LogAttempt(dsn string) {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	log := logrus.New()
	log.Level = logrus.DebugLevel


	//Airbreak
	log.AddHook(airbrake.NewHook(189919, "eb231154ea2c7431480a66cd7d174cdc", "test"))
	log.Error("Something failed but I'm not quitting.")
}
