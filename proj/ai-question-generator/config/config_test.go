package config_test

import (
	. "quizbot/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("../.env")
	t.Log(GetAPIKey())
}
