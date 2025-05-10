package service

import (
	"encoding/json"
	"testing"
)

func TestQuestion(t *testing.T) {
	q := &ProgrammingQuestion{
		Title:          "Go 语言中，不属于关键字的是（）",
		Options:        []string{"func", "struct", "class", "defer"},
		CorrectAnswers: []int{2},
	}
	output, _ := json.MarshalIndent(q, "", " ")
	t.Log(string(output))
}
