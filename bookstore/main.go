package main

import "fmt"

type IShoe struct {
	x int
	y int
}

type IShirt struct {
    x int
    y int
    z int
}

type ISportsFactory interface {
	makeShoe() IShoe
}

func main() {
	array := make([]uint8, 10)

	fmt.Println(array)

	offset := 1
	copy(array[offset:], []uint8{1, 2, 3})

	fmt.Println(array)

	sub := array[2:]
	fmt.Println(sub)

	sub2 := array[2:4]
	fmt.Println(sub2)
}
