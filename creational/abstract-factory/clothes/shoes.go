package clothes

type shoes struct {
	purpose string
	name    string
	size    int
	price   int
}

func (s *shoes) GetPurpose() string {
	return s.purpose
}

func (s *shoes) GetName() string {
	return s.name
}

func (s *shoes) GetSize() int {
	return s.size
}

func (s *shoes) GetPrice() int {
	return s.price
}

type IShoes interface {
	GetPurpose() string
	GetName() string
	GetSize() int
	GetPrice() int
}
