package logrus_hooks

import (
	"github.com/CIP-NL/logrus-hooks/airbrake"
	"github.com/CIP-NL/logrus-hooks/sentry"
	"github.com/sirupsen/logrus"
)

type Hook struct {
	Name        string `toml:"name"`
	Type        string `toml:"type"`
	ProjectID   int64  `toml:"project_id,omitempty"`
	APIKey      string `toml:"api_key,omitempty"`
	Environment string `toml:"environment,omitempty"`
	Backup      string `toml:"backup,omitempty"`
	Kind        string `toml:"kind,omitempty"`
	DNS         string `toml:"dns,omitempty"`
	Level       string `toml:"level,omitempty"`
}

type Loggers []struct {
	Name  string `toml:"name"`
	Level string `toml:"level"`
	Hooks []struct {
		Name string `toml:"name"`
	}
}

type Logrus struct {
	Hooks   []Hook  `toml:"hooks"`
	Loggers Loggers `toml:"loggers"`
}

// Configuration is just a wrapper used during tests.
type Configuration struct {
	Logrus Logrus `toml:"logrus"`
}

func GenerateHooks(hooks []Hook) map[string]logrus.Hook {
	hks := make(map[string]logrus.Hook)
	// First we generate the hooks with no backups
	for _, h := range hooks {
		if h.Backup == "" {
			switch h.Type {
			case "sentry":
				hks[h.Name] = genSentryHook(h)
			case "airbrake":
				hks[h.Name] = genAirbrakeHook(h)
			}
		}
	}

	for _, h := range hooks {
		if h.Backup != "" {
			switch h.Type {
			case "sentry":
				hks[h.Name] = genSentryHook(h, hks[h.Backup])
			case "airbrake":
				hks[h.Name] = genAirbrakeHook(h, hks[h.Backup])
			}
		}
	}
	return hks
}

func GenerateLoggers(log Logrus) map[string]*logrus.Logger {
	loggers := make(map[string]*logrus.Logger)

	hks := GenerateHooks(log.Hooks)
	for _, l := range log.Loggers {
		logger := logrus.New()
		lvl := getLevelFromString(l.Level)
		logger.SetLevel(lvl)

		if len(l.Hooks) > 0 {
			for _, x := range l.Hooks {
				logger.AddHook(hks[x.Name])
			}
		}
		loggers[l.Name] = logger
	}
	return loggers
}

func genAirbrakeHook(h Hook, backups ...logrus.Hook) logrus.Hook {
	return airbrake.NewHook(h.ProjectID, h.APIKey, h.Environment)
}

func genSentryHook(h Hook, backups ...logrus.Hook) logrus.Hook {
	var hook logrus.Hook
	var err error

	levels := getLevelFromHook(h)

	switch h.Kind {
	case "default":
		hook = sentry.New(h.DNS)
	case "async":
		hook, err = sentry.NewAsyncHook(h.DNS, levels)
	default:
		panic("Did not recognise hook kind for hook: " + h.Name)
	}
	if err != nil {
		panic("Unable to create hook: " + h.Name + err.Error())
	}
	return hook
}

// Helper function to convert levels to []logrus levels.
// Allowed aliases: DEBUG, INFO, WARN, ERROR, CRITICAL
func getLevelFromHook(h Hook) []logrus.Level {
	lvl := []logrus.Level{logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel}

	switch h.Level {
	case "DEBUG":
		return lvl
	case "INFO":
		return lvl[1:]
	case "WARN":
		return lvl[2:]
	case "ERROR":
		return lvl[3:]
	case "CRITICAL":
		return lvl[4:]
	default:
		panic("Unable to determine logging level for hook: " + h.Name)
	}
}

// Helper function to convert levels to []logrus levels.
// Allowed aliases: DEBUG, INFO, WARN, ERROR, CRITICAL
func getLevelFromString(s string) logrus.Level {

	switch s {
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "CRITICAL":
		return logrus.FatalLevel
	default:
		panic("Unable to determine logging level from string: " + s)
	}
}
