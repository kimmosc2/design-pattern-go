package parse

import (
	"errors"
)

// Parser all parser need implement Parser
type Parser interface {
	Parse() []byte
}

func NewParser(parserType string) (Parser, error) {
	switch parserType {
	case "json":
		return NewJsonParser(), nil
	case "yaml":
		return NewYamlParser(), nil
	default:
		return nil, errors.New("no specified parser")
	}
}
