package engine

//Request ...
type Request struct {
	URL        string
	ParserFunc func([]byte) ParseResult
}

//ParseResult ...
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//NilParse ...
func NilParse([]byte) ParseResult {
	return ParseResult{}
}
