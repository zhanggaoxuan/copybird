package pagerduty

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetName(t *testing.T) {
	var pd PagerDuty
	name := GetName()
	require.Equal(t, "pagerduty", name)
}

func TestGetConfig(t *testing.T) {
	var pd PagerDuty
	conf := GetConfig()
	require.Equal(t, Config{}, conf)
}

func TestInitModule(t *testing.T) {
	var pd PagerDuty
	err := InitModule(Config{})
	require.NoError(t, err, "should not be any error here")
}

func TestRun(t *testing.T) {
	t.Skip("Skip this test: proper config not provided")
	var pd PagerDuty
	err := InitModule(Config{
		AuthToken: "insert auth token here",
		From:      "example@example.com",
	})
	require.NoError(t, err, "should not be any error here")
	err = Run()
	require.NoError(t, err, "should not be any error here")

}