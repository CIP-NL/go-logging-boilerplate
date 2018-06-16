package sentry

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
	"flag"
)

// A sentry dns made specifically for this test. Don't use elsewhere please.
var dsn = `https://f73d12f7192f4943b71f1b4ba85a8a73:488214b852684b998ff13698c0525948@sentry.io/1226371`

var integration bool

func init() {
	flag.BoolVar(&integration, "int", false, "run integration tests alongside unit tests")
	flag.Parse()
}

// Unit tests.





// Integration tests.
func TestNewSentryHook(t *testing.T) {
	if !integration {t.Skip()}
	_, err := NewSentryHook(dsn)
	if err != nil {
		t.Error("Failed to create a new sentry hook (check the dsn): ", err)
	}
}

func TestSentryHook_Fire(t *testing.T) {
	if !integration {t.Skip()}
	event := logrus.Entry{}
	client, _ := NewSentryHook(dsn)
	err := client.Fire(&event)
	if err != nil {
		t.Error("Failed to fire an event.")
	}
}

func TestSentryHook_Levels(t *testing.T) {
	if !integration {t.Skip()}
	client, _ := NewSentryHook(dsn)
	levels := client.Levels()

	equal := reflect.DeepEqual(levels, []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	})
	if !equal {
		t.Error("Levels are not what they should be...")
	}
}
