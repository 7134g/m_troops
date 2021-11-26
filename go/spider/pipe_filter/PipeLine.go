package pipe_filter

import "errors"

type PipeLine struct {
	Name    string
	Filters *[]Filter
}

var PipeError = errors.New("err data pipe")

func NewPipeLine(n string, filte ...Filter) *PipeLine {
	return &PipeLine{
		Name:    n,
		Filters: &filte,
	}
}

func (p *PipeLine) Precess(data Requests) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *p.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return nil, PipeError
		}
		data = ret
	}

	return ret, nil

}
