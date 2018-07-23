package logcontroller

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	testEnv string = "test"
)

var (
	DSN         string
	integration bool
	projectID   int64
	testAPIKey  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file" + err.Error())
	}

	err = viper.BindEnv("DSN")
	if err != nil {
		panic(err)
	}
	DSN = viper.GetString("DSN")

	err = viper.BindEnv("INTEGRATION")
	if err != nil {
		integration = false
	}
	integration = viper.GetBool("INTEGRATION")

	err = viper.BindEnv("PROJECT_ID")
	if err != nil {
		panic(err)
	}
	projectID = viper.GetInt64("PROJECT_ID")

	err = viper.BindEnv("API_KEY")
	if err != nil {
		panic(err)
	}
	testAPIKey = viper.GetString("API_KEY")
}

func TestLogAttempt(t *testing.T) {
	if !integration {
		t.Skip()
	}
	LogAttempt(projectID, testAPIKey, testEnv, DSN)
}
