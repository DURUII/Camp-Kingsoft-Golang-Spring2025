package pipefilter

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("ToIntFilterWrongFormatError")

type ToIntFilter struct{}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tif *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	var ret []int
	for _, part := range parts {
		i, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, i)
	}
	return ret, nil
}
