package main

import (
	"fmt"
)

func Add(x, y int) (res int) {
	return x + y
}

func Subtract(x, y int) (res int) {
	return x - y
}

func main() {
	fmt.Println(Add(10, 20))
}
