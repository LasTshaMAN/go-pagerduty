package pagerduty_test

import (
	"testing"

	"github.com/LasTshaMAN/go-pagerduty"
)

func TestClient_SendTestNotification(t *testing.T) {
	t.Skip("This test is designed to check that interactions with real PD server work properly")

	const (
		authToken = "" // Set your PagerDuty API auth token here.

		userID = "" // Set your user ID here.
		contactMethodID = "" // Set the contact method of your preference here.
	)
	client := pagerduty.NewClient(authToken)

	headers := map[string]string{
		"From":                  "", // Set your user name here.
		"x-pagerduty-api-local": "1",
		"x-requested-with":      "XMLHttpRequest",
	}
	err := client.SendTestNotification(headers, userID, contactMethodID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
