package pipe_filter

import (
	"errors"
	"strings"
)

var SplitError = errors.New("err data split")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{
		delimiter: delimiter,
	}
}

func (s *SplitFilter) Process(data Requests) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitError
	}

	parts := strings.Split(str, s.delimiter)

	return parts, nil
}
