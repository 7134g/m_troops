package pipe_filter

type Requests interface{}
type Response interface{}

type Filter interface {
	Process(data Requests) (Response, error)
}
