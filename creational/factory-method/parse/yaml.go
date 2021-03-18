package parse

type yaml struct {
	data []byte
}


func (y yaml) Parse() []byte {
	return []byte{}
}

func NewYamlParser() Parser {
	return yaml{}
}
