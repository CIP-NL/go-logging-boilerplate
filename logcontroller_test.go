package logcontroller

import "testing"

func TestLogAttempt(t *testing.T) {
	if !integration {
		t.Skip()
	}
	LogAttempt(projectID, testAPIKey, testEnv)
}
