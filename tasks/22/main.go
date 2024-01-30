package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1111111111111111111)
	b := big.NewInt(22222222222)

	res := big.NewInt(0)
	res.Mul(a, b)
	fmt.Println("Умножение: ", res)

	res = big.NewInt(0)
	if b.Cmp(big.NewInt(0)) != 0 {
		res.Div(a, b)
	}
	fmt.Println("Деление: ", res)

	res = big.NewInt(0)
	res.Add(a, b)
	fmt.Println("Сложение: ", res)

	res = big.NewInt(0)
	res.Sub(a, b)
	fmt.Println("Вычитание: ", res)
}
