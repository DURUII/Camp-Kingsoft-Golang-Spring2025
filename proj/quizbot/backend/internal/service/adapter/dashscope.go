package adapter

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"

	. "quizbot/internal/service"
	. "quizbot/internal/service/parser"
)

type DashScopeGenerator struct {
	baseURL        string
	apiKey         string
	model          string
	sysPrompt      string
	userPromptFunc func(k string, l string, t QuestionType) string
}

func NewDashScopeGenerator(apiKey string, model string) *DashScopeGenerator {
	return &DashScopeGenerator{
		baseURL: "https://dashscope.aliyuncs.com/compatible-mode/v1/",
		apiKey:  apiKey,
		model:   model,
		sysPrompt: "你是一个互联网行业技术专家，请根据需求出题。" + `以 JSON 的形式输出，输出的 JSON 需遵守以下的格式：
{
  "title": <题目，例如，Go 语言中关于变量的自增和自减操作，下面语句正确的是？>,
  "answers": <选项数组，例如，["i := 1; i++","i := 1; j = i++","i := 1; ++i","i := 1; i--"]>,
  "right": <正确选项下标数组，例如，[0, 2]>
}`,
		userPromptFunc: func(k string, l string, t QuestionType) string {
			if t == SingleChoice {
				return fmt.Sprintf("请针对%s，出%s一道语言的单选题，选项数目控制在4个，不多不少", k, l)
			}
			return fmt.Sprintf("请针对%s，出一道%s语言的多选题，选项数目控制在4个，不多不少", k, l)
		},
	}
}

func (d *DashScopeGenerator) Generate(keyword string, language string, questionType QuestionType) (ProgrammingQuestion, error) {
	client := openai.NewClient(
		option.WithAPIKey(d.apiKey),
		option.WithBaseURL(d.baseURL),
	)
	// Create the chat completion request
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(d.sysPrompt),
				openai.UserMessage(d.userPromptFunc(keyword, language, questionType)),
			},
			Model: d.model, // qwen-plus, deepseek-v3
		},
	)
	if err != nil {
		return ProgrammingQuestion{}, fmt.Errorf("failed to get chat completion: %w", err)
	}
	// For some reason, the response may have unwanted characters at the beginning
	content := JSONParser{}.ExtractJSONContent(chatCompletion.Choices[0].Message.Content)
	// Unmarshal the response into a ProgrammingQuestion struct
	p := new(ProgrammingQuestion)
	if err := json.Unmarshal([]byte(content), p); err != nil {
		return ProgrammingQuestion{}, fmt.Errorf("failed to unmarshal programming question: %w. Here is the output: %s", err, content)
	}
	return *p, nil
}
