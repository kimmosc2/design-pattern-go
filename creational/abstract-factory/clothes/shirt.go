package clothes

type shirt struct {
	style string
	size  int
	price int
	name  string
}

func (s *shirt) GetStyle() string {
	return s.style
}

func (s *shirt) GetName() string {
	return s.name
}

func (s *shirt) GetSize() int {
	return s.size
}

func (s *shirt) GetPrice() int {
	return s.price
}

type IShirt interface {
	GetStyle() string
	GetName() string
	GetSize() int
	GetPrice() int
}
