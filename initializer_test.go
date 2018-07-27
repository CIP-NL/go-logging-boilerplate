package logrus_hooks

import (
	"github.com/spf13/viper"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/sirupsen/logrus"
)
var c Configuration

func init() {
	viper.SetConfigName("example_config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
}

func TestGenerateHooks(t *testing.T) {
	hooks := GenerateHooks(c.Logrus.Hooks)
	assert.NotEmpty(t, hooks["airbrake"])
	assert.NotEmpty(t, hooks["sentry"])
}

func TestGenerateLoggers(t *testing.T) {
	loggers := GenerateLoggers(c)
	assert.NotEmpty(t, loggers["api_logger"])
	assert.NotEmpty(t, loggers["store_logger"])

	assert.Equal(t, logrus.InfoLevel, loggers["api_logger"].Level)
	assert.Equal(t, logrus.FatalLevel, loggers["store_logger"].Level)

}