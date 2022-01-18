package pipe_filter

import "errors"

var SumError = errors.New("err data sum")

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (s *SumFilter) Process(data Requests) (Response, error) {
	var output int
	parts, err := data.([]int)
	if !err {
		return nil, SumError
	}

	for _, v := range parts {
		output += v
	}

	return output, nil

}
