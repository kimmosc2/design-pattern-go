package clothes

import (
	"errors"
)

type ClothesFactory interface {
	buildShoes() IShoes
	buildShirt() IShirt
}

func GetClothes(brand string) (ClothesFactory, error) {
	switch brand {
	case "adidas":
		return adidasFactory{}, nil
	case "nike":
		return nikeFactory{}, nil
	default:
		return nil,errors.New("no specified clothes")
	}
}
