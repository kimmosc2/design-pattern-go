package parse

type json struct {
	data []byte
}

func (j json) Parse() []byte {
	return []byte{}
}

func NewJsonParser() Parser {
	return json{}
}
