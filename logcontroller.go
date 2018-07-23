package logcontroller

import (
	"os"

	air "github.com/CIP-NL/logrus-hooks/airbrake"
	"github.com/CIP-NL/logrus-hooks/sentry"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

// logAttempt used to test error messages
func logAttempt(projectID int64, apiKey string, env string, dsn string) {
	log.Level = logrus.DebugLevel
}
func SentryHook(dsn string) {
	hook, err := sentry.NewHook(dsn, []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})
	if err == nil {
		log.AddHook(hook)
	}
}

func Airbrake(projectID int64, apiKey string, env string) {
	log.AddHook(air.NewHook(projectID, apiKey, env))
}

func WriteToSTDout() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

}
