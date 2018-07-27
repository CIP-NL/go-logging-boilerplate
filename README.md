[![Go Report Card](https://goreportcard.com/badge/github.com/CIP-NL/logrus-hooks)](https://goreportcard.com/report/github.com/CIP-NL/logrus-hooks)
[![Build Status](https://travis-ci.org/CIP-NL/logrus-hooks.svg?branch=master)](https://travis-ci.org/CIP-NL/logrus-hooks)

# logrus hooks
Different logging hooks for logrus (go based logging service)

Also implements two functions for easy configuration: `GenerateHooks` and `GenerateLoggers`.

These functions parse a struct constructed as follows:

```go
type Myconf struct {
	Logrus struct {
		Hooks []struct {
			Name string `toml:"name"`
			Type string `toml:"type"`
			ProjectID int `toml:"project_id,omitempty"`
			APIKey string `toml:"api_key,omitempty"`
			Environment string `toml:"environment,omitempty"`
			Backup string `toml:"backup,omitempty"`
			Kind string `toml:"kind,omitempty"`
			DNS string `toml:"dns,omitempty"`
			Level string `toml:"level,omitempty"`
		} `toml:"hooks"`
		Loggers []struct {
			Name string `toml:"name"`
			Level string `toml:"level"`
			Hooks []struct {
				Name string `toml:"name"`
			} `toml:"hooks"`
		} `toml:"loggers"`
	} `toml:"logrus"`
	
	
	SomeOtherfield myotherTypeStruct
	...
	...
}
```

The example uses a TOML file to parse the logging configuration: 

```toml
[logrus]

[[logrus.hooks]]
    name = "airbrake"
    type = "airbrake"
    project_id = 1
    api_key = ""
    environment = "local"

    backup = "sentry"   # Name of the backup hook

[[logrus.hooks]]
    name = "sentry"
    type = "sentry"
    kind = "default"    # Options: default, async
    dns = ""
    level = "WARN"      # Options: DEBUG, INFO, WARN, ERROR, CRITICAL


[[logrus.loggers]]
    name = "api_logger"
    level = "INFO"
    [[logrus.loggers.hooks]]
        name = "airbrake"

[[logrus.loggers]]
    name = "store_logger"
    level = "CRITICAL"
    [[logrus.loggers.hooks]]
        name = "sentry"
```

99% of the time you will only need to call GenerateLoggers(MyConf). This will return a `map[string]*logrus.Logger` where the key is the specified name in the config file.

