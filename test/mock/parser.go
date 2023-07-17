package mock

type Parser struct{}

func (Parser) Parse(_ []byte) (interface{}, error) {
	return "AST", nil
}
