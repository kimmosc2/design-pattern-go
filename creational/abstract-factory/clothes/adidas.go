package clothes

type adidasShoes struct {
	shoes
}

type adidasShirt struct {
	shirt
}

type adidasFactory struct{}

func (a adidasFactory) buildShoes() IShoes {
	return &adidasShoes{}
}

func (a adidasFactory) buildShirt() IShirt {
	return &adidasShirt{}
}
