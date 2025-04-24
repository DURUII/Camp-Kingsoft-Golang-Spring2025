package adapter

import (
	"errors"

	. "quizbot/internal/service"
)

var ErrModelNotFound = errors.New("model not found")

type ModelNotFoundGenerator struct{}

func NewModelNotFoundGenerator() *ModelNotFoundGenerator {
	return &ModelNotFoundGenerator{}
}

func (d *ModelNotFoundGenerator) Generate(keyword string, language string, questionType QuestionType) (ProgrammingQuestion, error) {
	return ProgrammingQuestion{}, ErrModelNotFound
}
