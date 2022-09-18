package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Request []Request
	Items   []any
}

func NilFunc([]byte) ParseResult {
	return ParseResult{nil, nil}
}
