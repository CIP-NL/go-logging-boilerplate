package sentry

import (
	"flag"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// A sentry dns made specifically for this test. Don't use elsewhere please.
var dsn = `https://f73d12f7192f4943b71f1b4ba85a8a73:488214b852684b998ff13698c0525948@sentry.io/1226371`

var integration bool

func init() {
	flag.BoolVar(&integration, "int", false, "run integration tests alongside unit tests")
	flag.Parse()
}

// Unit tests.
// TODO(karel) mock interfaces and add unit tests.

// Integration tests.
func TestNew(t *testing.T) {
	if !integration {
		t.Skip()
	}

	_, err := New(dsn)
	assert.NoError(t, err)
}

func TestSentryHook_Fire(t *testing.T) {
	if !integration {
		t.Skip()
	}

	event := logrus.Entry{}
	client, _ := New(dsn)
	err := client.Fire(&event)
	assert.NoError(t, err)
}

func TestSentryHook_Levels(t *testing.T) {
	if !integration {
		t.Skip()
	}

	client, _ := New(dsn)
	levels := client.Levels()

	assert.Equal(t, levels, []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	})
}
