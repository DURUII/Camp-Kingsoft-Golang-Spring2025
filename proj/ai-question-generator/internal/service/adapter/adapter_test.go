package adapter_test

import (
	"encoding/json"
	"testing"

	"quizbot/config"
	. "quizbot/internal/service/adapter"
)

func TestDashScopeGenerator(t *testing.T) {
	config.LoadConfig("../../../.env")
	g := NewDashScopeGenerator(config.GetAPIKey(), "qwen-plus")
	q, _ := g.Generate("Gin框架", "Go", 2)
	output, _ := json.MarshalIndent(q, "", " ")
	t.Log(string(output))
}
