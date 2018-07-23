package airbrake

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/airbrake/gobrake"
	"github.com/sirupsen/logrus"
)

// Hook to send exceptions to an exception-tracking service compatible
// with the Airbrake API.
type Hook struct {
	Airbrake *gobrake.Notifier
}

// NewHook Returns a new Airbrake hook given the projectID, apiKey and environment
func NewHook(projectID int64, apiKey, env string) *Hook {
	airbrake := gobrake.NewNotifier(projectID, apiKey)
	airbrake.AddFilter(func(notice *gobrake.Notice) *gobrake.Notice {
		if env == "development" {
			return nil
		}
		notice.Context["environment"] = env
		return notice
	})
	hook := &Hook{
		Airbrake: airbrake,
	}
	return hook
}

// Fire sends the notifyErr to airbrake
func (hook *Hook) Fire(entry *logrus.Entry) error {
	var notifyErr error
	err, ok := entry.Data["error"].(error)
	if ok {
		notifyErr = err
	} else {
		notifyErr = errors.New(entry.Message)
	}
	var req *http.Request
	for k, v := range entry.Data {
		if r, ok := v.(*http.Request); ok {
			req = r
			delete(entry.Data, k)
			break
		}
	}
	notice := hook.Airbrake.Notice(notifyErr, req, 3)
	for k, v := range entry.Data {
		notice.Context[k] = fmt.Sprintf("%s", v)
	}

	hook.Verify(notice)
	return nil
}

// Verify checks whether the airbrake service can be used
func (hook *Hook) Verify(notice *gobrake.Notice) bool {
	if _, err := hook.Airbrake.SendNotice(notice); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send error to Airbrake: %v\n", err)
		return false
	}
	return true
}

// Levels returns the standard levels for logrus
func (hook *Hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

// LogAttempt used to test error messages
// func LogAttempt(projectID int64, testAPIKey string, testEnv string) {
// 	log := logrus.New()
// 	log.Level = logrus.DebugLevel
// 	log.AddHook(NewHook(projectID, testAPIKey, testEnv))
// 	log.Error("Bitcoin price: 0")
// }
