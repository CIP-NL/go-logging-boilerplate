package logcontroller

import (
	air "github.com/CIP-NL/logrus-hooks/airbrake"
	"github.com/sirupsen/logrus"
)

// LogAttempt used to test error messages
func LogAttempt(projectID int64, testAPIKey string, testEnv string) {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	//Airbrake
	log.AddHook(air.NewHook(projectID, testAPIKey, testEnv))
	//log.Error("Bitcoin price: 0")
}
