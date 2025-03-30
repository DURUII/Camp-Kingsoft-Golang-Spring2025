package pipefilter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("SplitFilterWrongFormatError")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter: delimiter}
}

func (sf *SplitFilter) Process(r Request) (Response, error) {
	str, ok := r.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
