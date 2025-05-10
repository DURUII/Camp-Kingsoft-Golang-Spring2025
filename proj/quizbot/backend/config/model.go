package config

import (
	. "quizbot/internal/service"
	. "quizbot/internal/service/adapter"
)

var modelToVariant = map[string]string{
	"deepseek": "deepseek-v3",
	"tongyi":   "qwen-plus",
}

// i.e. deepseek is a model name, and deepseek-v3 is a variant name
func SpecifyModelVariant(model string) string {
	return modelToVariant[model]
}

// since generator is a interface, different model can be supported
func NewGenerator(model string) ProgrammingQuestionGenerator {
	switch model {
	case "":
		// default model is tongyi, and its variant would be qwen-plus
		model = "tongyi"
		fallthrough
	case "deepseek", "tongyi":
		// alibaba cloud
		return NewDashScopeGenerator(GetAPIKey(), SpecifyModelVariant(model))
	default:
		// unsupported
		return NewModelNotFoundGenerator()
	}
}
