package mock

type Parser struct {
	ParseFunc func([]byte) (interface{}, error)
}

func (p *Parser) Parse(data []byte) (interface{}, error) {
	if p.ParseFunc != nil {
		return p.ParseFunc(data)
	}

	return "AST", nil
}
