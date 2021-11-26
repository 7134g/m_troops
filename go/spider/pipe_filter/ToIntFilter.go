package pipe_filter

import (
	"errors"
	"strconv"
)

var ToIntError = errors.New("err data toint")

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (s *ToIntFilter) Process(data Requests) (Response, error) {
	var ret []int
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntError
	}

	for _, v := range parts {
		s, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}

	return ret, nil
}
