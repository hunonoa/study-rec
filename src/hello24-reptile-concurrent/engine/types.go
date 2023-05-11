package engine

type ParserResult struct {
	Requests []Request
	//  interface{} 表示任意类型
	Items []interface{}
}

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

func NilParser([]byte) ParserResult {

	return ParserResult{}
}
