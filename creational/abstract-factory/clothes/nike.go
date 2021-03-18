package clothes

type nikeShoes struct {
	shoes
}

type nikeShirt struct {
	shirt
}

type nikeFactory struct {}

func (n nikeFactory) buildShoes() IShoes {
	return &nikeShoes{}
}

func (n nikeFactory) buildShirt() IShirt {
	return &nikeShirt{}
}


