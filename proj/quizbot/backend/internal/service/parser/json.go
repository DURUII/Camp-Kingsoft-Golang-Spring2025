package parser

import (
	"regexp"
)

type JSONParser struct{}

func (p JSONParser) ExtractJSONContent(input string) string {
	re := regexp.MustCompile("```json(.*?)```")
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		return matches[1]
	}
	return input
}
