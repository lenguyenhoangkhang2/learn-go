package main

import "fmt"

type Brand int

const (
	AdidasBrand Brand = iota
	NikeBrand
)

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand Brand) (ISportsFactory, error) {
	if brand == AdidasBrand {
		return &Adidas{}, nil
	}

	if brand == NikeBrand {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
