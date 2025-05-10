package service

// 类型/常量
type QuestionType int

const (
	SingleChoice QuestionType = iota + 1
	MultipleChoice
)

// 结构体（用 tag 对应需求文档中的格式）
type ProgrammingQuestion struct {
	Title          string   `json:"title"`
	Options        []string `json:"answers"`
	CorrectAnswers []int    `json:"right"`
}

// 接口（问题生成、解析）
type ProgrammingQuestionGenerator interface {
	Generate(keyword string, language string, questionType QuestionType) (ProgrammingQuestion, error)
}

type ProgrammingQuestionParser interface {
	Parse(content string) (ProgrammingQuestion, error)
}
