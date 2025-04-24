package parser_test

import (
	"testing"

	. "quizbot/internal/service/parser"
)

func TestJSONParser(t *testing.T) {
	j := JSONParser{}
	input := "这里是一些额外的内容，" +
		"```json {\"name\":\"John\", \"age\":30, \"city\":\"New York\"} ```" +
		" 这是后续内容"
	t.Log(j.ExtractJSONContent(input))
}
