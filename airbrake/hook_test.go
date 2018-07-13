package sentry

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var projectID string
var apiKey string
var integration bool

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file" + err.Error())
	}
	err = viper.BindEnv("PROJECT_ID")
	if err != nil {
		panic(err)
	}
	projectID = viper.GetString("PROJECT_ID")

	err = viper.BindEnv("API_KEY")
	if err != nil {
		panic(err)
	}
	apiKey = viper.GetString("API_KEY")

	err = viper.BindEnv("INTEGRATION")
	if err != nil {
		integration = false
	}
	integration = viper.GetBool("INTEGRATION")
	fmt.Printf(dsn)
}

// Unit tests.
// TODO(karel) mock interfaces and add unit tests.

func TestLogAttempt(t *testing.T) {
	LogAttempt(dsn)
}
